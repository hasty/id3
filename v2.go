package id3

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"io"
	"os"

	"github.com/golang/glog"
)

const (
	id3v23FrameSizeSize  uint32 = 4
	id3v23FrameFlagsSize uint32 = 2
)

type frameFactory struct {
	description string
	maker       func(*Tag, *frameHeader, []byte) (Frame, error)
}

type versionParams struct {
	frameIdSize        uint32
	frameSizeSize      uint32
	frameFlagsSize     uint32
	sizeUnsynchronized bool
	frames             map[string]*frameFactory
}

func (tag *Tag) readV2(framesSize uint32, params *versionParams, r io.ReadSeeker) error {
	var i uint32
	var n int
	var err error
	br := bufio.NewReader(r)
	frameId := make([]byte, params.frameIdSize)
	frameSize := make([]byte, params.frameSizeSize)
	frameFlags := make([]byte, params.frameFlagsSize)
	for i < framesSize {

		n, err = br.Read(frameId[:1])
		if err != nil {
			return err
		}
		if n != 1 {
			return ErrTooShort
		}
		if frameId[0] == 0x0 {
			// This is the end of the frames; we're in padding now
			tag.Header.paddingSize = framesSize - i
			r.Seek(int64(tag.Header.paddingSize-1), os.SEEK_CUR)
			break
		}
		var statusFlags, formatFlags byte
		n, err = br.Read(frameId[1:params.frameIdSize])
		if err != nil || n != int(params.frameIdSize-1) {
			return ErrTooShort
		}
		i += params.frameIdSize
		n, err = br.Read(frameSize[0:params.frameSizeSize])
		if err != nil || n != int(params.frameSizeSize) {
			return ErrTooShort
		}
		i += params.frameSizeSize
		if params.frameFlagsSize > 0 {
			n, err = br.Read(frameFlags[0:params.frameFlagsSize])
			if err != nil || n != int(params.frameFlagsSize) {
				return ErrTooShort
			}
			if params.frameFlagsSize > 0 {
				statusFlags = frameFlags[0]
			}
			if params.frameFlagsSize > 1 {
				formatFlags = frameFlags[1]
			}
			i += params.frameFlagsSize
		}
		var frameLength uint32
		if params.sizeUnsynchronized {
			frameLength = unsafe(frameSize[:])
		} else {
			if params.frameSizeSize == 4 {
				frameLength = binary.BigEndian.Uint32(frameSize[:])
			} else {
				frameLength = uint32(frameSize[0])<<16 | uint32(binary.BigEndian.Uint16(frameSize[1:]))
			}
		}
		if frameLength > framesSize-i {
			return ErrTooShort
		}
		data := make([]byte, frameLength)
		n, err = br.Read(data)
		if err != nil || n != int(frameLength) {
			return ErrTooShort
		}
		i += frameLength

		//glog.Infof("TAG: %v, LENGTH: %v", string(frameId[:]), frameLength)
		//glog.Infof("DATA: %v", hex.EncodeToString(data))
		factory, ok := params.frames[string(frameId[:])]
		if !ok {
			glog.Errorf("Unknown tag: %v", string(frameId[:]))
			glog.Errorf("DATA: %v", hex.EncodeToString(data))
			continue
		}
		frame, err := factory.maker(tag, newFrameHeader(string(frameId), statusFlags, formatFlags, frameLength), data)
		if err != nil {
			glog.Errorf("Error parsing tag %v: %v", string(frameId[:]), err)
			glog.Errorf("DATA: %v", hex.EncodeToString(data))
			continue
		}

		/*
			switch t := frame.(type) {
			case *DataFrame:
				glog.Infof("DATA: %v", hex.EncodeToString(frame.Bytes()))
			case *TextFrame:
				glog.Infof("TEXT %v: %v", len(t.String()), t.String())
			//for index, runeValue := range t.String() {
			//	glog.Infof("%#U starts at byte position %d\n", runeValue, index)
			//}
			case *PictureFrame:
				glog.Infof("PIC %v: %v", t.String(), len(t.Bytes()))
			//glog.Infof("PICDATA: %v", hex.EncodeToString(t.Bytes()))
			//out := fmt.Sprintf("%v.png", time.Now().UnixNano())
			//glog.Infof("Wrote %v: %v", out, len(t.Bytes()))
			//ioutil.WriteFile(out, t.Bytes(), 0)
			case *FullTextFrame:
				glog.Infof("FULLTEXT %v: %v", t.Description(), t.String())

			}
		*/

		tag.addFrame(frame)

	}
	return nil
}

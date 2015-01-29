package id3

import (
	"errors"
	"fmt"

	"golang.org/x/text/encoding/charmap"
)

type PictureType byte

const (
	PictureTypeOther PictureType = iota
	PictureTypeFileIcon
	PictureTypeFrontCover
	PictureTypeBackCover
	PictureTypeLeafletPage
	PictureTypeMedia
	PictureTypeLeadArtist
	PictureTypeArtist
	PictureTypeConductor
	PictureTypeLyricist
	PictureTypeRecordingLocation
	PictureTypeDuringRecording
	PictureTypeDuringPerformance
	PictureTypeVideoCapture
	PictureTypeBrightColoredFish
	PictureTypeIllustration
	PictureTypeBandLogo
	PictureTypePublisherLogo
)

type PictureFrame struct {
	frameBase

	pictureType PictureType
	description string
	mime        string
	data        []byte
}

func newPictureFrame(tag *Tag, header *frameHeader, data []byte) (Frame, error) {
	pf := &PictureFrame{}
	pf.header = header

	var err error
	switch tag.Header.version {
	case 2:
		err = pf.read22(data)
	case 3, 4:
		err = pf.read23(data)
	default:
		return nil, errors.New(fmt.Sprintf("Unknown picture frame revision: %v", tag.Header.version))
	}

	if err != nil {
		return nil, err
	}
	return pf, nil
}

func (pf *PictureFrame) read22(data []byte) error {
	l := len(data)
	if l < 7 {
		return ErrTooShort
	}
	textEncoding, encoding, err := extractEncoding(l, data)
	if err != nil {
		return err
	}
	imgFmt := string(data[1:4])
	switch imgFmt {
	case "PNG":
		pf.mime = "image/png"
	case "JPG":
		pf.mime = "image/jpeg"
	default:
		return errors.New(fmt.Sprintf("Unknown picture format: %v", imgFmt))
	}
	pf.pictureType = PictureType(data[4])
	description, j, err := trimForEncoding(l-5, data[5:], textEncoding, false)
	if err != nil {
		return err
	}
	pf.description, err = decodeString(description, encoding)
	pf.data = data[5+j:]
	return nil
}

func (pf *PictureFrame) read23(data []byte) error {
	l := len(data)
	textEncoding, encoding, err := extractEncoding(l, data)
	if err != nil {
		return err
	}

	mime, i, err := trimForEncoding(l, data, ISO88591, true)
	i++

	pf.mime, err = decodeString(mime, charmap.Windows1252)
	if err != nil {
		return err
	}
	pf.pictureType = PictureType(data[i])

	description, j, err := trimForEncoding(l-i, data[i:], textEncoding, false)
	if err != nil {
		return err
	}
	pf.description, err = decodeString(description, encoding)

	pf.data = data[i+j:]
	return nil
}

func (pf *PictureFrame) Bytes() []byte {
	return pf.data
}

func (pf *PictureFrame) String() string {
	return fmt.Sprintf("%v of type %v (%v bytes)", pf.description, pf.mime, len(pf.data))
}

package id3

import (
	"fmt"
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

func newPictureFrame(header *frameHeader, data []byte) (Frame, error) {
	pf := &PictureFrame{}
	pf.header = header
	val, err := readString(data)
	if err != nil {
		return nil, err
	}
	pf.description = val
	return pf, nil
}

func (pf *PictureFrame) Bytes() []byte {
	return pf.data
}

func (pf *PictureFrame) String() string {
	return fmt.Sprintf("%v (%v bytes)", pf.mime, len(pf.data))
}

package id3

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/text/encoding/charmap"
)

type UniqueFileIDFrame struct {
	frameBase

	owner    string
	uniqueID []byte
}

func newUniqueFileIDFrame(tag *Tag, header *frameHeader, data []byte) (Frame, error) {
	uidf := &UniqueFileIDFrame{}
	uidf.header = header

	l := len(data)

	var err error

	owner, i, err := trimForEncoding(l, data, ISO88591, true)
	if err != nil {
		return nil, err
	}
	i++

	uidf.owner, err = decodeString(owner, charmap.Windows1252)
	if err != nil {
		return nil, err
	}

	uidf.uniqueID = data[i:]

	return uidf, nil
}

func (uidf *UniqueFileIDFrame) String() string {
	return fmt.Sprintf("%v (%v)", hex.EncodeToString(uidf.uniqueID), uidf.owner)
}

func (uidf *UniqueFileIDFrame) Bytes() []byte {
	return uidf.uniqueID
}

func (uidf *UniqueFileIDFrame) Owner() string {
	return uidf.owner
}

func (uidf *UniqueFileIDFrame) UniqueID() []byte {
	return uidf.uniqueID
}

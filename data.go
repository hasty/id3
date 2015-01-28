package id3

import (
	"encoding/hex"
)

type DataFrame struct {
	frameBase

	value []byte
}

func newDataFrame(tag *Tag, header *frameHeader, data []byte) (Frame, error) {
	df := &DataFrame{}
	df.header = header
	df.value = data
	return df, nil
}

func (df *DataFrame) String() string {
	return hex.EncodeToString(df.value)
}

func (df *DataFrame) Bytes() []byte {
	return []byte(df.value)
}

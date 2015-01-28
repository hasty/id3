package id3

import "io"

type ExtendedHeader struct {
	size uint32
	data []byte
}

const extendedHeaderSizeSize uint = 4

func newExtendedHeader(r io.ReadSeeker) (*ExtendedHeader, error) {
	s := make([]byte, extendedHeaderSizeSize)
	n, err := io.ReadFull(r, s)
	if err != nil {
		return nil, err
	}
	if uint(n) < extendedHeaderSizeSize {
		return nil, ErrCorruptExtendedHeader
	}
	size := unsafe(s)
	if size < 6 {
		return nil, ErrCorruptExtendedHeader
	}
	e := &ExtendedHeader{
		size: size,
		data: make([]byte, size),
	}
	n, err = io.ReadFull(r, e.data)
	if err != nil {
		return nil, err
	}
	if uint32(n) != size {
		return nil, ErrCorruptExtendedHeader
	}
	return e, nil

}

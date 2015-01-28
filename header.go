package id3

import "io"

const headerSize uint32 = 10

type Header struct {
	version  uint8
	revision uint8

	flags       byte
	frameSize   uint32
	paddingSize uint32
}

func newHeader(r io.ReadSeeker) (*Header, error) {
	header := make([]byte, headerSize)
	n, err := io.ReadFull(r, header)
	if err != nil {
		return nil, err
	}
	if uint32(n) < headerSize {
		return nil, ErrTooShort
	}
	if string(header[:3]) != "ID3" {
		return nil, ErrNoHeader
	}
	h := &Header{
		version:   header[3],
		revision:  header[4],
		flags:     header[5],
		frameSize: unsafe(header[6:]),
	}
	return h, nil
}

func (header *Header) Size() uint32 {
	return header.frameSize
}

func (header *Header) PaddingSize() uint32 {
	return header.paddingSize
}

func (header *Header) Version() uint8 {
	return header.version
}

func (header *Header) Revision() uint8 {
	return header.revision
}

func (header *Header) Unsynchronized() bool {
	return (header.flags & 0x80) == 0x80
}

func (header *Header) HasExtendedHeader() bool {
	return (header.flags & 0x40) == 0x40
}

func (header *Header) IsExperimental() bool {
	return (header.flags & 0x20) == 0x20
}

func (header *Header) HasFooter() bool {
	return (header.flags & 0x10) == 0x10
}

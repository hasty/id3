package id3

type Frame interface {
	Id() string
	Size() uint32
	StatusFlags() byte
	FormatFlags() byte
	String() string
	Bytes() []byte
}

type frameBase struct {
	header *frameHeader
}

func (fb *frameBase) Id() string {
	return fb.header.id
}

func (fb *frameBase) StatusFlags() byte {
	return fb.header.statusFlags
}

func (fb *frameBase) FormatFlags() byte {
	return fb.header.formatFlags
}

func (fb *frameBase) Size() uint32 {
	return fb.header.size
}

type frameHeader struct {
	id          string
	statusFlags byte
	formatFlags byte
	size        uint32
}

func newFrameHeader(id string, statusFlags byte, formatFlags byte, size uint32) *frameHeader {
	return &frameHeader{
		id:          id,
		statusFlags: statusFlags,
		formatFlags: formatFlags,
		size:        size,
	}
}

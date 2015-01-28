package id3

type TextFrame struct {
	frameBase

	value string
}

func newTextFrame(tag *Tag, header *frameHeader, data []byte) (Frame, error) {
	tf := &TextFrame{}
	tf.header = header
	val, err := readString(data)
	if err != nil {
		return nil, err
	}
	tf.value = val
	return tf, nil
}

func simpleTextFrame(tag *Tag, id string, val string) Frame {
	tf := &TextFrame{}

	tf.header = newFrameHeader(id, 0, 0, uint32(len(val)))
	tf.value = val
	return tf
}

func (tf *TextFrame) String() string {
	return tf.value
}

func (tf *TextFrame) Bytes() []byte {
	return []byte(tf.value)
}

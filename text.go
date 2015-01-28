package id3

type TextFrame struct {
	frameBase

	value string
}

func newTextFrame(header *frameHeader, data []byte) (Frame, error) {
	tf := &TextFrame{}
	tf.header = header
	val, err := readString(data)
	if err != nil {
		return nil, err
	}
	tf.value = val
	return tf, nil
}

func (tf *TextFrame) String() string {
	return tf.value
}

func (tf *TextFrame) Bytes() []byte {
	return []byte(tf.value)
}

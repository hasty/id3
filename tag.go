package id3

type Tag struct {
	Header         *Header
	ExtendedHeader *ExtendedHeader

	frameMap map[string][]Frame
}

func newTag(header *Header, extendedHeader *ExtendedHeader) *Tag {
	return &Tag{
		Header:         header,
		ExtendedHeader: extendedHeader,
		frameMap:       make(map[string][]Frame),
	}
}

func (tag *Tag) addFrame(frame Frame) {
	frames, ok := tag.frameMap[frame.Id()]
	frames = append(frames, frame)
	if !ok {
		tag.frameMap[frame.Id()] = frames
	}
}

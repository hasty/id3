package id3

type Tag struct {
	Header         *Header
	ExtendedHeader *ExtendedHeader

	frameMap      map[string][]Frame
	titleFrame    Frame
	artistFrame   Frame
	albumFrame    Frame
	yearFrame     Frame
	genreFrame    Frame
	commentFrames []Frame
}

func newTag(header *Header, extendedHeader *ExtendedHeader) *Tag {
	return &Tag{
		Header:         header,
		ExtendedHeader: extendedHeader,
		frameMap:       make(map[string][]Frame),
	}
}

func emptyTag() *Tag {
	return &Tag{
		frameMap: make(map[string][]Frame),
	}
}

func (tag *Tag) addFrame(frame Frame) {
	id := frame.Id()
	frames, ok := tag.frameMap[id]
	frames = append(frames, frame)
	if !ok {
		tag.frameMap[id] = frames
	}
	switch id {
	case "TT2", "TIT2":
		tag.titleFrame = frame
	case "TP1", "TPE1":
		tag.artistFrame = frame
	case "TAL", "TALB":
		tag.albumFrame = frame
	case "TYE", "TYER":
		tag.yearFrame = frame
	case "TCO", "TCON":
		tag.genreFrame = frame
	case "COM", "COMM":
		tag.commentFrames = append(tag.commentFrames, frame)
	}
}

func (tag *Tag) Title() string {
	if tag.titleFrame != nil {
		return tag.titleFrame.String()
	}
	return ""
}

func (tag *Tag) Artist() string {
	if tag.artistFrame != nil {
		return tag.artistFrame.String()
	}
	return ""

}

func (tag *Tag) Album() string {
	if tag.albumFrame != nil {
		return tag.albumFrame.String()
	}
	return ""

}

func (tag *Tag) Year() string {
	if tag.yearFrame != nil {
		return tag.yearFrame.String()
	}
	return ""

}

func (tag *Tag) Genre() string {
	if tag.genreFrame != nil {
		return tag.genreFrame.String()
	}
	return ""

}

func (tag *Tag) Comments() []string {
	if len(tag.commentFrames) > 0 {
		var comments []string
		for _, comment := range tag.commentFrames {
			comments = append(comments, comment.String())
		}
		return comments
	}
	return []string{}
}

func (tag *Tag) missingCoreInfo() bool {
	return tag.albumFrame == nil || tag.artistFrame == nil || tag.genreFrame == nil || tag.titleFrame == nil || tag.yearFrame == nil
}

func (tag *Tag) mergeTag(tag2 *Tag) {
	if tag.albumFrame == nil && tag2.albumFrame != nil {
		tag.albumFrame = tag2.albumFrame
	}
	if tag.artistFrame == nil && tag2.artistFrame != nil {
		tag.artistFrame = tag2.artistFrame
	}
	if tag.genreFrame == nil && tag2.genreFrame != nil {
		tag.genreFrame = tag2.genreFrame
	}
	if tag.titleFrame == nil && tag2.titleFrame != nil {
		tag.titleFrame = tag2.titleFrame
	}
	if tag.yearFrame == nil && tag2.yearFrame != nil {
		tag.yearFrame = tag2.yearFrame
	}
}

package id3

import "golang.org/x/text/language"

type FullTextFrame struct {
	frameBase

	language    language.Base
	description string
	text        string
}

func newFullTextFrame(tag *Tag, header *frameHeader, data []byte) (Frame, error) {
	ftf := &FullTextFrame{}
	ftf.header = header

	l := len(data)
	textEncoding, encoding, err := extractEncoding(l, data)
	if err != nil {
		return nil, err
	}
	langCode := string(data[1:4])

	if langCode == "xxx" || langCode == "XXX" {
		langCode = "ENG"
	}

	ftf.language, err = language.ParseBase(langCode)
	if err != nil {
		return nil, err
	}

	description, i, err := trimForEncoding(l-4, data[4:], textEncoding, false)
	if err != nil {
		return nil, err
	}
	ftf.description, err = decodeString(description, encoding)
	if err != nil {
		return nil, err
	}
	text, _, err := trimForEncoding(l-(i+4), data[i+4:], textEncoding, false)
	if err != nil {
		return nil, err
	}
	ftf.description, err = decodeString(text, encoding)
	if err != nil {
		return nil, err
	}
	return ftf, nil
}

func newDescribedFrame(tag *Tag, header *frameHeader, data []byte) (Frame, error) {
	ftf := &FullTextFrame{}
	ftf.header = header

	l := len(data)
	textEncoding, encoding, err := extractEncoding(l, data)
	if err != nil {
		return nil, err
	}

	description, i, err := trimForEncoding(l, data, textEncoding, true)
	if err != nil {
		return nil, err
	}
	ftf.description, err = decodeString(description, encoding)
	if err != nil {
		return nil, err
	}
	text, _, err := trimForEncoding(l-i, data[i:], textEncoding, false)
	if err != nil {
		return nil, err
	}
	ftf.description, err = decodeString(text, encoding)
	if err != nil {
		return nil, err
	}
	return ftf, nil
}

func (ftf *FullTextFrame) String() string {
	return ftf.text
}

func (ftf *FullTextFrame) Bytes() []byte {
	return []byte(ftf.text)
}

func (ftf *FullTextFrame) Description() string {
	return ftf.description
}

func (ftf *FullTextFrame) Language() language.Base {
	return ftf.language
}

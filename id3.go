package id3

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func Read(r io.ReadSeeker) (*Tag, error) {
	//glog.Infof("READING: %v", path)

	header, err := newHeader(r)
	if err != nil {
		if err == ErrNoHeader {
			return readv1(r)

		}
		return nil, err
	}

	//glog.Infof("Unsynchronized: %v", header.Unsynchronized())

	var extendedHeader *ExtendedHeader

	size := header.Size()

	if header.HasExtendedHeader() {
		extendedHeader, err = newExtendedHeader(r)
		if err != nil {
			return nil, err
		}
		size -= extendedHeader.size
	}

	tag := newTag(header, extendedHeader)

	switch header.version {
	case 2:
		tag.readV2(size, version22Params, r)
	case 3:
		tag.readV2(size, version23Params, r)
	case 4:
		tag.readV2(size, version24Params, r)
	default:
		return nil, errors.New(fmt.Sprintf("Unknown major revision: %v", header.version))
	}

	if tag.missingCoreInfo() {
		id3v1Tag, err := readv1(r)
		if err == nil {
			tag.mergeTag(id3v1Tag)
		}
	}

	return tag, nil
}

func readAll(path string, parseBoth bool) (*Tag, *Tag, error) {
	r, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, nil, err
	}

	header, err := newHeader(r)
	if err != nil {
		if err == ErrNoHeader {
			v1, err := readv1(r)
			return nil, v1, err

		}
		return nil, nil, err
	}

	//glog.Infof("Unsynchronized: %v", header.Unsynchronized())

	var extendedHeader *ExtendedHeader

	size := header.Size()

	if header.HasExtendedHeader() {
		extendedHeader, err = newExtendedHeader(r)
		if err != nil {
			return nil, nil, err
		}
		size -= extendedHeader.size
	}

	var id3v1Tag *Tag
	tag := newTag(header, extendedHeader)

	switch header.version {
	case 2:
		tag.readV2(size, version22Params, r)
	case 3:
		tag.readV2(size, version23Params, r)
	case 4:
		tag.readV2(size, version24Params, r)
	default:
		return nil, nil, errors.New(fmt.Sprintf("Unknown major revision: %v", header.version))
	}

	if parseBoth {
		id3v1Tag, err = readv1(r)
	}

	return tag, id3v1Tag, nil
}

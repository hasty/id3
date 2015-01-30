package id3

import (
	"errors"
	"fmt"
	"os"
)

func Read(path string) (*Tag, error) {
	//glog.Infof("READING: %v", path)

	r, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

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
		tag.readV2(path, size, version22Params, r)
	case 3:
		tag.readV2(path, size, version23Params, r)
	case 4:
		tag.readV2(path, size, version24Params, r)
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

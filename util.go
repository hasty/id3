package id3

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type TextEncoding byte

const (
	ISO88591 TextEncoding = iota
	UTF16
	UTF16BE
	UTF8
)

func unsafe(b []byte) uint32 {
	o := make([]byte, 4)
	o[3] = ((b[3] >> 0) & 0x7F) | ((b[2] & 0x01) << 7)
	o[2] = ((b[2] >> 1) & 0x3F) | ((b[1] & 0x03) << 6)
	o[1] = ((b[1] >> 2) & 0x1F) | ((b[0] & 0x07) << 5)
	o[0] = (b[0] >> 3) & 0x0F

	return binary.BigEndian.Uint32(o)
}

func readString(data []byte) (string, error) {
	l := len(data)
	if l < 2 {
		return "", nil
	}
	var encoding encoding.Encoding
	textEncoding := TextEncoding(data[0])

	switch textEncoding {
	case ISO88591:
		// Technically a superset of ISO-8859-1, but we're only reading so it's ok
		encoding = charmap.Windows1252
		data, _ = trimToNull(l, data)
	case UTF16:
		if len(data) < 3 {
			encoding = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
		} else if data[1] == 0xFE && data[2] == 0xFF {
			encoding = unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM)
		} else if data[1] == 0xFF && data[2] == 0xFE {
			encoding = unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM)
		} else {
			encoding = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
		}
		data, _ = trimToDoubleNull(l, data)
	case UTF16BE:
		encoding = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
		data, _ = trimToDoubleNull(l, data)
	case UTF8:
		encoding = nil
		data, _ = trimToNull(l, data)
	}
	return decodeString(data, encoding)
}

func decodeString(data []byte, encoding encoding.Encoding) (string, error) {
	if encoding != nil {
		reader := transform.NewReader(bytes.NewReader(data), encoding.NewDecoder())
		n, err := ioutil.ReadAll(reader)
		if err != nil {
			return "", err
		}
		data = n
	}
	return string(data), nil
}
func trimToNull(l int, data []byte) ([]byte, int) {
	var i int = 1
	for i < l {
		if data[i] == 0x0 {
			break
		}
		i++
	}
	return data[1:i], i
}

func trimToDoubleNull(l int, data []byte) ([]byte, int) {
	var i int = 1
	for i < l {
		if data[i] == 0x0 && data[i+1] == 0x0 {
			break
		}
		i += 2
	}
	if i >= l {
		i = l
	}
	return data[1:i], i
}

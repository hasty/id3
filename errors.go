package id3

import "errors"

var ErrTooShort = errors.New("invalid file; too short")
var ErrNoHeader = errors.New("invalid file; missing ID3 Header")
var ErrCorruptExtendedHeader = errors.New("invalid file; Extended Header is too short")

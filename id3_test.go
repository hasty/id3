// Copyright 2013 Michael Yang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package id3

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/glog"
)

type testData struct {
	name        string
	path        string
	artist      string
	title       string
	album       string
	description string
	comment     string
}

func walkFunc(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}
	//glog.Infof("Reading %v err:%v", path, err)
	_, err = Read(path)
	if err != nil {
		if err == ErrNoHeader || err == ErrTooShort {
			return nil
		}
		glog.Errorf("err:%v", err)
		return err
	}
	return nil
}

var root string

func init() {
	flag.StringVar(&root, "root", "M:\\Music", "The root to parse")
	flag.Parse()

}

func TestAll(t *testing.T) {

	filepath.Walk(root, walkFunc)
}

func DontTestParse(t *testing.T) {
	id3v2Files := []*testData{
		&testData{"Simple MP3",
			"test/test.mp3",
			"Paloalto",
			"Nice Life (Feat. Basick)",
			"Chief Life",
			"✓",
			"✓"},

		&testData{"Dummy Frames",
			"test/dummyframes.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},
		&testData{"Incomplete MPEG Frame",
			"test/incompletempegframe.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},
		&testData{"Obsolete (No Image)",
			"test/obsolete-noimage.mp3",
			"ARTISTABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"NAMEABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"ALBUMNAMEABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"",
			""},
		&testData{"Obsolete",
			"test/obsolete.mp3",
			"ARTIST1234567890123456789012345678901234567890",
			"NAME1234567890123456789012345678901234567890",
			"ALBUM1234567890123456789012345678901234567890",
			"",
			""},
		&testData{"id3v1 and id3v2.3 (Custom Tags)",
			"test/v1andv23andcustomtags.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},
		&testData{"id3v1 and id3v2.3", "test/v1andv23tags.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},

		&testData{"id3v1 and id3v2.3 (Album Image & Little-Endian UTF-16)",
			"test/v1andv23tagswithalbumimage-utf16le.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"ID3v1 Comment",
			"COMMENT123456789012345678901"},

		&testData{"id3v1 and id3v2.3 (Album Image)",
			"test/v1andv23tagswithalbumimage.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},
		&testData{"id3v1 and id3v2.4",
			"test/v1andv24tags.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},

		&testData{"id3v2.3",
			"test/v23tag.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},

		/*&testData{"id3v2.3 (Chapters)",
		"test/v23tagwithchapters.mp3",
		"Paloalto",
		"Nice Life (Feat. Basick)",
		"Chief Life",
		"✓",
		"✓"},*/

		&testData{"id3v2.3 (Unicode)",
			"test/v23unicodetags.mp3",
			"γειά σου",
			"中文",
			"こんにちは",
			"",
			""},

		&testData{"id3v2.4 (Album Image)",
			"test/v24tagswithalbumimage.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},
		&testData{"iTunes Comment",
			"test/withitunescomment.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},
	}
	for _, v := range id3v2Files {
		err := testReadV2(v, t)
		if err != nil {
			t.Errorf("%v: %v", v.name, err)
		}
	}
	id3v1Files := []testData{
		testData{"id3v1",
			"test/v1tag.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901\x00\x01"},
		testData{"id3v1 (No Track)",
			"test/v1tagwithnotrack.mp3",
			"ARTIST123456789012345678901234",
			"TITLE1234567890123456789012345",
			"ALBUM1234567890123456789012345",
			"",
			"COMMENT123456789012345678901"},
	}
	for _, v := range id3v1Files {
		_, err := Read(v.path)
		if err != nil {
			t.Errorf("%v: %v", v.name, err)
		}
	}
	invalidFiles := []testData{
		testData{"No Tags",
			"test/notags.mp3",
			"",
			"",
			"",
			"",
			""},
		testData{"Not An MP3",
			"test/notanmp3.mp3",
			"",
			"",
			"",
			"",
			""},
	}
	for _, v := range invalidFiles {
		_, err := Read(v.path)
		if err != ErrNoHeader {
			t.Errorf("%v: %v", v.name, err)
		}
	}
}

func testReadV2(testData *testData, t *testing.T) error {
	glog.Infof("Reading %v...", testData.path)
	tag, err := Read(testData.path)
	if err != nil {
		return err
	}

	if s := tag.Artist(); s != testData.artist {
		return errors.New(fmt.Sprintf("incorrect artist, \"%v\", expected \"%v\"", s, testData.artist))
	}

	if s := tag.Title(); s != testData.title {
		return errors.New(fmt.Sprintf("incorrect title, \"%v\", expected \"%v\"", s, testData.title))
	}

	if s := tag.Album(); s != testData.album {
		return errors.New(fmt.Sprintf("incorrect album, \"%v\", expected \"%v\"", s, testData.album))
	}

	/*parsedFrame := file.Frame("COMM")
	if parsedFrame == nil {
		if testData.comment != "" {
			return errors.New(fmt.Sprintf("missing comment, expected \"%v\"", testData.comment))
		}
		return nil
	}
	resultFrame, ok := parsedFrame.(*v2.UnsynchTextFrame)
	if !ok {
		return errors.New("couldn't cast frame")
	}

	actual := resultFrame.Description()

	if testData.description != actual {
		return errors.New(fmt.Sprintf("incorrect description, \"%v\", expected \"%v\"", actual, testData.description))
	}

	actual = resultFrame.Text()
	if testData.comment != actual {
		return errors.New(fmt.Sprintf("incorrect comment, \"%v\", expected \"%v\"", actual, testData.comment))
	}*/
	return nil
}

/*func testReadV2(testData *testData, t *testing.T) error {
	fmt.Println(testData.path)
	file, err := Read(testData.path)
	if err != nil {
		return err
	}

	tag, ok := file.Tagger.(*v2.Tag)
	if !ok {
		return errors.New("incorrect tagger type")
	}

	if s := tag.Artist(); s != testData.artist {
		return errors.New(fmt.Sprintf("incorrect artist, \"%v\", expected \"%v\"", testData.name, s, testData.artist))
	}

	if s := tag.Title(); s != testData.title {
		return errors.New(fmt.Sprintf("incorrect title, \"%v\", expected \"%v\"", s, testData.title))
	}

	if s := tag.Album(); s != testData.album {
		return errors.New(fmt.Sprintf("incorrect album, \"%v\", expected \"%v\"", s, testData.album))
	}

	parsedFrame := file.Frame("COMM")
	if parsedFrame == nil {
		if testData.comment != "" {
			return errors.New(fmt.Sprintf("missing comment, expected \"%v\"", testData.comment))
		}
		return nil
	}
	resultFrame, ok := parsedFrame.(*v2.UnsynchTextFrame)
	if !ok {
		return errors.New("couldn't cast frame")
	}

	actual := resultFrame.Description()

	if testData.description != actual {
		return errors.New(fmt.Sprintf("incorrect description, \"%v\", expected \"%v\"", actual, testData.description))
	}

	actual = resultFrame.Text()
	if testData.comment != actual {
		return errors.New(fmt.Sprintf("incorrect comment, \"%v\", expected \"%v\"", actual, testData.comment))
	}
	return nil
}

func testReadV1(testData *testData, t *testing.T) error {
	fmt.Println(testData.path)
	file, err := Read(testData.path)
	if err != nil {
		return err
	}

	tag, ok := file.Tagger.(*v1.Tag)
	if !ok {
		return errors.New("incorrect tagger type")
	}

	if s := tag.Artist(); s != testData.artist {
		return errors.New(fmt.Sprintf("incorrect artist, \"%v\", expected \"%v\"", testData.name, s, testData.artist))
	}

	if s := tag.Title(); s != testData.title {
		return errors.New(fmt.Sprintf("incorrect title, \"%v\", expected \"%v\"", s, testData.title))
	}

	if s := tag.Album(); s != testData.album {
		return errors.New(fmt.Sprintf("incorrect album, \"%v\", expected \"%v\"", s, testData.album))
	}

	if testData.comment != "" {
		var found bool
		comments := tag.Comments()
		for _, comment := range comments {
			if strings.TrimSpace(comment) == testData.comment {
				found = true
				break
			}
		}
		if !found {
			return errors.New(fmt.Sprintf("missing comment, expected \"%v\"", testData.comment))
		}
	}
	return nil
}

func TestOpen(t *testing.T) {
	file, err := Open(testFile)
	if err != nil {
		t.Errorf("Open: unable to open file")
		return
	}

	tag, ok := file.Tagger.(*v2.Tag)
	if !ok {
		t.Errorf("Open: incorrect tagger type")
		return
	}

	if s := tag.Artist(); s != "Paloalto" {
		t.Errorf("Open: incorrect artist, %v", s)
		return
	}

	if s := tag.Title(); s != "Nice Life (Feat. Basick)" {
		t.Errorf("Open: incorrect title, %v", s)
		return
	}

	if s := tag.Album(); s != "Chief Life" {
		t.Errorf("Open: incorrect album, %v", s)
		return
	}

	parsedFrame := file.Frame("COMM")
	resultFrame, ok := parsedFrame.(*v2.UnsynchTextFrame)
	if !ok {
		t.Error("Couldn't cast frame")
		return
	}

	expected := "✓"
	actual := resultFrame.Description()

	if expected != actual {
		t.Errorf("Expected %x, got %x", expected, actual)
		return
	}

	actual = resultFrame.Text()
	if expected != actual {
		t.Errorf("Expected %q, got %q", expected, actual)
		return
	}
}

func TestClose(t *testing.T) {
	before, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("test file error")
	}

	file, err := Open(testFile)
	if err != nil {
		t.Errorf("Close: unable to open file")
	}
	beforeCutoff := file.originalSize

	file.SetArtist("Paloalto")
	file.SetTitle("Test test test test test test")

	afterCutoff := file.Size()

	if err := file.Close(); err != nil {
		t.Errorf("Close: unable to close file")
	}

	after, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("Close: unable to reopen file")
	}

	if !bytes.Equal(before[beforeCutoff:], after[afterCutoff:]) {
		t.Errorf("Close: nontag data lost on close")
	}

	if err := ioutil.WriteFile(testFile, before, 0666); err != nil {
		t.Errorf("Close: unable to write original contents to test file")
	}
}

func TestReadonly(t *testing.T) {
	before, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("test file error")
	}

	file, err := Open(testFile)
	if err != nil {
		t.Errorf("Readonly: unable to open file")
	}

	file.Title()
	file.Artist()
	file.Album()
	file.Year()
	file.Genre()
	file.Comments()

	if err := file.Close(); err != nil {
		t.Errorf("Readonly: unable to close file")
	}

	after, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("Readonly: unable to reopen file")
	}

	if !bytes.Equal(before, after) {
		t.Errorf("Readonly: tag data modified without set")
	}
}

func TestAddTag(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "notag")
	if err != nil {
		t.Fatal(err)
	}

	file, err := Open(tempFile.Name())
	if err != nil {
		t.Errorf("AddTag: unable to open empty file")
	}

	tag := file.Tagger

	if tag == nil {
		t.Errorf("AddTag: no tag added to file")
	}

	file.SetArtist("Michael")

	err = file.Close()
	if err != nil {
		t.Errorf("AddTag: error closing new file")
	}

	reopenBytes, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Errorf("AddTag: error reopening file")
	}

	expectedBytes := tag.Bytes()
	if !bytes.Equal(expectedBytes, reopenBytes) {
		t.Errorf("AddTag: tag not written correctly: %v", reopenBytes)
	}
}

func TestUnsynchTextFrame_RoundTrip(t *testing.T) {
	var (
		err              error
		tempfile         *os.File
		f                *File
		tagger           *v2.Tag
		ft               v2.FrameType
		utextFrame       *v2.UnsynchTextFrame
		parsedFrame      v2.Framer
		resultFrame      *v2.UnsynchTextFrame
		ok               bool
		expected, actual string
	)

	tempfile, err = ioutil.TempFile("", "id3v2")
	if err != nil {
		t.Fatal(err)
	}

	tagger = v2.NewTag(3)
	ft = v2.V23FrameTypeMap["COMM"]
	utextFrame = v2.NewUnsynchTextFrame(ft, "Comment", "Foo")
	tagger.AddFrames(utextFrame)

	_, err = tempfile.Write(tagger.Bytes())
	tempfile.Close()
	if err != nil {
		t.Fatal(err)
	}

	f, err = Open(tempfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	parsedFrame = f.Frame("COMM")
	if resultFrame, ok = parsedFrame.(*v2.UnsynchTextFrame); !ok {
		t.Error("Couldn't cast frame")
	} else {
		expected = utextFrame.Description()
		actual = resultFrame.Description()
		if expected != actual {
			t.Errorf("Expected %q, got %q", expected, actual)
		}
	}
}

func TestUTF16CommPanic(t *testing.T) {
	osFile, err := os.Open(testFile)
	if err != nil {
		t.Error(err)
	}
	tempfile, err := ioutil.TempFile("", "utf16_comm")
	if err != nil {
		t.Error(err)
	}
	io.Copy(tempfile, osFile)
	osFile.Close()
	tempfile.Close()
	for i := 0; i < 2; i++ {
		file, err := Open(tempfile.Name())
		if err != nil {
			t.Error(err)
		}
		file.Close()
	}
}*/

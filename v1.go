package id3

import (
	"errors"
	"io"
	"os"
	"strings"
)

const (
	v1TagSize = 128
)

func readv1(r io.ReadSeeker) (*Tag, error) {
	_, err := r.Seek(-v1TagSize, os.SEEK_END)
	if err != nil {
		return nil, err
	}

	data := make([]byte, v1TagSize)

	n, err := io.ReadFull(r, data)
	if err != nil {
		return nil, err
	}
	if n < v1TagSize {
		return nil, ErrTooShort
	}
	if string(data[:3]) != "TAG" {
		return nil, ErrNoHeader
	}

	tag := emptyTag()

	title := strings.TrimRight(string(data[3:33]), "\x00")
	artist := strings.TrimRight(string(data[33:63]), "\x00")
	album := strings.TrimRight(string(data[63:93]), "\x00")
	year := strings.TrimRight(string(data[93:97]), "\x00")
	comment := strings.TrimRight(string(data[97:127]), "\x00")

	genreByte := int(data[127])

	if genreByte >= len(v1Genres) {
		return nil, errors.New("Unknown v1 genre")
	}
	genre := v1Genres[genreByte]

	tag.addFrame(simpleTextFrame(tag, "TT2", title))
	tag.addFrame(simpleTextFrame(tag, "TP1", artist))
	tag.addFrame(simpleTextFrame(tag, "TAL", album))
	tag.addFrame(simpleTextFrame(tag, "TYE", year))
	tag.addFrame(simpleTextFrame(tag, "TCO", genre))
	tag.addFrame(simpleTextFrame(tag, "COM", comment))

	return tag, nil
}

var (
	v1Genres = []string{
		"Blues",
		"Classic Rock",
		"Country",
		"Dance",
		"Disco",
		"Funk",
		"Grunge",
		"Hip-Hop",
		"Jazz",
		"Metal",
		"New Age",
		"Oldies",
		"Other",
		"Pop",
		"R&B",
		"Rap",
		"Reggae",
		"Rock",
		"Techno",
		"Industrial",
		"Alternative",
		"Ska",
		"Death Metal",
		"Pranks",
		"Soundtrack",
		"Euro-Techno",
		"Ambient",
		"Trip-Hop",
		"Vocal",
		"Jazz+Funk",
		"Fusion",
		"Trance",
		"Classical",
		"Instrumental",
		"Acid",
		"House",
		"Game",
		"Sound Clip",
		"Gospel",
		"Noise",
		"AlternRock",
		"Bass",
		"Soul",
		"Punk",
		"Space",
		"Meditative",
		"Instrumental Pop",
		"Instrumental Rock",
		"Ethnic",
		"Gothic",
		"Darkwave",
		"Techno-Industrial",
		"Electronic",
		"Pop-Folk",
		"Eurodance",
		"Dream",
		"Southern Rock",
		"Comedy",
		"Cult",
		"Gangsta",
		"Top 40",
		"Christian Rap",
		"Pop/Funk",
		"Jungle",
		"Native American",
		"Cabaret",
		"New Wave",
		"Psychadelic",
		"Rave",
		"Showtunes",
		"Trailer",
		"Lo-Fi",
		"Tribal",
		"Acid Punk",
		"Acid Jazz",
		"Polka",
		"Retro",
		"Musical",
		"Rock & Roll",
		"Hard Rock",
		"Folk",
		"Folk-Rock",
		"National Folk",
		"Swing",
		"Fast Fusion",
		"Bebop",
		"Latin",
		"Revival",
		"Celtic",
		"Bluegrass",
		"Avantgarde",
		"Gothic Rock",
		"Progressive Rock",
		"Psychedelic Rock",
		"Symphonic Rock",
		"Slow Rock",
		"Big Band",
		"Chorus",
		"Easy Listening",
		"Acoustic",
		"Humour",
		"Speech",
		"Chanson",
		"Opera",
		"Chamber Music",
		"Sonata",
		"Symphony",
		"Booty Bass",
		"Primus",
		"Porn groove",
		"Satire",
		"Slow Jam",
		"Club",
		"Tango",
		"Samba",
		"Folklore",
		"Ballad",
		"Power Ballad",
		"Rhythmic Soul",
		"Freestyle",
		"Duet",
		"Punk rock",
		"Drum Solo",
		"A capella",
		"Euro-House",
		"Dance Hall",
		"Goa Trance",
		"Drum & Bass",
		"Club-House",
		"Hardcore Techno",
		"Terror",
		"Indie",
		"BritPop",
		"Afro-punk",
		"Polsk Punk",
		"Beat",
		"Christian Gangsta Rap",
		"Heavy Metal",
		"Black Metal",
		"Crossover",
		"Contemporary Christian",
		"Christian Rock",
		"Merengue",
		"Salsa",
		"Thrash Metal",
		"Anime",
		"Jpop",
		"Synthpop",
		"Abstract",
		"Art Rock",
		"Baroque",
		"Bhangra",
		"Big Beat",
		"Breakbeat",
		"Chillout",
		"Downtempo",
		"Dub",
		"EBM",
		"Eclectic",
		"Electro",
		"Electroclash",
		"Emo",
		"Experimental",
		"Garage",
		"Global",
		"IDM",
		"Illbient",
		"Industro-Goth",
		"Jam Band",
		"Krautrock",
		"Leftfield",
		"Lounge",
		"Math Rock",
		"New Romantic",
		"Nu-Breakz",
		"Post-Punk",
		"Post-Rock",
		"Psytrance",
		"Shoegaze",
		"Space Rock",
		"Trop Rock",
		"World Music",
		"Neoclassical",
		"Audiobook",
		"Audio Theatre",
		"Neue Deutsche Welle",
		"Podcast",
		"Indie Rock",
		"G-Funk",
		"Dubstep",
		"Garage Rock",
		"Psybient",
	}
)

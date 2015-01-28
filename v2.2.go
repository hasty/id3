package id3

var version22Params = &versionParams{
	frameIdSize:    3,
	frameSizeSize:  3,
	frameFlagsSize: 0,
	frames: map[string]*frameFactory{
		"BUF": &frameFactory{description: "Recommended buffer size", maker: newDataFrame},
		"CNT": &frameFactory{description: "Play counter", maker: newDataFrame},
		// TODO: "COM": &frameFactory{description: "Comments", maker: ParseUnsynchTextFrame},
		"CRA": &frameFactory{description: "Audio encryption", maker: newDataFrame},
		"CRM": &frameFactory{description: "Encrypted meta frame", maker: newDataFrame},
		"ETC": &frameFactory{description: "Event timing codes", maker: newDataFrame},
		"EQU": &frameFactory{description: "Equalization", maker: newDataFrame},
		"GEO": &frameFactory{description: "General encapsulated object", maker: newDataFrame},
		"IPL": &frameFactory{description: "Involved people list", maker: newDataFrame},
		"LNK": &frameFactory{description: "Linked information", maker: newDataFrame},
		"MCI": &frameFactory{description: "Music CD Identifier", maker: newDataFrame},
		"MLL": &frameFactory{description: "MPEG location lookup table", maker: newDataFrame},
		"PIC": &frameFactory{description: "Attached picture", maker: newDataFrame},
		"POP": &frameFactory{description: "Popularimeter", maker: newDataFrame},
		"REV": &frameFactory{description: "Reverb", maker: newDataFrame},
		"RVA": &frameFactory{description: "Relative volume adjustment", maker: newDataFrame},
		"SLT": &frameFactory{description: "Synchronized lyric/text", maker: newDataFrame},
		"STC": &frameFactory{description: "Synced tempo codes", maker: newDataFrame},
		"TAL": &frameFactory{description: "Album/Movie/Show title", maker: newTextFrame},
		"TBP": &frameFactory{description: "BPM (Beats Per Minute)", maker: newTextFrame},
		"TCM": &frameFactory{description: "Composer", maker: newTextFrame},
		"TCO": &frameFactory{description: "Content type", maker: newTextFrame},
		"TCR": &frameFactory{description: "Copyright message", maker: newTextFrame},
		"TDA": &frameFactory{description: "Date", maker: newTextFrame},
		"TDY": &frameFactory{description: "Playlist delay", maker: newTextFrame},
		"TEN": &frameFactory{description: "Encoded by", maker: newTextFrame},
		"TFT": &frameFactory{description: "File type", maker: newTextFrame},
		"TIM": &frameFactory{description: "Time", maker: newTextFrame},
		"TKE": &frameFactory{description: "Initial key", maker: newTextFrame},
		"TLA": &frameFactory{description: "Language(s)", maker: newTextFrame},
		"TLE": &frameFactory{description: "Length", maker: newTextFrame},
		"TMT": &frameFactory{description: "Media type", maker: newTextFrame},
		"TOA": &frameFactory{description: "Original artist(s)/performer(s)", maker: newTextFrame},
		"TOF": &frameFactory{description: "Original filename", maker: newTextFrame},
		"TOL": &frameFactory{description: "Original Lyricist(s)/text writer(s)", maker: newTextFrame},
		"TOR": &frameFactory{description: "Original release year", maker: newTextFrame},
		"TOT": &frameFactory{description: "Original album/Movie/Show title", maker: newTextFrame},
		"TP1": &frameFactory{description: "Lead artist(s)/Lead performer(s)/Soloist(s)/Performing group", maker: newTextFrame},
		"TP2": &frameFactory{description: "Band/Orchestra/Accompaniment", maker: newTextFrame},
		"TP3": &frameFactory{description: "Conductor/Performer refinement", maker: newTextFrame},
		"TP4": &frameFactory{description: "Interpreted, remixed, or otherwise modified by", maker: newTextFrame},
		"TPA": &frameFactory{description: "Part of a set", maker: newTextFrame},
		"TPB": &frameFactory{description: "Publisher", maker: newTextFrame},
		"TRC": &frameFactory{description: "ISRC (International Standard Recording Code)", maker: newTextFrame},
		"TRD": &frameFactory{description: "Recording dates", maker: newTextFrame},
		"TRK": &frameFactory{description: "Track number/Position in set", maker: newTextFrame},
		"TSI": &frameFactory{description: "Size", maker: newTextFrame},
		"TSS": &frameFactory{description: "Software/hardware and settings used for encoding", maker: newTextFrame},
		"TT1": &frameFactory{description: "Content group description", maker: newTextFrame},
		"TT2": &frameFactory{description: "Title/Songname/Content description", maker: newTextFrame},
		"TT3": &frameFactory{description: "Subtitle/Description refinement", maker: newTextFrame},
		"TXT": &frameFactory{description: "Lyricist/text writer", maker: newTextFrame},
		// TODO: "TXX": &frameFactory{description: "User defined text information frame", maker: ParseDescTextFrame},
		"TYE": &frameFactory{description: "Year", maker: newTextFrame},
		"UFI": &frameFactory{description: "Unique file identifier", maker: newDataFrame},
		"ULT": &frameFactory{description: "Unsychronized lyric/text transcription", maker: newDataFrame},
		"WAF": &frameFactory{description: "Official audio file webpage", maker: newDataFrame},
		"WAR": &frameFactory{description: "Official artist/performer webpage", maker: newDataFrame},
		"WAS": &frameFactory{description: "Official audio source webpage", maker: newDataFrame},
		"WCM": &frameFactory{description: "Commercial information", maker: newDataFrame},
		"WCP": &frameFactory{description: "Copyright/Legal information", maker: newDataFrame},
		"WPB": &frameFactory{description: "Publishers official webpage", maker: newDataFrame},
		"WXX": &frameFactory{description: "User defined URL link frame", maker: newDataFrame},
	},
}
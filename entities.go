package seratoparser

var seratoVolume string

// MediaEntity defines Tracks/Songs entities in a Database or Crate
//aka Otrk
type MediaEntity struct {
	// META
	DVOL string // volume

	// UTFSTR
	PTRK string // filetrack
	PFIL string // filebase
	TTYP string // track type
	TSNG string // song title
	TART string // artist Title
	TGEN string // genre
	TLEN string // length
	TBIT string // bit rate
	TSMP string // sampling rate
	TALB string // album title
	TCOM string // comment
	TCMP string // composer
	TADD string // local time added to database
	TKEY string // key
	TLBL string // ??
	TBPM string // bpm
	TGRP string // group?
	TTYR string // year
	TSIZ string // size of track
	UDSC string // desc?
	TRMX string // remix tag
	TCOR string // group?
	PVID string // group?

	// INT1
	BMIS bool // missing
	BCRT bool // corrupt
	BHRT bool // ?
	BPLY bool // ?
	BLOP bool // ?
	BITU bool // ?
	BOVC bool // ?
	BIRO bool // ?
	BWLB bool // ?
	BWLL bool // ?
	BUNS bool // ?
	BBGL bool // ?
	BKRK bool // ?

	// INT4
	UADD int // timeadded
	UTKN int // ??
	UTME int // ??
	UTPC int // playcount
	UFSB int // ??

	// BYTE SLICE
	ULBL []byte // color - track colour
	SBAV []byte //?
}

// HistoryEntity defines Tracks/Songs entities in a History Session
type HistoryEntity struct {
	RROW int    // rrow
	RDIR string // rfullpath
	TTMS int    // rstarttime
	TTME int    // rendtime
	TDCK int    // rdeck
	RDTE string // rdate*
	RSRT int    // rstart*
	REND int    // rend*
	TPTM int    // rplaytime
	RSES int    // rsessionId
	RPLY int    // rplayed = 1
	RADD int    // radded
	RUPD int    // rupdatedAt
	RSWR string // rsoftware*
	RSWB int    // rsoftwareBuild*
	RDEV string // rdevice
}

// SeratoAdatMap Defines all the known keys with their integer key found in Serato Databases
// TODO: Identify all fields of an ADAT object
var SeratoAdatMap = map[int]string{
	1:  "RROW", // rrow
	2:  "RDIR", // rfullpath
	28: "TTMS", // rstarttime
	29: "TTME", // rendtime
	31: "TDCK", // rdeck

	41: "RDTE", // rdate
	43: "RSRT", // rstart
	44: "REND", // rend

	45: "TPTM", // rplaytime
	48: "RSES", // rsessionId
	50: "RPLY", // rplayed
	52: "RADD", // radded
	53: "RUPD", // rupdatedAt
	54: "RUNK", // rr54unknownTimestamp

	57: "RSWR", // rsoftware
	58: "RSWB", // rsoftwareBuild

	63: "RDEV", // rdevice
}

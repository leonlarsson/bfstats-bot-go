package shapes

// BaseData represents the base data for all image creations.
// Each game embeds BaseData together with its own GAMEStats struct.
type BaseData struct {
	Identifier string `json:"identifier"`
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
	Platform   int    `json:"platform"`
	TimePlayed string `json:"timePlayed"`
	Meta       Meta   `json:"meta"`
}

// Meta represents the meta data for the image.
type Meta struct {
	Game    string `json:"game"`
	Segment string `json:"segment"`
}

// Stat represents a single stat with a name, value and extra information.
type Stat struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Extra string `json:"extra"`
}

// RankStat represents a rank stat with a name, value, rankInt and extra information. Currently replaced by Stat.
type RankStat struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	RankInt int    `json:"rankInt"`
	Extra   string `json:"extra"`
}

// GenericRegularData represents the data for a regular image.
type GenericRegularData struct {
	BaseData BaseData
	Stats    GenericRegularStats
}

// GenericRegularStats represents the stats for the overview segment.
type GenericRegularStats struct {
	L1,
	L2,
	L3,
	L4,
	L5,
	L6,
	R1,
	R2,
	R3,
	R4 Stat
}

// GenericGridData represents the data for a grid image.
type GenericGridData struct {
	BaseData BaseData
	Entries  []Stat
}

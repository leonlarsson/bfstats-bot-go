package canvas

// BaseData represents the base data for all image creations.
// Each game embeds BaseData together with its own GAMEStats struct.
type BaseData struct {
	Identifier string `json:"identifier"`
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
	Platform   int    `json:"platform"`
	Meta       Meta   `json:"meta"`
}

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

// RankStat represents a rank stat with a name, value, rankInt and extra information.
type RankStat struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	RankInt int    `json:"rankInt"`
	Extra   string `json:"extra"`
}

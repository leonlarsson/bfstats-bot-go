package canvas

// BF2042OverviewCanvasData represents the data needed to create a Battlefield 2042 image.
type BF2042OverviewCanvasData struct {
	BaseData `json:"baseData"`
	Stats    BF2042OverviewCanvasStats `json:"stats"`
}

// BF2042OverviewCanvasStats represents the stats for Battlefield 2042.
type BF2042OverviewCanvasStats struct {
	TimePlayed     Stat     `json:"timePlayed"`
	Kills          Stat     `json:"kills"`
	Deaths         Stat     `json:"deaths"`
	Assists        Stat     `json:"assists"`
	Revives        Stat     `json:"revives"`
	WlRatio        Stat     `json:"wlRatio"`
	BestClass      Stat     `json:"bestClass"`
	KillsPerMatch  Stat     `json:"killsPerMatch"`
	KdRatio        Stat     `json:"kdRatio"`
	KillsPerMinute Stat     `json:"killsPerMinute"`
	MultiKills     Stat     `json:"multiKills"`
	Rank           RankStat `json:"rank"`
}

// BF2042WeaponsCanvasData represents the data needed to create a Battlefield 2042 image.
type BF2042WeaponsCanvasData struct {
	BaseData `json:"baseData"`
	Weapons  []Stat `json:"weapons"`
}

// BF2042VehiclesCanvasData represents the data needed to create a Battlefield 2042 image.
type BF2042VehiclesCanvasData struct {
	BaseData `json:"baseData"`
	Vehicles []Stat `json:"vehicles"`
}

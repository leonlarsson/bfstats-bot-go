package structs

// BF2042OverviewData represents the data needed to create a Battlefield 2042 image.
type BF2042OverviewData struct {
	BaseData `json:"baseData"`
	Stats    BF2042OverviewStats `json:"stats"`
}

// BF2042OverviewStats represents the stats for Battlefield 2042.
type BF2042OverviewStats struct {
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
	Rank           RankStat `json:"rank"`
}

// BF2042WeaponsData represents the data needed to create a Battlefield 2042 image.
type BF2042WeaponsData struct {
	BaseData `json:"baseData"`
	Weapons  []Stat `json:"weapons"`
}

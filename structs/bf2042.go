package structs

// BF2042Data represents the data needed to create a Battlefield 2042 image.
type BF2042Data struct {
	BaseData `json:"baseData"`
	Stats    BF2042Stats `json:"stats"`
}

// BF2042Stats represents the stats for Battlefield 2042.
type BF2042Stats struct {
	Kills          Stat `json:"kills"`
	Deaths         Stat `json:"deaths"`
	Assists        Stat `json:"assists"`
	Revives        Stat `json:"revives"`
	WlRatio        Stat `json:"wlRatio"`
	BestClass      Stat `json:"bestClass"`
	KillsPerMatch  Stat `json:"killsPerMatch"`
	KdRatio        Stat `json:"kdRatio"`
	KillsPerMinute Stat `json:"killsPerMinute"`
	Rank           Stat `json:"rank"`
}

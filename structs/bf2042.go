package structs

type BF2042Data struct {
	Identifier string      `json:"identifier"`
	Username   string      `json:"username"`
	TimePlayed string      `json:"timePlayed"`
	Platform   int         `json:"platform"`
	Stats      BF2042Stats `json:"stats"`
}

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

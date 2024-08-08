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
	MultiKills     Stat     `json:"multiKills"`
	Rank           RankStat `json:"rank"`
}

// BF2042WeaponsData represents the data needed to create a Battlefield 2042 image.
type BF2042WeaponsData struct {
	BaseData `json:"baseData"`
	Weapons  []Stat `json:"weapons"`
}

/* TRN RESPONSES */

// TRNBF2042OverviewResponse represents the response from the TRN API for Battlefield 2042.
type TRNBF2042OverviewResponse struct {
	Data struct {
		PlatformInfo struct {
			PLatformSlug       string
			PlatformUserHandle string
			AvatarURL          string
		}
		Segments []TRNBF2042Segment
	}
}

// TRNBF2042Segment represents a segment in the TRN API response for Battlefield 2042.
type TRNBF2042Segment struct {
	Type  string
	Stats TRNBF2042OverviewStats
}

type TRNBF2042BasicIntStat struct {
	Value        int
	DisplayValue string
	Percentile   float64
}

type TRNBF2042BasicFloatStat struct {
	Value        float64
	DisplayValue string
	Percentile   float64
}

// TRNBF2042OverviewStats represents the stats in the overview segment in the TRN API response for Battlefield 2042.
type TRNBF2042OverviewStats struct {
	TimePlayed       TRNBF2042BasicIntStat
	Kills            TRNBF2042BasicIntStat
	Deaths           TRNBF2042BasicIntStat
	Assists          TRNBF2042BasicIntStat
	Revives          TRNBF2042BasicIntStat
	WlPercentage     TRNBF2042BasicFloatStat
	BestClass        TRNBF2042BasicIntStat
	KillsPerMatch    TRNBF2042BasicFloatStat
	KdRatio          TRNBF2042BasicFloatStat
	KillsPerMinute   TRNBF2042BasicFloatStat
	MultiKills       TRNBF2042BasicIntStat
	Level            TRNBF2042BasicIntStat
	LevelProgression TRNBF2042BasicFloatStat
	XPAll            TRNBF2042BasicIntStat
}

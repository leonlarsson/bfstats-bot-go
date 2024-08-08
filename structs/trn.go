package structs

type TRNResponse struct {
	Data struct {
		PlatformInfo struct {
			PLatformSlug       string
			PlatformUserHandle string
		}
		Segments []TRNBF2042Segment
	}
}

type TRNBF2042Stats struct {
	TimePlayed struct {
		Value        int
		DisplayValue string
		Percentile   float64
	}
	Kills struct {
		Value        int
		DisplayValue string
		Percentile   float64
	}
}

type TRNBF2042Segment struct {
	Type  string
	Stats TRNBF2042Stats
}

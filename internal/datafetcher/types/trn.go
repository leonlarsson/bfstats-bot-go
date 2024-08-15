package types

// trnBasicStat is a basic stat struct for the TRN api.
type trnBasicStat[T any] struct {
	Value        T
	DisplayValue string
	Percentile   float64
}

type TrnOverviewResponse struct {
	Data struct {
		PlatformInfo struct {
			PLatformSlug       string
			PlatformUserHandle string
			AvatarURL          string
		}
		Segments []struct {
			Type  string
			Stats struct {
				TimePlayed       trnBasicStat[int]
				Kills            trnBasicStat[int]
				Deaths           trnBasicStat[int]
				Assists          trnBasicStat[int]
				Revives          trnBasicStat[int]
				WlPercentage     trnBasicStat[float64]
				BestClass        trnBasicStat[int]
				KillsPerMatch    trnBasicStat[float64]
				KdRatio          trnBasicStat[float64]
				HumanKdRatio     trnBasicStat[float64]
				KillsPerMinute   trnBasicStat[float64]
				MultiKills       trnBasicStat[int]
				Level            trnBasicStat[int]
				LevelProgression trnBasicStat[float64]
				XPAll            trnBasicStat[int]
			}
		}
	}
}

type TrnWeaponsResponse struct {
	Data []struct {
		Type       string
		Attributes struct {
			Weapon      string
			CategoryKey string
		}
		Metadata struct {
			Name     string
			Category string
		}
		Stats struct {
			TimePlayed     trnBasicStat[int]
			Kills          trnBasicStat[int]
			ShotsAccuracy  trnBasicStat[float64]
			KillsPerMinute trnBasicStat[float64]
		}
	}
}

type TrnVehiclesResponse struct {
	Data []struct {
		Type       string
		Attributes struct {
			Vehicle     string
			CategoryKey string
		}
		Metadata struct {
			Name     string
			Category string
		}
		Stats struct {
			TimePlayed     trnBasicStat[int]
			Kills          trnBasicStat[int]
			KillsPerMinute trnBasicStat[float64]
		}
	}
}

type TrnClassesResponse struct {
	Data []struct {
		Type       string
		Attributes struct {
			Soldier     string
			CategoryKey string
			Mode        string
		}
		Metadata struct {
			Name     string
			Category string
			ImageURL string
		}
		Stats struct {
			TimePlayed trnBasicStat[int]
			Kills      trnBasicStat[int]
		}
	}
}

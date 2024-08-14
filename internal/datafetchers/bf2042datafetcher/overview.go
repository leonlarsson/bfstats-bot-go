package bf2042datafetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/leonlarsson/bfstats-bot-go/internal/utils"
)

type trnOverviewResponse struct {
	Data struct {
		PlatformInfo struct {
			PLatformSlug       string
			PlatformUserHandle string
			AvatarURL          string
		}
		Segments []struct {
			Type  string
			Stats struct {
				TimePlayed       trnOverviewBasicStat[int]
				Kills            trnOverviewBasicStat[int]
				Deaths           trnOverviewBasicStat[int]
				Assists          trnOverviewBasicStat[int]
				Revives          trnOverviewBasicStat[int]
				WlPercentage     trnOverviewBasicStat[float64]
				BestClass        trnOverviewBasicStat[int]
				KillsPerMatch    trnOverviewBasicStat[float64]
				KdRatio          trnOverviewBasicStat[float64]
				HumanKdRatio     trnOverviewBasicStat[float64]
				KillsPerMinute   trnOverviewBasicStat[float64]
				MultiKills       trnOverviewBasicStat[int]
				Level            trnOverviewBasicStat[int]
				LevelProgression trnOverviewBasicStat[float64]
				XPAll            trnOverviewBasicStat[int]
			}
		}
	}
}

type trnOverviewBasicStat[T any] struct {
	Value        T
	DisplayValue string
	Percentile   float64
}

func FetchBF2042OverviewData(platform, username string) (trnOverviewResponse, error) {
	var trnResponse trnOverviewResponse

	url := utils.TRNBF2042OverviewURL(platform, username)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("TRN-Internal-Api-Key", os.Getenv("TRN_API_KEY"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return trnResponse, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return trnResponse, fmt.Errorf("API returned a non-200 status code: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&trnResponse); err != nil {
		return trnResponse, err
	}

	return trnResponse, nil
}

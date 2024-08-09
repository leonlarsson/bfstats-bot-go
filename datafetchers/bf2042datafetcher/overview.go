package bf2042datafetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/leonlarsson/bfstats-bot-go/utils"
)

// TRNBF2042OverviewResponse represents the response from the TRN API for Battlefield 2042.
type TRNOverviewResponse struct {
	Data struct {
		PlatformInfo struct {
			PLatformSlug       string
			PlatformUserHandle string
			AvatarURL          *string
		}
		Segments []TRNOverviewSegment
	}
}

// TRNOverviewSegment represents a segment in the TRN API response for Battlefield 2042.
type TRNOverviewSegment struct {
	Type  string
	Stats TRNOverviewSegmentStats
}

// TRNBF2042OverviewStats represents the stats in the overview segment in the TRN API response for Battlefield 2042.
type TRNOverviewSegmentStats struct {
	TimePlayed       TRNOverviewBasicStat[int]
	Kills            TRNOverviewBasicStat[int]
	Deaths           TRNOverviewBasicStat[int]
	Assists          TRNOverviewBasicStat[int]
	Revives          TRNOverviewBasicStat[int]
	WlPercentage     TRNOverviewBasicStat[float64]
	BestClass        TRNOverviewBasicStat[int]
	KillsPerMatch    TRNOverviewBasicStat[float64]
	KdRatio          TRNOverviewBasicStat[float64]
	KillsPerMinute   TRNOverviewBasicStat[float64]
	MultiKills       TRNOverviewBasicStat[int]
	Level            TRNOverviewBasicStat[int]
	LevelProgression TRNOverviewBasicStat[float64]
	XPAll            TRNOverviewBasicStat[int]
}

type TRNOverviewBasicStat[T any] struct {
	Value        T
	DisplayValue string
	Percentile   *float64
}

func FetchBF2042OverviewData(platform, username string) (TRNOverviewResponse, error) {
	var trnResponse TRNOverviewResponse

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

// Unused, but a cool idea
func (trnResponse *TRNOverviewResponse) Fetch(platform, username string) error {
	url := utils.TRNBF2042OverviewURL(platform, username)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("TRN-Internal-Api-Key", os.Getenv("TRN_API_KEY"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("API returned a non-200 status code: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(trnResponse); err != nil {
		return err
	}

	return nil
}

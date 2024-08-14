package bf2042datafetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/leonlarsson/bfstats-go/internal/utils"
)

type trnWeaponsResponse struct {
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
			TimePlayed     trnWeaponBasicStat[int]
			Kills          trnWeaponBasicStat[int]
			ShotsAccuracy  trnWeaponBasicStat[float64]
			KillsPerMinute trnWeaponBasicStat[float64]
		}
	}
}

type trnWeaponBasicStat[T any] struct {
	Value        T
	DisplayValue string
	Percentile   float64
}

func FetchBF2042WeaponsData(platform, username string) (trnWeaponsResponse, error) {
	var trnResponse trnWeaponsResponse

	url := utils.TRNBF2042WeaponsURL(platform, username)

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

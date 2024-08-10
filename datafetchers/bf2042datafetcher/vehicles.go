package bf2042datafetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/leonlarsson/bfstats-bot-go/utils"
)

type trnVehiclesResponse struct {
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
			TimePlayed     trnVehicleBasicStat[int]
			Kills          trnVehicleBasicStat[int]
			KillsPerMinute trnVehicleBasicStat[float64]
		}
	}
}

type trnVehicleBasicStat[T any] struct {
	Value        T
	DisplayValue string
	Percentile   float64
}

func FetchBF2042VehiclesData(platform, username string) (trnVehiclesResponse, error) {
	var trnResponse trnVehiclesResponse

	url := utils.TRNBF2042VehiclesURL(platform, username)

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

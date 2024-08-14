package bf2042datafetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/leonlarsson/bfstats-go/internal/utils"
)

type trnClassesResponse struct {
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
			TimePlayed trnClassBasicStat[int]
			Kills      trnClassBasicStat[int]
		}
	}
}

type trnClassBasicStat[T any] struct {
	Value        T
	DisplayValue string
	Percentile   float64
}

func FetchBF2042ClassesData(platform, username string) (trnClassesResponse, error) {
	var trnResponse trnClassesResponse

	url := utils.TRNBF2042ClassesURL(platform, username)

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

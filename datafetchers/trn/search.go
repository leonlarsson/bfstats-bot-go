package trndatafetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/leonlarsson/bfstats-bot-go/utils"
)

type trnSearchResponse struct {
	Data []struct {
		PlatformUserIdentifier string
	}
}

func FetchTRNSearchData(game, platform, username string) (trnSearchResponse, error) {
	var trnResponse trnSearchResponse

	url := utils.TRNSearchURL(game, platform, username)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("TRN-Internal-Api-Key", os.Getenv("TRN_API_KEY"))
	if game == "bf1" {
		req.Header.Set("TRN-Api-Key", os.Getenv("TRN_API_KEY"))
	}

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

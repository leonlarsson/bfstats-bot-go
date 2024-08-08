package datafetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/leonlarsson/bfstats-bot-go/structs"
)

type BF2042TRNOverviewData struct {
}

func FetchBF2042OverviewData(platform, username string) (structs.TRNBF2042OverviewResponse, error) {
	var trnResponse structs.TRNBF2042OverviewResponse

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://public-api.tracker.gg/v2/bf2042/standard/profile/%s/%s", platform, username), nil)
	req.Header.Set("TRN-Internal-Api-Key", os.Getenv("TRN_API_KEY"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return structs.TRNBF2042OverviewResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return structs.TRNBF2042OverviewResponse{}, errors.New("TRN API returned a non-200 status code")
	}

	if err := json.NewDecoder(res.Body).Decode(&trnResponse); err != nil {
		return structs.TRNBF2042OverviewResponse{}, err
	}

	return trnResponse, nil

}

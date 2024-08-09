package datafetcher

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

func FetchData[T any](url string) (T, error) {
	var result T

	req, _ := http.NewRequest("GET", url, nil)

	// Add the TRN API key to the request if the URL is a TRN API URL
	if strings.HasPrefix(url, "https://public-api.tracker.gg") {
		req.Header.Set("TRN-Internal-Api-Key", os.Getenv("TRN_API_KEY"))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return result, errors.New("API returned a non-200 status code")
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

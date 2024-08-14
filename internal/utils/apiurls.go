package utils

import (
	"fmt"
	"net/url"
)

func TRNSearchURL(game string, platform, username string) string {
	baseUrl := fmt.Sprintf("https://public-api.tracker.gg/v2/%s/standard/search", game)

	queryParams := url.Values{
		"autocomplete": {"true"},
		"platform":     {platform},
		"query":        {username},
	}

	return baseUrl + "?" + queryParams.Encode()
}

func TRNBF2042OverviewURL(platform string, username string) string {
	return fmt.Sprintf("https://public-api.tracker.gg/v2/bf2042/standard/profile/%s/%s", platform, username)
}

func TRNBF2042WeaponsURL(platform string, username string) string {
	return fmt.Sprintf("https://public-api.tracker.gg/v2/bf2042/standard/profile/%s/%s/segments/weapon", platform, username)
}

func TRNBF2042VehiclesURL(platform string, username string) string {
	return fmt.Sprintf("https://public-api.tracker.gg/v2/bf2042/standard/profile/%s/%s/segments/vehicle", platform, username)
}

func TRNBF2042ClassesURL(platform string, username string) string {
	return fmt.Sprintf("https://public-api.tracker.gg/v2/bf2042/standard/profile/%s/%s/segments/soldier", platform, username)
}

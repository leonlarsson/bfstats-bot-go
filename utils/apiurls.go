package utils

import "fmt"

func TRNBF2042OverviewURL(platform string, username string) string {
	return fmt.Sprintf("https://public-api.tracker.gg/v2/bf2042/standard/profile/%s/%s", platform, username)
}

func TRNBF2042WeaponsURL(platform string, username string) string {
	return fmt.Sprintf("https://public-api.tracker.gg/v2/bf2042/standard/profile/%s/%s/segments/weapon", platform, username)
}

func TRNBF2042VehiclesURL(platform string, username string) string {
	return fmt.Sprintf("https://public-api.tracker.gg/v2/bf2042/standard/profile/%s/%s/segments/vehicle", platform, username)
}

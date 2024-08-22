package utils

import (
	"cmp"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

const UTMString = "?utm_source=discord&utm_medium=full-stats&utm_campaign=mozzy-bot"

type StatsLinkSettings struct {
	ApiURL        string
	Provider      string
	StatsCategory string
}

// BuildStatsLink builds a stats link based on the provided settings.
func BuildStatsLink(settings StatsLinkSettings) string {
	parsedApiUrl := parseApiUrl(settings.ApiURL)
	if parsedApiUrl == nil {
		return ""
	}

	if settings.Provider == "trn" ||
		strings.HasPrefix(settings.ApiURL, "https://public-api.tracker.gg/v2/") ||
		strings.HasPrefix(settings.ApiURL, "https://api.tracker.gg/api/v2/") {

		return fmt.Sprintf("https://battlefieldtracker.com/%s/profile/%s/%s/%s%s",
			parsedApiUrl.Game,
			toTRNPlatform(cmp.Or(parsedApiUrl.Platform, "origin")),
			url.PathEscape(cmp.Or(parsedApiUrl.Username, "mozzyfx")),
			cmp.Or(settings.StatsCategory, "overview"),
			UTMString)
	}

	if settings.Provider == "gt" || strings.HasPrefix(settings.ApiURL, "https://api.gametools.network/") {
		return fmt.Sprintf("https://gametools.network/stats/%s/name/%s?game=%s",
			toGTPlatform(cmp.Or(parsedApiUrl.Platform, "pc")),
			url.PathEscape(cmp.Or(parsedApiUrl.Username, "mozzyfx")),
			parsedApiUrl.Game)
	}

	// Return empty string if no matching apiURL
	return ""
}

type ApiInfo struct {
	Provider string
	Game     string
	Platform string
	Username string
}

// parseApiUrl parses the API URL and returns the game, platform, and username.
func parseApiUrl(apiURL string) *ApiInfo {
	// TRN API
	if strings.HasPrefix(apiURL, "https://public-api.tracker.gg/v2/") || strings.HasPrefix(apiURL, "https://api.tracker.gg/api/v2/") {
		gameRegex := regexp.MustCompile(`/v2/(\w+)/`)
		game := gameRegex.FindStringSubmatch(apiURL)[1]

		var platform, username string
		if strings.Contains(apiURL, "/standard/matches/") || strings.Contains(apiURL, "/gamereports/") {
			platformRegex := regexp.MustCompile(`/(?:gamereports|matches)/(\w+)/`)
			platform = platformRegex.FindStringSubmatch(apiURL)[1]

			usernameRegex := regexp.MustCompile(`/(?:gamereports/.+|matches)/\w+/(.+?)(?:/|$)`)
			username = usernameRegex.FindStringSubmatch(apiURL)[1]
		} else {
			platformRegex := regexp.MustCompile(`/profile/(\w+)/`)
			platform = platformRegex.FindStringSubmatch(apiURL)[1]

			usernameRegex := regexp.MustCompile(`/profile/\w+/(.+?)(?:/|$)`)
			username = usernameRegex.FindStringSubmatch(apiURL)[1]
		}

		return &ApiInfo{Provider: "trn", Game: game, Platform: platform, Username: username}
	}

	// GT API
	if strings.HasPrefix(apiURL, "https://api.gametools.network/") {
		parsedUrl, _ := url.Parse(apiURL)
		gameRegex := regexp.MustCompile(`gametools.network/(\w+)/`)
		game := gameRegex.FindStringSubmatch(apiURL)[1]

		platform := parsedUrl.Query().Get("platform")
		username := parsedUrl.Query().Get("name")

		return &ApiInfo{Provider: "gt", Game: game, Platform: platform, Username: username}
	}

	// Return nil if no matching apiURL
	return nil
}

type BattlelogLinkSettings struct {
	Game        string
	Username    string
	BattlelogId string
	Platform    string
}

// func buildBattlelogLink(settings BattlelogLinkSettings) string {
// 	// Check if any required argument is missing or if the game is not supported
// 	if settings.Game == "" || settings.Username == "" || settings.BattlelogId == "" || settings.Platform == "" {
// 		return ""
// 	}

// 	if !strings.Contains("bf3 bf4 bfh", settings.Game) {
// 		return ""
// 	}

// 	// Build the Battlelog link
// 	return fmt.Sprintf(
// 		"https://battlelog.battlefield.com/%s/soldier/%s/stats/%s/%s",
// 		settings.Game,
// 		url.PathEscape(settings.Username),
// 		settings.BattlelogId,
// 		settings.Platform,
// 	)
// }

// toTRNPlatform converts the platform to the TRN platform format.
func toTRNPlatform(platform string) string {
	platform = strings.ReplaceAll(platform, "pc", "origin")
	platform = strings.ReplaceAll(platform, "ps4", "psn")
	platform = strings.ReplaceAll(platform, "xboxone", "xbl")
	return platform
}

// toGTPlatform converts the platform to the GameTools platform format.
func toGTPlatform(platform string) string {
	platform = strings.ReplaceAll(platform, "origin", "pc")
	platform = strings.ReplaceAll(platform, "psn", "ps4")
	platform = strings.ReplaceAll(platform, "xbl", "xboxone")
	return platform
}

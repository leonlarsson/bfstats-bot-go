package datafetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/leonlarsson/bfstats-go/internal/bot/data"
	"github.com/leonlarsson/bfstats-go/internal/localization"
	"github.com/leonlarsson/bfstats-go/internal/utils"
)

// Fetch fetches data from the provided URL and decodes it into the provided type.
func Fetch[T any](url string, interaction *events.ApplicationCommandInteractionCreate, loc localization.LanguageLocalizer, username string) (T, error) {
	var result T

	statsLink := utils.BuildStatsLink(utils.StatsLinkSettings{
		ApiURL: url,
	})

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
		handleUnsuccessfulFetch(interaction, username, res.StatusCode, loc, statsLink)
		return result, errors.New("API returned a non-200 status code")
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

// FetchNoHandling fetches data from the provided URL and decodes it into the provided type. Without handling the response.
func FetchNoHandling[T any](url string) (T, error) {
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

// handleUnsuccessfulFetch handles an unsuccessful fetch by sending an error message to the user.
func handleUnsuccessfulFetch(interaction *events.ApplicationCommandInteractionCreate, username string, statusCode int, loc localization.LanguageLocalizer, statsLink string) {
	errorEmbed := discord.Embed{
		Color: 15548997,
		Footer: &discord.EmbedFooter{
			Text:    "Battlefield Stats - By Mozzy • /about • /help",
			IconURL: *data.ApplicationData.IconURL(),
		},
	}

	// 5xx errors
	if strings.HasPrefix(string(statusCode), "5") {
		errorEmbed.Title = fmt.Sprintf("%s (%d)", loc.Translate("messages/errors/fetch_server_error_title"), statusCode)
		description := loc.Translate("messages/errors/fetch_server_error_body", map[string]string{"username": username})
		errorEmbed.Description = description
		if strings.HasPrefix(statsLink, "https://battlefieldtracker.com/bf2042/profile") && statusCode == 400 {
			errorEmbed.Description += "\n\n⚠ **PSA: In order to view your Battlefield 2042 stats, you must activate 'Share\u00A0Usage\u00A0Data'.** [Instructions here](https://battlefieldtracker.com/bf2042/articles/changing-your-bf2042-privacy?utm_source=discord&utm_medium=full-stats&utm_campaign=mozzy-bot)."
		}
	} else if statusCode == 400 || statusCode == 404 {
		errorEmbed.Title = fmt.Sprintf("%s (%d)", loc.Translate("messages/errors/fetch_player_not_found_title"), statusCode)
		errorEmbed.Description = loc.Translate("messages/errors/fetch_player_not_found_body", map[string]string{"username": username})
	} else {
		errorEmbed.Title = fmt.Sprintf("%s (%d)", loc.Translate("messages/errors/fetch_something_went_wrong_title"), statusCode)
		errorEmbed.Description = loc.Translate("messages/errors/fetch_something_went_wrong_body", map[string]string{"username": username})
	}

	// Send the error message embed + buttons
	interaction.Client().Rest().UpdateInteractionResponse(interaction.ApplicationID(), interaction.Token(), discord.MessageUpdate{
		Embeds: &[]discord.Embed{errorEmbed},
		Components: &[]discord.ContainerComponent{
			discord.ActionRowComponent{
				utils.BuildFullStatsButton(statsLink, loc),
				utils.BuildInviteButton(loc),
			},
		},
	})
}

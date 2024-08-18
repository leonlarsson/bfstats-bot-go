package events

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/leonlarsson/bfstats-go/internal/datafetchers/trndatafetcher"
)

func HandleAutocomplete(event *events.AutocompleteInteractionCreate) {
	data := event.Data
	game := event.Data.CommandName
	username := data.String("username")
	platform := data.String("platform")

	var choices []discord.AutocompleteChoice

	// If fewer than 3 characters. Else fetch data from TRN.
	if len(username) < 3 {
		// If no input, respond with user's display name. Else, respond with the input.
		if len(username) == 0 {
			choices = append(choices, discord.AutocompleteChoiceString{
				Name:  event.Member().EffectiveName(),
				Value: event.Member().EffectiveName(),
			})

		} else {
			choices = append(choices, discord.AutocompleteChoiceString{
				Name:  username,
				Value: username,
			})
		}
	} else {
		// Fetch data
		data, err := trndatafetcher.FetchTRNSearchData(game, platform, username)
		if err != nil {
			return
		}

		// Build the response data for the first 25 users
		for i, user := range data.Data {
			if i >= 25 {
				break
			}
			choices = append(choices, discord.AutocompleteChoiceString{
				Name:  user.PlatformUserIdentifier,
				Value: user.PlatformUserIdentifier,
			})
		}
	}

	event.Respond(discord.InteractionResponseTypeAutocompleteResult, discord.AutocompleteResult{
		Choices: choices,
	})
}

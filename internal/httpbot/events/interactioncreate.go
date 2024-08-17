package events

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/leonlarsson/bfstats-go/internal/datafetchers/trndatafetcher"
)

func HandleAutocompleteInteraction(event *events.AutocompleteInteractionCreate) {
	fmt.Println(event)
}

func HandleInteractionCreateEvent(event *events.ApplicationCommandInteractionCreate) {
	// Autocomplete
	if event.Type() == discord.InteractionTypeAutocomplete {
		data := event.SlashCommandInteractionData()
		game := data.CommandName()
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

	// Chat input command
	// if interaction.Type == discordgo.InteractionApplicationCommand {
	// 	cmdData := interaction.ApplicationCommandData()
	// 	options := commands.ParseOptions(cmdData.Options)

	// 	locale := cmp.Or(string(commands.GetOptionStringValue(options, "language")), string(interaction.Locale), "en")
	// 	loc := *localization.CreateLocForLanguage(string(locale))
	// 	commandUsed := utils.GetCommandName(interaction)

	// 	// Retrieve and execute the handler
	// 	if handler, ok := handlers.HandlersMap[commandUsed]; ok {
	// 		// Type assert the handler func
	// 		if handler, ok := handler.(func(*discordgo.Session, *discordgo.InteractionCreate, localization.LanguageLocalizer) error); ok {
	// 			handler(s, interaction, loc)
	// 		} else {
	// 			log.Println("Bot: Handler function signature mismatch")
	// 		}
	// 	} else {
	// 		log.Println("Bot: Handler not found for command:", commandUsed)
	// 	}
	// }
}

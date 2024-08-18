package events

import (
	"cmp"
	"log"
	"regexp"

	"github.com/disgoorg/disgo/events"
	"github.com/leonlarsson/bfstats-go/internal/httpbot/handlers"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

func HandleInteractionCreate(interaction *events.ApplicationCommandInteractionCreate) {
	println("Bot: Interaction received")
	data := interaction.SlashCommandInteractionData()
	locale := cmp.Or(data.String("language"), interaction.Locale().Code(), "en")
	loc := *localization.CreateLocForLanguage(locale)
	commandUsed := getCommandName(*interaction)

	// Retrieve and execute the handler
	if handler, ok := handlers.HandlersMap[commandUsed]; ok {
		// Type assert the handler func
		if handler, ok := handler.(func(interaction *events.ApplicationCommandInteractionCreate, loc localization.LanguageLocalizer) error); ok {
			println("Bot: Handler found and running for command:", commandUsed)
			if err := handler(interaction, loc); err != nil {
				log.Println("Bot: Handler error:", err)
			}
		} else {
			log.Println("Bot: Handler function signature mismatch")
		}
	} else {
		log.Println("Bot: Handler not found for command:", commandUsed)
	}
}

func getCommandName(interaction events.ApplicationCommandInteractionCreate) (commandUsed string) {
	data := interaction.SlashCommandInteractionData()
	commandName := data.CommandName()
	commandIsGameCommand, _ := regexp.MatchString(`bf2042|bfv|bf1|bfh|bf4|bf3|bfbc2|bf2`, commandName)

	commandUsed = commandName

	if commandIsGameCommand {
		subcommand := *data.SubCommandName
		if subcommand == "stats" {
			segment := data.String("segment")
			commandUsed += segment
		} else {
			commandUsed += subcommand
		}
	}

	return commandUsed
}

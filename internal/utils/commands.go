package utils

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/leonlarsson/bfstats-go/internal/bot/commands"
)

func GetCommandName(interaction *discordgo.InteractionCreate) (commandUsed string) {
	commandData := interaction.ApplicationCommandData()
	commandOptions := commands.ParseOptions(commandData.Options)
	commandIsGameCommand, _ := regexp.MatchString(`bf2042|bfv|bf1|bfh|bf4|bf3|bfbc2|bf2`, commandData.Name)

	commandUsed = commandData.Name

	if commandIsGameCommand {
		subcommand := commands.GetOptionStringValue(commandOptions, "subcommand")
		if subcommand == "stats" {
			segment := commands.GetOptionStringValue(commandOptions, "segment")
			commandUsed = commandData.Name + segment
		} else {
			commandUsed = commandData.Name + subcommand
		}
	}

	return commandUsed
}

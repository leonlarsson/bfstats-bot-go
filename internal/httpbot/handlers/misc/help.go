package misc

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func HandleHelp(interaction *events.ApplicationCommandInteractionCreate) error {
	return interaction.CreateMessage(discord.MessageCreate{
		Content: "help",
	})
}

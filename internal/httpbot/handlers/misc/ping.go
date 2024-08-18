package misc

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

func HandlePing(event *handler.CommandEvent) error {
	return event.CreateMessage(discord.MessageCreate{
		Content: "pong",
		Flags:   discord.MessageFlagEphemeral,
	})
}

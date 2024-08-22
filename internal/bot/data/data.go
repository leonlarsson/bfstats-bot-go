package data

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
)

var ApplicationData *discord.Application

func FetchApplicationData(client bot.Client) {
	ApplicationData, _ = client.Rest().GetBotApplicationInfo()
}

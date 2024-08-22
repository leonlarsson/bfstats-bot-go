package utils

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

func BuildBaseRow(statsLink string, loc localization.LanguageLocalizer) discord.ActionRowComponent {
	return discord.ActionRowComponent{
		BuildFullStatsButton(statsLink, loc),
		BuildInviteButton(loc),
		BuildWebsiteButton(loc),
	}
}

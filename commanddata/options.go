package commanddata

import (
	"github.com/bwmarrin/discordgo"
	"github.com/leonlarsson/bfstats-bot-go/localization"
)

type optionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

// ParseOptions parses the options from an interaction and turns it into a map.
func ParseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) (om optionMap) {
	om = make(optionMap)
	for _, opt := range options {
		om[opt.Name] = opt
	}
	return
}

func UsernameOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:              "username",
		NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/username/name"),
		Description:       "The username to get stats for.",
		Type:              discordgo.ApplicationCommandOptionString,
		Required:          true,
	}
}

func PlatformOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:              "platform",
		NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/platform/name"),
		Description:       "The platform to get stats for.",
		Type:              discordgo.ApplicationCommandOptionString,
		Required:          true,
		Choices: []*discordgo.ApplicationCommandOptionChoice{
			{
				Name:  "PC/Origin",
				Value: "origin",
			},
			{
				Name:  "Xbox",
				Value: "xbox",
			},
			{
				Name:  "PlayStation",
				Value: "playstation",
			},
		},
	}
}

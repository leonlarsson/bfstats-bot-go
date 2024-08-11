package commanddata

import (
	"github.com/bwmarrin/discordgo"
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

var usernameOption = discordgo.ApplicationCommandOption{
	Name:        "username",
	Description: "The username to get stats for.",
	Type:        discordgo.ApplicationCommandOptionString,
	Required:    true,
}

var platformOption = discordgo.ApplicationCommandOption{
	Name:        "platform",
	Description: "The platform to get stats for.",
	Type:        discordgo.ApplicationCommandOptionString,
	Required:    true,
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

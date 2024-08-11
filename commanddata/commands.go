package commanddata

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "bf2042overview",
		Description: "Get Battlefield 2042 overview stats.",
		Options: []*discordgo.ApplicationCommandOption{
			&usernameOption,
			&platformOption,
		},
	},
	{
		Name:        "bf2042weapons",
		Description: "Get Battlefield 2042 weapon stats.",
		Options: []*discordgo.ApplicationCommandOption{
			&usernameOption,
			&platformOption,
		},
	},
	{
		Name:        "bf2042vehicles",
		Description: "Get Battlefield 2042 vehicle stats.",
		Options: []*discordgo.ApplicationCommandOption{
			&usernameOption,
			&platformOption,
		},
	},
}

package commands

import (
	"github.com/disgoorg/disgo/discord"
)

// GetCommands returns the commands for the bot.
func GetCommands() []discord.ApplicationCommandCreate {
	// Commonly used things
	// statsName := localization.GetEnglishString("slash_commands/stats/name")
	// statsNameLocalizations := localization.BuildDiscordLocalizations("slash_commands/stats/name")
	// statsDescription := localization.GetEnglishString("slash_commands/stats/bf2042_description")
	// statsDescriptionLocalizations := localization.BuildDiscordLocalizations("slash_commands/stats/bf2042_description")

	// segmentName := localization.GetEnglishString("slash_commands/stats/options/segment/name")
	// segmentNameLocalizations := localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/name")
	// segmentDescription := localization.GetEnglishString("slash_commands/stats/options/segment/description")
	// segmentDescriptionLocalizations := localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/description")

	return []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "ping",
			Description: "Ping command",
		},
		// discord.SlashCommandCreate{
		// 	Name:                     "bf2042",
		// 	Description:              localization.GetEnglishString("slash_commands/base/bf2042_description"),
		// 	DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/base/bf2042_description"),
		// 	Options: []discord.ApplicationCommandOption{
		// 		discord.ApplicationCommandOptionSubCommand{
		// 			Name:                     statsName,
		// 			NameLocalizations:        statsNameLocalizations,
		// 			Description:              statsDescription,
		// 			DescriptionLocalizations: statsDescriptionLocalizations,
		// 			Options: []discord.ApplicationCommandOption{
		// 				discord.ApplicationCommandOptionString{
		// 					Name:                     segmentName,
		// 					NameLocalizations:        segmentNameLocalizations,
		// 					Description:              segmentDescription,
		// 					DescriptionLocalizations: segmentDescriptionLocalizations,
		// 					Required:                 true,
		// 					Choices: []discord.ApplicationCommandOptionChoiceString{
		// 						OverviewSegment(),
		// 						WeaponsSegment(),
		// 						VehiclesSegment(),
		// 						ClassesSegment(),
		// 						GadgetsSegment(),
		// 						MapsSegment(),
		// 						ModesSegment(),
		// 						HazardZoneSegment(),
		// 					},
		// 				},
		// 				PlatformOption([]discord.ApplicationCommandOptionChoiceString{
		// 					{
		// 						Name:  "PC/Origin",
		// 						Value: "origin",
		// 					},
		// 					{
		// 						Name:  "Xbox",
		// 						Value: "xbl",
		// 					},
		// 					{
		// 						Name:  "PlayStation",
		// 						Value: "psn",
		// 					},
		// 				}),
		// 				UsernameOption(true),
		// 				LanguageOption(),
		// 				FormatOption(),
		// 				PoemGPTOption(),
		// 			},
		// 		},
		// 	},
		// },
	}
}

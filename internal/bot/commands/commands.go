package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/leonlarsson/bfstats-bot-go/internal/localization"
)

// GetCommands returns the commands for the bot.
func GetCommands() []*discordgo.ApplicationCommand {
	// Commonly used things
	statsName := localization.GetEnglishString("slash_commands/stats/name")
	statsNameLocalizations := localization.BuildDiscordLocalizations("slash_commands/stats/name")
	statsDescription := localization.GetEnglishString("slash_commands/stats/bf2042_description")
	statsDescriptionLocalizations := localization.BuildDiscordLocalizations("slash_commands/stats/bf2042_description")

	segmentName := localization.GetEnglishString("slash_commands/stats/options/segment/name")
	segmentNameLocalizations := localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/name")
	segmentDescription := localization.GetEnglishString("slash_commands/stats/options/segment/description")
	segmentDescriptionLocalizations := localization.BuildDiscordLocalizations("slash_commands/stats/options/segment/description")

	// Workaround. Ref: https://github.com/bwmarrin/discordgo/issues/1555
	descriptionLocations := localization.BuildDiscordLocalizations("slash_commands/base/bf2042_description")

	return []*discordgo.ApplicationCommand{
		{
			Name:                     "bf2042",
			Description:              localization.GetEnglishString("slash_commands/base/bf2042_description"),
			DescriptionLocalizations: &descriptionLocations,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:                     statsName,
					NameLocalizations:        statsNameLocalizations,
					Description:              statsDescription,
					DescriptionLocalizations: statsDescriptionLocalizations,

					Type: discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:                     segmentName,
							NameLocalizations:        segmentNameLocalizations,
							Description:              segmentDescription,
							DescriptionLocalizations: segmentDescriptionLocalizations,
							Type:                     discordgo.ApplicationCommandOptionString,
							Required:                 true,
							Choices: []*discordgo.ApplicationCommandOptionChoice{
								OverviewSegment(),
								WeaponsSegment(),
								VehiclesSegment(),
								ClassesSegment(),
								GadgetsSegment(),
								MapsSegment(),
								ModesSegment(),
								HazardZoneSegment(),
							},
						},
						PlatformOption([]*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  "PC/Origin",
								Value: "origin",
							},
							{
								Name:  "Xbox",
								Value: "xbl",
							},
							{
								Name:  "PlayStation",
								Value: "psn",
							},
						}),
						UsernameOption(true),
						LanguageOption(),
						FormatOption(),
						PoemGPTOption(),
					},
				},
			},
		},
	}
}

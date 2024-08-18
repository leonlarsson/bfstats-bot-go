package commands

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

// UsernameOption returns the username option for the stats command. Autocomplete is provided as an argument.
func UsernameOption(autocomplete bool) discord.ApplicationCommandOptionString {
	var maxLength = 100
	return discord.ApplicationCommandOptionString{
		Name:                     localization.GetEnglishString("slash_commands/stats/options/username/name"),
		NameLocalizations:        localization.BuildDiscordLocalizations("slash_commands/stats/options/username/name"),
		Description:              localization.GetEnglishString("slash_commands/stats/options/username/description"),
		DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/username/description"),
		MaxLength:                &maxLength,
		Autocomplete:             autocomplete,
		Required:                 true,
	}
}

// PlatformOption returns the platform option for the stats command. Actual platforms and values are provided as an argument.
func PlatformOption(platformChoices []discord.ApplicationCommandOptionChoiceString) discord.ApplicationCommandOptionString {
	return discord.ApplicationCommandOptionString{
		Name:                     localization.GetEnglishString("slash_commands/stats/options/platform/name"),
		NameLocalizations:        localization.BuildDiscordLocalizations("slash_commands/stats/options/platform/name"),
		Description:              localization.GetEnglishString("slash_commands/stats/options/platform/description"),
		DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/platform/description"),
		Required:                 true,
		Choices:                  platformChoices,
	}
}

// FormatOption returns the format option for the stats command.
func FormatOption() discord.ApplicationCommandOptionString {
	return discord.ApplicationCommandOptionString{
		Name:                     localization.GetEnglishString("slash_commands/stats/options/format/name"),
		NameLocalizations:        localization.BuildDiscordLocalizations("slash_commands/stats/options/format/name"),
		Description:              localization.GetEnglishString("slash_commands/stats/options/format/description"),
		DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/format/description"),
		Required:                 false,
		Choices: []discord.ApplicationCommandOptionChoiceString{
			{
				Name:              localization.GetEnglishString("slash_commands/stats/options/format/image_name"),
				NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/format/image_name"),
				Value:             "image",
			},
			{
				Name:              localization.GetEnglishString("slash_commands/stats/options/format/image_name") + "++",
				NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/format/image_name", "++"),
				Value:             "image2",
			},
			{
				Name:              localization.GetEnglishString("slash_commands/stats/options/format/text_name"),
				NameLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/format/text_name"),
				Value:             "text",
			},
		},
	}
}

// PoemGPTOption returns the poem GPT option for the stats command.
func PoemGPTOption() discord.ApplicationCommandOptionBool {
	return discord.ApplicationCommandOptionBool{
		Name:        "poem_gpt",
		Description: "Receive a beautiful poem about your stats.",
		Required:    false,
	}
}

// OverviewSegment returns the overview segment choice for the stats command.
func LanguageOption() discord.ApplicationCommandOptionString {
	languageChoices := []discord.ApplicationCommandOptionChoiceString{}
	locales := localization.GetLocales()

	for _, locale := range locales {
		loc := localization.CreateLocForLanguage(locale)
		languageChoices = append(languageChoices, discord.ApplicationCommandOptionChoiceString{
			Name:  fmt.Sprintf("%s (%s)", loc.Translate("meta/lang_localized"), loc.Translate("meta/lang")),
			Value: loc.Translate("meta/lang_code_discord"),
		})
	}

	return discord.ApplicationCommandOptionString{
		Name:                     localization.GetEnglishString("slash_commands/stats/options/language/name"),
		NameLocalizations:        localization.BuildDiscordLocalizations("slash_commands/stats/options/language/name"),
		Description:              localization.GetEnglishString("slash_commands/stats/options/language/description"),
		DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/language/description"),
		Required:                 false,
		Choices:                  languageChoices,
	}
}

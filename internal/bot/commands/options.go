package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

type optionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

// ParseOptions parses the options from an interaction and turns it into a map.
func ParseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) (om optionMap) {
	om = make(optionMap)
	for _, opt := range options {
		// If the option has suboptions, add subcommand to the map and then add the suboptions. Else just add the option.
		if len(opt.Options) > 0 {
			om["subcommand"] = opt
			for _, suboption := range opt.Options {
				om[suboption.Name] = suboption
			}
		} else {
			om[opt.Name] = opt
		}
	}
	return
}

// GetOptionStringValue returns the string value of an option if it exists, otherwise an empty string.
func GetOptionStringValue(options optionMap, key string) string {
	if option, exists := options[key]; exists {
		if option.Type != discordgo.ApplicationCommandOptionString {
			return option.Name
		} else {
			return option.StringValue()
		}
	}
	return ""
}

// GetOptionBoolValue returns the bool value of an option if it exists, otherwise false.
func GetOptionBoolValue(options optionMap, key string) bool {
	if option, exists := options[key]; exists {
		if option.Type != discordgo.ApplicationCommandOptionBoolean {
			return false
		} else {
			return option.BoolValue()
		}
	}
	return false
}

// UsernameOption returns the username option for the stats command. Autocomplete is provided as an argument.
func UsernameOption(autocomplete bool) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:                     localization.GetEnglishString("slash_commands/stats/options/username/name"),
		NameLocalizations:        localization.BuildDiscordLocalizations("slash_commands/stats/options/username/name"),
		Description:              localization.GetEnglishString("slash_commands/stats/options/username/description"),
		DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/username/description"),
		MaxLength:                100,
		Type:                     discordgo.ApplicationCommandOptionString,
		Autocomplete:             autocomplete,
		Required:                 true,
	}
}

// PlatformOption returns the platform option for the stats command. Actual platforms and values are provided as an argument.
func PlatformOption(platformChoices []*discordgo.ApplicationCommandOptionChoice) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:                     localization.GetEnglishString("slash_commands/stats/options/platform/name"),
		NameLocalizations:        localization.BuildDiscordLocalizations("slash_commands/stats/options/platform/name"),
		Description:              localization.GetEnglishString("slash_commands/stats/options/platform/description"),
		DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/platform/description"),
		Type:                     discordgo.ApplicationCommandOptionString,
		Required:                 true,
		Choices:                  platformChoices,
	}
}

// FormatOption returns the format option for the stats command.
func FormatOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:                     localization.GetEnglishString("slash_commands/stats/options/format/name"),
		NameLocalizations:        localization.BuildDiscordLocalizations("slash_commands/stats/options/format/name"),
		Description:              localization.GetEnglishString("slash_commands/stats/options/format/description"),
		DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/format/description"),
		Type:                     discordgo.ApplicationCommandOptionString,
		Required:                 false,
		Choices: []*discordgo.ApplicationCommandOptionChoice{
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
func PoemGPTOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:        "poem_gpt",
		Description: "Receive a beautiful poem about your stats.",
		Type:        discordgo.ApplicationCommandOptionBoolean,
		Required:    false,
	}
}

// OverviewSegment returns the overview segment choice for the stats command.
func LanguageOption() *discordgo.ApplicationCommandOption {
	languageChoices := []*discordgo.ApplicationCommandOptionChoice{}
	locales := localization.GetLocales()

	for _, locale := range locales {
		loc := localization.CreateLocForLanguage(locale)
		languageChoices = append(languageChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  fmt.Sprintf("%s (%s)", loc.Translate("meta/lang_localized"), loc.Translate("meta/lang")),
			Value: loc.Translate("meta/lang_code_discord"),
		})
	}

	return &discordgo.ApplicationCommandOption{
		Name:                     localization.GetEnglishString("slash_commands/stats/options/language/name"),
		NameLocalizations:        localization.BuildDiscordLocalizations("slash_commands/stats/options/language/name"),
		Description:              localization.GetEnglishString("slash_commands/stats/options/language/description"),
		DescriptionLocalizations: localization.BuildDiscordLocalizations("slash_commands/stats/options/language/description"),
		Type:                     discordgo.ApplicationCommandOptionString,
		Required:                 false,
		Choices:                  languageChoices,
	}
}

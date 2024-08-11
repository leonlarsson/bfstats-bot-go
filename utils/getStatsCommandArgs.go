package utils

import (
	"cmp"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/leonlarsson/bfstats-bot-go/commanddata"
	"github.com/leonlarsson/bfstats-bot-go/localization"
)

// GetStatsCommandArgs gets the arguments for the stats commands and defers the response if username validation passes.
func GetStatsCommandArgs(session *discordgo.Session, interaction *discordgo.InteractionCreate, loc *localization.LanguageLocalizer) (username, platform string, usernameFailedValidation bool) {
	options := commanddata.ParseOptions(interaction.ApplicationCommandData().Options)
	username = options["username"].StringValue()
	platform = options["platform"].StringValue()

	// If the username is "me", use the member's display name
	if strings.ToLower(username) == "me" {
		username = cmp.Or(interaction.Member.DisplayName())
	}

	// Build username regex
	// If BF2, allow periods
	usernameRegex := regexp.MustCompile("^[A-z0-9-–_ ]+$")
	if interaction.Interaction.ApplicationCommandData().Name == "bf2" {
		usernameRegex = regexp.MustCompile("^[A-z0-9-–_.,()/|*@=# ]+$")
	}

	// If the username contains newlines, backticks, or doesn't match the regex, fail the validation and respond with an error message
	if strings.Contains(username, "\n") || strings.Contains(username, "`") || !usernameRegex.MatchString(username) {
		usernameFailedValidation = true
		session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: loc.Translate("messages/invalid_username_unallowed_characters") + "\n\nIf this is your actual in-game name, [contact me](https://x.com/mozzyfx).",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}

	// If the validation passed, defer the response
	if !usernameFailedValidation {
		session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		})
	}

	return username, platform, usernameFailedValidation
}

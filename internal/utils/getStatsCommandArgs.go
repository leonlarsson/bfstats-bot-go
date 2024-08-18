package utils

import (
	"cmp"
	"regexp"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

// GetStatsCommandArgs gets the arguments for the stats commands and defers the response if username validation passes.
func GetStatsCommandArgs(interaction *events.ApplicationCommandInteractionCreate, loc localization.LanguageLocalizer) (username, platform string, usernameFailedValidation bool) {
	data := interaction.SlashCommandInteractionData()
	commandName := data.CommandName()
	username = data.String("username")
	platform = data.String("platform")

	// If the username is "me", use the member's display name
	if strings.ToLower(username) == "me" {
		username = cmp.Or(interaction.Member().EffectiveName())
	}

	// Build username regex
	// If BF2, allow periods
	usernameRegex := regexp.MustCompile("^[A-z0-9-–_ ]+$")
	if commandName == "bf2" {
		usernameRegex = regexp.MustCompile("^[A-z0-9-–_.,()/|*@=# ]+$")
	}

	// If the username contains newlines, backticks, or doesn't match the regex, fail the validation and respond with an error message
	if strings.Contains(username, "\n") || strings.Contains(username, "`") || !usernameRegex.MatchString(username) {
		usernameFailedValidation = true
		interaction.CreateMessage(discord.MessageCreate{
			Content: loc.Translate("messages/invalid_username_unallowed_characters") + "\n\nIf this is your actual in-game name, [contact me](https://x.com/mozzyfx).",
			Flags:   discord.MessageFlagEphemeral,
		})
	}

	// If the validation passed, defer the response
	if !usernameFailedValidation {
		interaction.DeferCreateMessage(false)
	}

	return username, platform, usernameFailedValidation
}

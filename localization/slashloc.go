package localization

import (
	"slices"

	"github.com/bwmarrin/discordgo"
)

// BuildDiscordLocalizations builds a map of Discord localizations for a given key.
func BuildDiscordLocalizations(key string) map[discordgo.Locale]string {
	locales := GetLocales()

	// Locales to skip. en is the default locale, ar is not supported by Discord.
	localesToSkip := []string{"en", "ar"}

	translations := make(map[discordgo.Locale]string, len(locales)-len(localesToSkip))
	for _, locale := range locales {
		if skip := slices.Contains(localesToSkip, locale); skip {
			continue
		}
		loc := CreateLocForLanguage(locale)
		translations[discordgo.Locale(loc.Translate("meta/lang_code_discord"))] = loc.Translate(key)
	}
	return translations
}

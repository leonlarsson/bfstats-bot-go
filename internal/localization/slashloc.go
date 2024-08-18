package localization

import (
	"slices"

	"github.com/disgoorg/disgo/discord"
)

// GetEnglishString returns the English translation for a given key.
func GetEnglishString(key string) string {
	return CreateLocForLanguage("en").Translate(key)
}

// BuildDiscordLocalizations builds a map of Discord localizations for a given key.
func BuildDiscordLocalizations(key string, suffix ...string) map[discord.Locale]string {
	locales := GetLocales()

	// Locales to skip. en is the default locale, ar is not supported by Discord.
	localesToSkip := []string{"en", "ar"}

	translations := make(map[discord.Locale]string, len(locales)-len(localesToSkip))
	for _, locale := range locales {
		if skip := slices.Contains(localesToSkip, locale); skip {
			continue
		}
		loc := CreateLocForLanguage(locale)

		suffixToAdd := ""

		// If suffix is provided, append it to the result of loc.Translate(key)
		if len(suffix) > 0 {
			suffixToAdd = suffix[0]
		}

		translations[discord.Locale(loc.Translate("meta/lang_code_discord"))] = loc.Translate(key) + suffixToAdd
	}
	return translations
}

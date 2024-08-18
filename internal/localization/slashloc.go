package localization

import (
	"slices"
)

// GetEnglishString returns the English translation for a given key.
func GetEnglishString(key string) string {
	return CreateLocForLanguage("en").Translate(key)
}

// BuildDiscordLocalizations builds a map of Discord localizations for a given key.
func BuildDiscordLocalizations[T ~string](key string, suffix ...string) map[T]string {
	locales := GetLocales()

	// Locales to skip. en is the default locale, ar is not supported by Discord.
	localesToSkip := []string{"en", "ar"}

	translations := make(map[T]string, len(locales)-len(localesToSkip))
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

		translations[T(loc.Translate("meta/lang_code_discord"))] = loc.Translate(key) + suffixToAdd
	}
	return translations
}

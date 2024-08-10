package localization

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	Bundle *i18n.Bundle
)

func init() {
	Bundle = GetBundle()
}

// GetBundle returns a new i18n.Bundle with the English language as the default
func GetBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	return bundle
}

// LoadLocales loads all locale files from the assets/locales directory
func LoadLocales() error {
	locFiles := []string{
		"assets/locales/ar.json",
		"assets/locales/de.json",
		"assets/locales/en.json",
		"assets/locales/es.json",
		"assets/locales/fi.json",
		"assets/locales/fr.json",
		"assets/locales/it.json",
		"assets/locales/nb.json",
		"assets/locales/pl.json",
		"assets/locales/pt.json",
		"assets/locales/ru.json",
		"assets/locales/sv.json",
		"assets/locales/tr.json",
	}

	for _, file := range locFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read locale file %s: %w", file, err)
		}
		modifiedContent := replaceDelimiters(string(content))
		if _, err := Bundle.ParseMessageFileBytes([]byte(modifiedContent), file); err != nil {
			return fmt.Errorf("failed to parse locale file %s: %w", file, err)
		}
	}
	return nil
}

// replaceDelimiters replaces the delimiters used in the locale files
func replaceDelimiters(s string) string {
	s = strings.ReplaceAll(s, "{{{", "{{.")
	s = strings.ReplaceAll(s, "}}}", "}}")
	return s
}

// Localize returns a localized string based on the key and data provided. This is just a shorthand for the Localizer.MustLocalize method.
func Localize(localizer *i18n.Localizer, key string, data map[string]string) string {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: data,
	})

	if err != nil {
		return "<string not found>"
	}

	return msg
}

// LanguageLocalizer is a struct that contains functions for translating and formatting strings based on a specific language
type LanguageLocalizer struct {
	Translate               func(string, ...map[string]string) string
	TranslateWithColon      func(string, ...map[string]string) string
	FormatInt               func(int) string
	FormatFloat             func(float64, int) string
	FormatPercent           func(float64, int) string
	SelectedLocale          string
	SelectedLocaleNumbers   string
	SelectedLocaleHumanizer string
	SelectedLocaleDiscord   string
}

// CreateLocForLanguage creates a new LanguageLocalizer for the specified language
func CreateLocForLanguage(lang string) *LanguageLocalizer {
	// Create a new localizer for the language
	localizer := i18n.NewLocalizer(Bundle, lang)

	// Translate function that takes a key and optional data map
	translate := func(key string, data ...map[string]string) string {
		// Default to an empty map if no data is provided
		dataMap := map[string]string{}
		if len(data) > 0 {
			dataMap = data[0]
		}
		return Localize(localizer, key, dataMap)
	}

	// Shortcut for translating a key with a colon at the end
	translateWithColon := func(key string, data ...map[string]string) string {
		return fmt.Sprintf("%s%s", translate(key, data...), translate("meta/colon"))
	}

	selectedLocaleNumbers := translate("meta/lang_code_numbers", nil)
	selectedLocaleHumanizer := translate("meta/lang_code_humanizer", nil)
	selectedLocaleDiscord := translate("meta/lang_code_discord", nil)

	printer := message.NewPrinter(language.Make(selectedLocaleNumbers))

	formatInt := func(i int) string {
		return printer.Sprintf("%d", i)
	}

	// TODO: Accept max fraction digits as a parameter
	formatFloat := func(f float64, maxFractionDigits int) string {
		return printer.Sprintf("%.*f", maxFractionDigits, f)
	}

	// TODO: Accept max fraction digits as a parameter
	formatPercent := func(f float64, maxFractionDigits int) string {
		return translate("stats/extra/x_percent", map[string]string{"number": printer.Sprintf("%.*f", maxFractionDigits, f)})
	}

	return &LanguageLocalizer{
		Translate:               translate,
		TranslateWithColon:      translateWithColon,
		FormatInt:               formatInt,
		FormatFloat:             formatFloat,
		FormatPercent:           formatPercent,
		SelectedLocale:          lang,
		SelectedLocaleNumbers:   selectedLocaleNumbers,
		SelectedLocaleHumanizer: selectedLocaleHumanizer,
		SelectedLocaleDiscord:   selectedLocaleDiscord,
	}
}

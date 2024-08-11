package localization

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
		"ar.json",
		"de.json",
		"en.json",
		"es.json",
		"fi.json",
		"fr.json",
		"it.json",
		"nb.json",
		"pl.json",
		"pt.json",
		"ru.json",
		"sv.json",
		"tr.json",
	}

	for _, file := range locFiles {
		path := filepath.Join("assets", "locales", file)
		content, err := os.ReadFile(path)
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

// LanguageLocalizer is a struct that contains functions for translating and formatting strings based on a specific language
type LanguageLocalizer struct {
	localizer               *i18n.Localizer
	printer                 *message.Printer
	SelectedLocale          string
	SelectedLocaleNumbers   string
	SelectedLocaleHumanizer string
	SelectedLocaleDiscord   string
}

// CreateLocForLanguage creates a new LanguageLocalizer for the specified language
func CreateLocForLanguage(lang string) *LanguageLocalizer {
	localizer := i18n.NewLocalizer(Bundle, lang)

	selectedLocaleNumbers := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "meta/lang_code_numbers",
	})
	selectedLocaleHumanizer := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "meta/lang_code_humanizer",
	})
	selectedLocaleDiscord := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "meta/lang_code_discord",
	})

	printer := message.NewPrinter(language.Make(selectedLocaleNumbers))

	return &LanguageLocalizer{
		localizer:               localizer,
		printer:                 printer,
		SelectedLocale:          lang,
		SelectedLocaleNumbers:   selectedLocaleNumbers,
		SelectedLocaleHumanizer: selectedLocaleHumanizer,
		SelectedLocaleDiscord:   selectedLocaleDiscord,
	}
}

func (loc LanguageLocalizer) Translate(key string, data ...map[string]string) string {
	// Default to an empty map if no data is provided
	dataMap := map[string]string{}
	if len(data) > 0 {
		dataMap = data[0]
	}

	msg, err := loc.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: dataMap,
	})

	if err != nil {
		return "<string not found>"
	}

	return msg
}

func (loc LanguageLocalizer) TranslateWithColon(key string, data ...map[string]string) string {
	return fmt.Sprintf("%s%s", loc.Translate(key, data...), loc.Translate("meta/colon"))
}

func (loc LanguageLocalizer) FormatInt(i int) string {
	return loc.printer.Sprintf("%d", i)
}

func (loc LanguageLocalizer) FormatFloat(f float64, maxFractionDigits ...int) string {
	maxDigits := 2
	if len(maxFractionDigits) > 0 {
		maxDigits = maxFractionDigits[0]
	}

	// If the float is a whole number, return it as an integer to avoid unnecessary decimal points
	if f == float64(int64(f)) {
		return strconv.FormatInt(int64(f), 10)
	}

	return fmt.Sprintf("%.*f", maxDigits, f)
}

func (loc LanguageLocalizer) FormatPercent(f float64, maxFractionDigits ...int) string {
	maxDigits := 1
	if len(maxFractionDigits) > 0 {
		maxDigits = maxFractionDigits[0]
	}

	// If the float is a whole number, return it as an integer to avoid unnecessary decimal points
	if f == float64(int64(f)) {
		return loc.Translate("stats/extra/x_percent", map[string]string{"number": loc.printer.Sprintf("%d", int(f))})
	}

	return loc.Translate("stats/extra/x_percent", map[string]string{"number": loc.printer.Sprintf("%.*f", maxDigits, f)})
}

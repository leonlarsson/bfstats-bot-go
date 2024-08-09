package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/leonlarsson/bfstats-bot-go/apihandlers"
	commandhandlers "github.com/leonlarsson/bfstats-bot-go/commandhandlers/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/localization"
)

func init() {
	// Load .env
	godotenv.Load()

	// Load locales
	localization.LoadLocales()
}

func main() {
	loc := *localization.CreateLocForLanguage("sv")
	msg := loc.Translate("stats/extra/x_score", map[string]string{"score": loc.FormatInt(123456789)})
	println(msg)

	err := commandhandlers.HandleBF2042OverviewCommand("origin", "MozzyFX")
	if err != nil {
		println(err.Error())
	}

	r := http.NewServeMux()
	r.HandleFunc("/bf2042/overview", apihandlers.BF2042OverviewHandler)
	r.HandleFunc("/bf2042/weapons", apihandlers.BF2042WeaponsHandler)
	http.ListenAndServe(":8080", r)
}

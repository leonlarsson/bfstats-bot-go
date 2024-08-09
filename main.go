package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/leonlarsson/bfstats-bot-go/apihandlers"
	commandhandlers "github.com/leonlarsson/bfstats-bot-go/commandhandlers/bf2042"
)

func init() {
	godotenv.Load()
}

func main() {
	err := commandhandlers.HandleBF2042OverviewCommand("origin", "MozzyFX")
	if err != nil {
		println(err.Error())
	}

	r := http.NewServeMux()
	r.HandleFunc("/bf2042/overview", apihandlers.BF2042OverviewHandler)
	r.HandleFunc("/bf2042/weapons", apihandlers.BF2042WeaponsHandler)
	http.ListenAndServe(":8080", r)
}

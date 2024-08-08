package main

import (
	"os"

	commandhandlers "github.com/leonlarsson/bfstats-bot-go/commandhandlers/bf2042"
)

func init() {
	os.Setenv("TRN_API_KEY", "XXX")
}

func main() {
	err := commandhandlers.HandleBF2042OverviewCommand("origin", "MozzyFX")
	if err != nil {
		println(err.Error())
	}
}

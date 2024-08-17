package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/leonlarsson/bfstats-go/internal/bot"
	"github.com/leonlarsson/bfstats-go/internal/localization"
)

var deployCommands bool

func init() {
	// Load locales
	localization.LoadLocales()

	// Load .env
	godotenv.Load()

	// Check if all required environment variables are set
	requiredEnvVars := []string{"TRN_API_KEY", "BOT_TOKEN", "BOT_ID", "GUILD_ID", "BOT_PUBLIC_KEY"}

	for _, varName := range requiredEnvVars {
		if os.Getenv(varName) == "" {
			log.Fatalf("Bot: Environment variable %s is not set", varName)
		}
	}

	// Parse command line flags
	flag.BoolVar(&deployCommands, "dc", false, "Deploy commands to Discord on start")
	flag.Parse()
}

func main() {
	log.Println("Bot: Starting")
	bot.Start(deployCommands)
}

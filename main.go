package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/leonlarsson/bfstats-bot-go/api"
	"github.com/leonlarsson/bfstats-bot-go/bot"
	"github.com/leonlarsson/bfstats-bot-go/localization"
)

func init() {
	// Load .env
	godotenv.Load()

	// Load locales
	localization.LoadLocales()

	// Check environment variables
	bot.CheckEnvVars()
}

func main() {
	var wg sync.WaitGroup

	// Start API service
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("API: Attempting to start API service")
		api.Start(":8080")
	}()

	// Start Discord bot
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Bot: Attempting to start Discord bot")
		bot.Start()
	}()

	// Wait for both goroutines to finish
	wg.Wait()
}

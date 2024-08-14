package main

import (
	"flag"
	"log"
	"sync"

	"github.com/leonlarsson/bfstats-bot-go/internal/api"
	"github.com/leonlarsson/bfstats-bot-go/internal/bot"
)

var (
	startApi       bool
	startBot       bool
	deployCommands bool
)

func init() {
	// Parse command line flags
	flag.BoolVar(&startApi, "api", false, "Start API service")
	flag.BoolVar(&startBot, "bot", false, "Start bot service")
	flag.BoolVar(&deployCommands, "dc", false, "Deploy commands to Discord on start")
	flag.Parse()
}

func main() {
	var wg sync.WaitGroup

	// If no services are to be started, log and return
	if !startApi && !startBot {
		log.Println("No services to start")
		return
	}

	// Start API service
	if startApi {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("API: Attempting to start API service")
			api.StartServer(":8080")
		}()
	}

	// Start Discord bot
	if startBot {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("Bot: Attempting to start Discord bot")
			bot.Start(deployCommands)
		}()
	}

	// Wait for both goroutines to finish
	wg.Wait()
}

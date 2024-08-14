package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/leonlarsson/bfstats-go/internal/bot/commands"
	"github.com/leonlarsson/bfstats-go/internal/bot/events"
	"github.com/servusdei2018/shards/v2"
)

func Start(deployCommands bool) {
	var err error

	mgr, err := shards.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Bot: Error creating shards manager: %v", err)
	}

	// Deploy commands (currently on every start, will change to an API call later probably)
	if deployCommands {
		log.Println("Bot: Attempting to deploy commands to Discord")
		err := mgr.ApplicationCommandBulkOverwrite(os.Getenv("GUILD_ID"), commands.GetCommands())
		if err != nil {
			log.Printf("Bot: Error deploying commands: %v", err)
			return
		}
		log.Println("Bot: Successfully deployed commands to Discord")
	}

	// Ready event
	mgr.AddHandler(events.HandleReadyEvent)

	// Interaction event
	mgr.AddHandler(events.HandleInteractionCreateEvent)

	// Open a websocket connection to Discord and begin listening.
	err = mgr.Start()
	if err != nil {
		log.Fatalf("Bot: Error opening connection: %v", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	err = mgr.Shutdown()
	if err != nil {
		log.Printf("Bot: Error closing connection: %v", err)
		return
	}
	log.Println("Bot: Stopped")
}

package bot

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/leonlarsson/bfstats-bot-go/commanddata"
	"github.com/leonlarsson/bfstats-bot-go/localization"
	"github.com/servusdei2018/shards/v2"
)

var deployCommands bool

func init() {
	// Load locales
	localization.LoadLocales()

	// Load .env
	godotenv.Load()

	// Check if all required environment variables are set
	CheckEnvVars()

	// Parse command line flags
	flag.BoolVar(&deployCommands, "deploy-commands", false, "Deploy commands to Discord on start")
}

// CheckEnvVars checks if all required environment variables are set.
func CheckEnvVars() {
	// Verify required environment variables
	requiredEnvVars := []string{"TRN_API_KEY", "BOT_TOKEN", "BOT_ID", "GUILD_ID"}

	for _, varName := range requiredEnvVars {
		if os.Getenv(varName) == "" {
			log.Fatalf("Bot: Environment variable %s is not set", varName)
		}
	}
}

// Start starts the Discord bot.
func Start() {

	var err error

	mgr, err := shards.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Bot: Error creating shards manager: %v", err)
	}

	// Deploy commands (currently on every start, will change to an API call later probably)
	if deployCommands {
		log.Println("Bot: Attempting to deploy commands to Discord")
		err := mgr.ApplicationCommandBulkOverwrite(os.Getenv("GUILD_ID"), commanddata.GetCommands())
		if err != nil {
			log.Printf("Bot: Error deploying commands: %v", err)
			return
		}
		log.Println("Bot: Successfully deployed commands to Discord")
	}

	// Ready event
	mgr.AddHandler(HandleReady)

	// Interaction event
	mgr.AddHandler(HandleInteractionCreate)

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

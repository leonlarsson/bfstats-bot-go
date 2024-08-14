package bot

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/leonlarsson/bfstats-bot-go/commanddata"
	"github.com/leonlarsson/bfstats-bot-go/localization"
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
	// Create a new Discord session using the provided bot token. The equivalent of discord.js's new Client()
	session, _ := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	var err error

	// Deploy commands (currently on every start, will change to an API call later probably)
	if deployCommands {
		log.Println("Bot: Attempting to deploy commands to Discord")
		commands, err := session.ApplicationCommandBulkOverwrite(os.Getenv("BOT_ID"), os.Getenv("GUILD_ID"), commanddata.GetCommands())
		if err != nil {
			log.Printf("Bot: Error deploying commands: %v", err)
			return
		}
		log.Printf("Bot: Successfully deployed %v commands to Discord", len(commands))
	}

	// Ready event
	session.AddHandler(HandleReady)

	// Interaction event
	session.AddHandler(HandleInteractionCreate)

	// Open a websocket connection to Discord and begin listening.
	err = session.Open()
	if err != nil {
		log.Printf("Bot: Error opening connection: %v", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	// Cleanly close down the Discord session.
	err = session.Close()
	if err != nil {
		log.Printf("Bot: Error closing connection: %v", err)
		return
	}
}

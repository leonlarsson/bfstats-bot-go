package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

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
	// _, err = session.ApplicationCommandBulkOverwrite(os.Getenv("BOT_ID"), os.Getenv("GUILD_ID"), commanddata.GetCommands())
	// if err != nil {
	// 	log.Printf("Bot: Error deploying commands: %v", err)
	// 	return
	// }

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

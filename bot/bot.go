package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/leonlarsson/bfstats-bot-go/commanddata"
	commandhandlers "github.com/leonlarsson/bfstats-bot-go/commandhandlers/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/localization"
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

func Start() {
	// Create a new Discord session using the provided bot token. The equivalent of discord.js's new Client()
	session, _ := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	// Deploy commands (currently on every start, will change this later)
	_, err := session.ApplicationCommandBulkOverwrite(os.Getenv("BOT_ID"), os.Getenv("GUILD_ID"), commanddata.Commands)
	if err != nil {
		log.Printf("Bot: Error deploying commands: %v", err)
		return
	}

	// Ready event
	session.AddHandler(func(s *discordgo.Session, readyEvent *discordgo.Ready) {
		log.Printf("Bot: Running as %s", readyEvent.User.Username)
	})

	// Interaction event
	session.AddHandler(func(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if interaction.Type == discordgo.InteractionApplicationCommand {
			loc := *localization.CreateLocForLanguage(string(interaction.Locale))

			switch interaction.ApplicationCommandData().Name {
			case "bf2042overview":
				err := commandhandlers.HandleBF2042OverviewCommand(s, interaction, loc)
				if err != nil {
					log.Printf("Bot: Error in bf2042overview command: %v", err)
				}
			case "bf2042weapons":
				err := commandhandlers.HandleBF2042WeaponsCommand(s, interaction, loc)
				if err != nil {
					log.Printf("Bot: Error in bf2042weapons command: %v", err)
				}
			case "bf2042vehicles":
				err := commandhandlers.HandleBF2042VehiclesCommand(s, interaction, loc)
				if err != nil {
					log.Printf("Bot: Error in bf2042vehicles command: %v", err)
				}
			}
		}
	})

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

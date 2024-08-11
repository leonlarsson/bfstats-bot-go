package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/leonlarsson/bfstats-bot-go/api"
	"github.com/leonlarsson/bfstats-bot-go/commanddata"
	commandhandlers "github.com/leonlarsson/bfstats-bot-go/commandhandlers/bf2042"
	"github.com/leonlarsson/bfstats-bot-go/localization"
)

func init() {
	// Load .env
	godotenv.Load()

	// Verify required environment variables
	requiredEnvVars := []string{"TRN_API_KEY", "BOT_TOKEN", "BOT_ID", "GUILD_ID"}

	for _, varName := range requiredEnvVars {
		if os.Getenv(varName) == "" {
			log.Fatalf("environment variable %s is not set", varName)
		}
	}

	// Load locales
	localization.LoadLocales()

	// Start API service
	api.Start()

	// TODO: Start Discord bot
	// discordbot.Start()
}

func main() {

	// Create a new Discord session using the provided bot token. The equivalent of discord.js's new Client()
	session, _ := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	// Deploy commands (currently on every start, will change this later)
	_, err := session.ApplicationCommandBulkOverwrite(os.Getenv("BOT_ID"), os.Getenv("GUILD_ID"), commanddata.Commands)
	if err != nil {
		println("Error creating commands: ", err)
		return
	}

	// Ready event
	session.AddHandler(func(s *discordgo.Session, readyEvent *discordgo.Ready) {
		log.Println("Bot is running as", readyEvent.User.Username)
	})

	// Interaction event
	session.AddHandler(func(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if interaction.Type == discordgo.InteractionApplicationCommand {
			loc := *localization.CreateLocForLanguage(string(interaction.Locale))

			switch interaction.ApplicationCommandData().Name {
			case "bf2042overview":
				err := commandhandlers.HandleBF2042OverviewCommand(s, interaction, loc)
				if err != nil {
					println(err.Error())
				}
			case "bf2042weapons":
				err := commandhandlers.HandleBF2042WeaponsCommand(s, interaction, loc)
				if err != nil {
					println(err.Error())
				}
			case "bf2042vehicles":
				err := commandhandlers.HandleBF2042VehiclesCommand(s, interaction, loc)
				if err != nil {
					println(err.Error())
				}
			}
		}
	})

	// Open a websocket connection to Discord and begin listening.
	err = session.Open()
	if err != nil {
		println("Error opening connection: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	// Cleanly close down the Discord session.
	err = session.Close()
	if err != nil {
		println("Error closing connection: ", err)
		return
	}
}

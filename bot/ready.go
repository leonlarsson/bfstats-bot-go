package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandleReady(s *discordgo.Session, readyEvent *discordgo.Ready) {
	log.Printf("Bot: Running as %s", readyEvent.User.String())
}

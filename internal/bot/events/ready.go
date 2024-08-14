package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandleReadyEvent(s *discordgo.Session, readyEvent *discordgo.Ready) {
	log.Printf("Bot: Shard #%d connected with %d guilds", s.ShardID, len(readyEvent.Guilds))
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Type: discordgo.ActivityTypeGame,
				Name: fmt.Sprintf("Battlefield • /about • /help • s%d", s.ShardID),
			},
		},
	})
}

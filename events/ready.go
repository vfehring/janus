package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateListeningStatus("keep Discord safe")
	log.Println("Logged in and ready to serve", len(session.State.Guilds), "servers")
}

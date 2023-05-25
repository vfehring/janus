package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateListeningStatus("keep Discord safe")
	fmt.Println(session.State.User.Username, "is Online")
}

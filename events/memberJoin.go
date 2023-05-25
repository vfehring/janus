package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	log.Println(m.User.Username, "joined the server")
}

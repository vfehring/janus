package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	fmt.Println(s.State.User.Username, "just joined the server.")
}

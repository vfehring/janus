package events

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "."

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bott itself
	// This isn't required in this specific example, but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(m.Content, " ")
	// if args[0] != prefix {
	// 	return
	// }

	// If the command is "ping" reply with "Pong!"
	if args[0] == prefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the command is "pong" reply with "Ping!"
	if args[0] == prefix+"pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

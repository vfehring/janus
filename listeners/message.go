package listeners

import (
	"vfehring/janus/util"

	"github.com/bwmarrin/discordgo"
)

type ListenerMsg struct {
}

func NewListenerMsg() *ListenerMsg {
	return &ListenerMsg{}
}

func (l *ListenerMsg) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Message.Author.ID == s.State.User.ID {
		return
	}
	util.Log.Infof("New message sent")
}

package listeners

import (
	"vfehring/janus/util"

	"github.com/bwmarrin/discordgo"
)

type ListenerReady struct {
}

func NewListenerReady() *ListenerReady {
	return &ListenerReady{}
}

func (l *ListenerReady) Handler(s *discordgo.Session, e *discordgo.Ready) {
	util.Log.Infof("Logged in as %s#%s (%s) - Running on %d servers",
		e.User.Username, e.User.Discriminator, e.User.ID, len(e.Guilds))
	util.Log.Infof("Invite Link: https://discordapp.com/api/oauth2/authorize?client_id=%s&scope=bot&permissions=%d",
		e.User.ID, util.InvitePermission)

	s.UpdateGameStatus(0, util.StdMotd)
	for _, g := range e.Guilds {
		if err := s.GuildMemberNickname(g.ID, "@me", util.AutoNick); err != nil {
			util.Log.Errorf("Failed updating nickname on guild %s (%s): %s", g.Name, g.ID, err)
		}
	}
}

package listeners

import (
	"fmt"
	"os"
	"strings"
	"vfehring/janus/commands"
	"vfehring/janus/util"

	"github.com/bwmarrin/discordgo"
)

type ListenerCmds struct {
	cmdHandler *commands.CmdHandler
}

func NewListenerCmd(cmdHandler *commands.CmdHandler) *ListenerCmds {
	return &ListenerCmds{
		cmdHandler: cmdHandler,
	}
}

func (l *ListenerCmds) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Message.Author.ID == s.State.User.ID {
		return
	}
	channel, err := s.Channel(e.ChannelID)
	if err != nil {
		util.Log.Errorf("Failed getting Discord channel from ID (%s): %s", e.ChannelID, err.Error())
		return
	}
	if channel.Type != discordgo.ChannelTypeGuildText {
		return
	}
	guildPrefix := os.Getenv("BASE_PREFIX")
	var pre string
	if guildPrefix != "" && strings.HasPrefix(e.Message.Content, guildPrefix) {
		pre = guildPrefix
	} else {
		return
	}

	contSplit := strings.Fields(e.Message.Content)
	invoke := contSplit[0][len(pre):]
	invoke = strings.ToLower(invoke)

	// UNFINISHED
	if cmdInstance, ok := l.cmdHandler.GetCommand(invoke); ok {
		guild, _ := s.Guild(e.GuildID)
		cmdArgs := &commands.CommandArgs{
			Args:       contSplit[1:],
			Channel:    channel,
			CmdHandler: l.cmdHandler,
			Guild:      guild,
			Message:    e.Message,
			Session:    s,
			User:       e.Author,
		}

		err = cmdInstance.Exec(cmdArgs)
		if err != nil {
			util.SendEmbedError(s, channel.ID, fmt.Sprintf("Failed executing command: ```\n%s\n```", err.Error()), "Command execution failed")
		}
	}
}

package util

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func SendEmbedError(s *discordgo.Session, chanID, content string, title ...string) (*discordgo.Message, error) {
	emb := &discordgo.MessageEmbed{
		Description: content,
		Color:       ColorEmbedError,
	}
	if len(title) > 0 {
		emb.Title = title[0]
	}
	return s.ChannelMessageSendEmbed(chanID, emb)
}

func SendEmbed(s *discordgo.Session, chanID, content string, title string, color int) (*discordgo.Message, error) {
	emb := &discordgo.MessageEmbed{
		Description: content,
		Color:       ColorEmbedError,
	}
	emb.Title = title
	if color == 0 {
		emb.Color = ColorEmbedDefault
	}
	return s.ChannelMessageSendEmbed(chanID, emb)
}

func DeleteMessageLater(s *discordgo.Session, msg *discordgo.Message, duration time.Duration) {
	if msg == nil {
		return
	}
	time.AfterFunc(duration, func() {
		s.ChannelMessageDelete(msg.ChannelID, msg.ID)
	})
}

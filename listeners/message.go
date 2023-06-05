package listeners

import (
	"vfehring/janus/models"
	"vfehring/janus/util"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm/clause"
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
	usr := models.User{ID: e.Message.Author.ID}
	userCreate := util.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&usr)

	if userCreate.Error != nil {
		util.Log.Fatal("Couldn't create the user")
	}
}

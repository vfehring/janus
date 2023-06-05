package listeners

import (
	"vfehring/janus/models"
	"vfehring/janus/util"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm/clause"
)

type ListenerAdd struct {
}

func NewListenerAdd() *ListenerAdd {
	return &ListenerAdd{}
}

func (l *ListenerAdd) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	if e.User.ID == s.State.User.ID {
		return
	}
	usr := models.User{ID: e.User.ID}
	result := util.DB.Model(&models.User{}).First(&usr)
	if result.RowsAffected > 0 {

	} else {
		// Add the user to the database if they do not already exist.
		userCreate := util.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&usr)
		if userCreate.Error != nil {
			util.Log.Fatal("Couldn't create the user")
		}
	}
}

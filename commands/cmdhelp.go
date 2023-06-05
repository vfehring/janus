package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"vfehring/janus/util"

	"github.com/bwmarrin/discordgo"
)

type CmdHelp struct {
}

func (c *CmdHelp) GetInvokes() []string {
	return []string{"help", "h", "?"}
}

func (c *CmdHelp) GetDescription() string {
	return "Display a list of commands or get help for a specific command"
}

func (c *CmdHelp) GetHelp() string {
	return "`help` - Display command list\n" +
		"`help <command>` - Display help of specific command"
}

func (c *CmdHelp) GetGroup() string {
	return GroupGeneral
}

func (c *CmdHelp) GetPermission() int {
	return 0
}

func (c *CmdHelp) Exec(args *CommandArgs) error {
	args.Session.ChannelMessageDelete(args.Channel.ID, args.Message.ID)

	emb := &discordgo.MessageEmbed{
		Color:  util.ColorEmbedDefault,
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	if len(args.Args) == 0 {
		cmds := make(map[string][]Command)
		for _, c := range args.CmdHandler.registeredCmdInstances {
			group := c.GetGroup()
			if _, ok := cmds[group]; !ok {
				cmds[group] = make([]Command, 0)
			}
			cmds[group] = append(cmds[group], c)
		}
		emb.Title = "Command List"
		for cat, catCmds := range cmds {
			commandHelpLines := ""
			for _, c := range catCmds {
				commandHelpLines += fmt.Sprintf("`%s` - *%s*\n", c.GetInvokes()[0], c.GetDescription())
			}
			emb.Fields = append(emb.Fields, &discordgo.MessageEmbedField{
				Name:  cat,
				Value: commandHelpLines,
			})
		}
	} else {
		cmd, ok := args.CmdHandler.GetCommand(args.Args[0])
		if !ok {
			msg, err := util.SendEmbedError(args.Session, args.Channel.ID,
				fmt.Sprintf("Sorry, there is no command with the invoke `%s`", args.Args[0]))
			util.DeleteMessageLater(args.Session, msg, 5*time.Second)
			return err
		}
		emb.Title = "Command Description"
		emb.Fields = []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Invokes",
				Value: strings.Join(cmd.GetInvokes(), "\n"),
			},
			&discordgo.MessageEmbedField{
				Name:   "Group",
				Value:  cmd.GetGroup(),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Permission Lvl",
				Value:  strconv.Itoa(cmd.GetPermission()),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:  "Description",
				Value: util.EnsureNotEmpty(cmd.GetDescription(), "`No description`"),
			},
			&discordgo.MessageEmbedField{
				Name:  "Usage",
				Value: util.EnsureNotEmpty(cmd.GetHelp(), "`No usage information`"),
			},
		}
	}

	userChan, err := args.Session.UserChannelCreate(args.User.ID)
	if err != nil {
		return err
	}
	_, err = args.Session.ChannelMessageSendEmbed(userChan.ID, emb)
	if err != nil {
		if strings.Contains(err.Error(), `{"code": 50007, "message": "Cannot send messages to this user"}`) {
			emb.Footer = &discordgo.MessageEmbedFooter{
				Text: "Actually, this message appears in your DM, but you have deactivated receiving DMs from" +
					"server members, so I can not send you this message via DM and you see this here right now.",
			}
			_, err = args.Session.ChannelMessageSendEmbed(args.Channel.ID, emb)
			return err
		}
	}

	return err
}

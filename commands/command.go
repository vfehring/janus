package commands

type Command interface {
	GetInvokes() []string
	GetDescription() string
	GetHelp() string
	GetGroup() string
	GetPermission() int
	Exec(args *CommandArgs) error
}

const (
	GroupGlobalAdmin = "GLOBAL ADMIN"
	GroupGuildAdmin  = "GUILD ADMIN"
	GroupModeration  = "MODERATION"
	GroupEtc         = "ETC"
	GroupGeneral     = "GENERAL"
)

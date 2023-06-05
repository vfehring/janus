package main

import (
	"os"
	"os/signal"
	"syscall"

	"vfehring/janus/commands"
	"vfehring/janus/listeners"
	"vfehring/janus/util"

	"github.com/bwmarrin/discordgo"
)

func init() {
	util.ConnectDB()
}

func main() {
	////////////////////
	// DATABASE LOGIN //
	////////////////////

	//////////////////////////
	// COMMAND REGISTRATION //
	//////////////////////////
	cmdHandler := commands.NewCmdHandler()
	cmdHandler.RegisterCommand(new(commands.CmdHelp))

	//////////////////////////
	// BOT SESSION CREATION //
	//////////////////////////
	discordToken := os.Getenv("DISCORD_TOKEN")
	session, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		util.Log.Fatal("Failed creating Discord bot sessiion:", err)
	}

	session.AddHandler(listeners.NewListenerReady().Handler)
	session.AddHandler(listeners.NewListenerCmd(cmdHandler).Handler)

	err = session.Open()
	if err != nil {
		util.Log.Fatal("Failed connecting Discord bot sessiion:", err)
	}

	util.Log.Info("Started event loop. Stop with CTRL-C...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	util.Log.Info("Shutting down...")
	session.Close()
}

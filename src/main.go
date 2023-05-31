package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/vfehring/janus/src/commands"
	"github.com/vfehring/janus/src/core"
	"github.com/vfehring/janus/src/listeners"
	"github.com/vfehring/janus/src/util"
)

var (
	configLocation = flag.String("c", "config.yml", "The location of the main config file")
)

func main() {
	flag.Parse()
	util.Log.Info("Starting up...")

	cfgParser := new(core.YAMLConfigParser)

	///////////////////////////
	// CONFIG INITIALIZATION //
	///////////////////////////

	cfgFile, err := os.Open(*configLocation)
	if os.IsNotExist(err) {
		cfgFile, err = os.Create(*configLocation)
		if err != nil {
			log.Fatal("Config file was not found and failed creating default config:", err)
		}
		err = cfgParser.Encode(cfgFile, core.NewDefaultConfig())
		if err != nil {
			log.Fatal("Config file was not found and failed writing to new config file:", err)
		}
		log.Fatal("Config file was not found. Created default config file. Please open it and enter your configuration.")
	} else if err != nil {
		log.Fatal("Failed opening config file:", err)
	}

	config, err := cfgParser.Decode(cfgFile)
	if err != nil {
		util.Log.Fatal("Failed decoding config file:", err)
	}

	if config.Discord.OwnerID == "" {
		util.Log.Warning("ATTENTION: Bot onwer ID is not set in config!",
			"You will not be identified as the owner of this bot so you will not have access to the owner-only commands!")
	}

	util.Log.Info("Config file loaded")

	////////////////////
	// DATABASE LOGIN //
	////////////////////

	database := new(core.MySql)
	if err := database.Connect(config.Database); err != nil {
		util.Log.Fatal("Failed connecting to database:", err)
	}
	util.Log.Info("Connected to database")

	//////////////////////////
	// COMMAND REGISTRATION //
	//////////////////////////

	cmdHandler := commands.NewCmdHandler(database, config)
	cmdHandler.RegisterCommand(new(commands.CmdTest))
	cmdHandler.RegisterCommand(new(commands.CmdHelp))
	cmdHandler.RegisterCommand(new(commands.CmdPrefix))

	//////////////////////////
	// BOT SESSION CREATION //
	//////////////////////////

	session, err := discordgo.New("Bot " + config.Discord.Token)
	if err != nil {
		util.Log.Fatal("Failed creating Discord bot session:", err)
	}

	session.AddHandler(listeners.NewListenerReady(config).Handler)
	session.AddHandler(listeners.NewListenerCmd(config, database, cmdHandler).Handler)

	err = session.Open()
	if err != nil {
		util.Log.Fatal("Failed connecting Discord bot session:", err)
	}

	util.Log.Info("Started event loop. Stop with CTRL-C...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	util.Log.Info("Shutting down...")
	session.Close()
}

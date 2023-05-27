package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandleReady(s *discordgo.Session, _ *discordgo.Ready) {
	log.Println("Logged in and ready to serve", len(s.State.Guilds), "servers")
}

func HandleGuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	log.Println("Joined guild", g.Name, " Connected to", len(s.State.Guilds), "Guilds")
}

func HandleGuildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	log.Println("Left guild", g.Name, " Connected to", len(s.State.Guilds))
}

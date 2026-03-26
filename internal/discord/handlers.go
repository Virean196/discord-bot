package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterHandlers(s *discordgo.Session) {
	s.AddHandler(onReady)
	s.AddHandler(onMessageCreate)
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as %s", r.User.String())
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("%s sent: %s", m.Author.Username, m.Content)
}

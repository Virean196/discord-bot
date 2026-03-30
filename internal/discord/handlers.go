package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/virean196/discord-bot/internal/features/ping"
)

func RegisterHandlers(s *discordgo.Session) {
	s.AddHandler(onReady)
	s.AddHandler(onMessageCreate)
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as %s", r.User.String())
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	response, ok := ping.Handle(m.Content)
	if ok {
		s.ChannelMessageSend(m.ChannelID, response)
		return
	}

	/* userID := m.Author.ID

	channel, err := s.UserChannelCreate(userID)
	if err != nil {
		return
	}

	_, err = s.ChannelMessageSend(channel.ID, "Pong in DM!")
	if err != nil {
		return
	} */

}

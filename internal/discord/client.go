package discord

import "github.com/bwmarrin/discordgo"

func NewClient(token string) (*discordgo.Session, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	s.Identify.Intents = discordgo.IntentsGuildMessages

	return s, nil
}

package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/virean196/discord-bot/internal/discord"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	token := os.Getenv("DISCORD_TOKEN")
	session, err := discord.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}

	discord.RegisterHandlers(session)

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Println("Bot is running...")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh
}

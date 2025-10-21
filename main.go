package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pupunha-code/Manicoba/articles"
	"github.com/pupunha-code/Manicoba/bot"
)

func main() {

	// Pega as variaveis de ambiente
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	bot.ArticleSender(bot.SessionCreator(token), channelID, articles.FetchMorningNews())
}

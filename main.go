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

	session, err := bot.SessionCreator(token)
	if err != nil {
		log.Fatal("Não foi possivel criar a seção: ", err)
	}

	news, err := articles.FetchMorningNews()
	if err != nil {
		log.Print("Não foi possivel retornar os artigos: ", err)
	} else {
		bot.ArticleSender(session, channelID, news)
	}

}

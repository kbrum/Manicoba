package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pupunha-code/Manicoba/articles"
	"github.com/pupunha-code/Manicoba/bot"
	"github.com/robfig/cron/v3"
)

func main() {

	// Pega as variaveis de ambiente
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregara a .env file")
	}

	token := os.Getenv("DISCORD_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	session, err := bot.SessionCreator(token)
	if err != nil {
		log.Fatal("Não foi possivel criar a seção: ", err)
	}

	log.Println("Bot iniciado")

	c := cron.New() // cria o agendador

	c.AddFunc("0 9 * * *", func() { // manda artigos de frontend
		log.Println("Buscando artigos frontend")
		article, err := articles.FetchArticles("javascript,typescript,react,nextjs,vue,angular,svelte,tailwindcss,vite&top=1")
		if err != nil {
			log.Println("Erro ao buscar artigo: ", err)
		} else {
			bot.ArticleSender(session, channelID, article)
		}
	})

	c.AddFunc("30 12 * * *", func() { // manda artigos de backend
		log.Println("Buscando artigos frontend")
		article, err := articles.FetchArticles("node,python,go,java,rust,springboot,django,fastapi,laravel,graphql,postgres,redis,backend,architecture")
		if err != nil {
			log.Println("Erro ao buscar artigo: ", err)
		} else {
			bot.ArticleSender(session, channelID, article)
		}
	})

	c.AddFunc(" 18 * * *", func() { // manda artigos de devops e cloud
		log.Println("Buscando artigos frontend")
		article, err := articles.FetchArticles("devops,aws,azure,gcp,docker,kubernetes,terraform,ansible,githubactions,cicd,prometheus,grafana,observability,security,linux")
		if err != nil {
			log.Println("Erro ao buscar artigo: ", err)
		} else {
			bot.ArticleSender(session, channelID, article)
		}
	})
}

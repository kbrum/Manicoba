package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/pupunha-code/Manicoba/articles"
)

// ClientStarter Função que faz start no client do discord
func ClientStarter(token string) {
	session, err := discordgo.New("Maniçoba " + token)

	// Tratamento de erros de sessão
	if err != nil {
		log.Fatal("Session Error:", err)
	}

}

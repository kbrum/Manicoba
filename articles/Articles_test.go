package articles

import (
	"log"
	"testing"
)

func TestArticles(t *testing.T) {
	_, err := FetchArticles("")
	if err != nil {
		log.Print("Não foi possivel fazer a solicitação:", err)
	}
}

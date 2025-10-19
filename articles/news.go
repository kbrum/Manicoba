package articles

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Article mapeamento do dados do json para a struct
type Article struct {
	Title       string `json:"title"`
	Description string `Json:"description"`
	URL         string `json:"url"`
	ImageURL    string `json:"social_image"`
}

// FetchArticles Função base que retorna o artigo que sera enviado
func FetchArticles(url string) Article {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}

	var articles []Article
	err = json.Unmarshal(bodyBytes, &articles)
	if err != nil {
		log.Fatal("Error:", err)
	}

	if len(articles) == 0 {
		log.Fatal("No articles found")
	}

	return articles[0]
}

// FetchMorningNews Noticias que serão mandadas as 9 horas (assuntos frontend)
func FetchMorningNews() Article {
	return FetchArticles("https://dev.to/api/articles?tags=javascript,typescript,react,nextjs,vue,angular,svelte,tailwindcss,vite&top=1")
}

// FetchAfternoonNews Noticias que serão mandadas as 14 horas (assuntos backend)
func FetchAfternoonNews() Article {
	return FetchArticles("https://dev.to/api/articles?tags=node,python,go,java,rust,springboot,django,fastapi,laravel,graphql,postgres,redis,backend,architecture")
}

// FetchNightNews Noticias que serão mandadas as 19 horas (assuntos devops/cloud)
func FetchNightNews() Article {
	return FetchArticles("https://dev.to/api/articles?tags=devops,aws,azure,gcp,docker,kubernetes,terraform,ansible,githubactions,cicd,prometheus,grafana,observability,security,linux")
}

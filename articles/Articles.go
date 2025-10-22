package articles

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Article Mapeamento do dados do json para a struct
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ImageURL    string `json:"social_image"`
}

// FetchArticles Função base que retorna o artigo que sera enviado
func FetchArticles(url string) (Article, error) {

	//Faz a solicitação http
	resp, err := http.Get(url)
	if err != nil {
		return Article{}, err
	}

	// Fecha a conexão com o endpoint no final da função
	resp.Body.Close()

	// trata codigos http diferentes de 200
	if resp.StatusCode != 200 {
		return Article{}, errors.New("http error: " + resp.Status)
	}

	// Acessa os dados da response e guarda na na slice do tipo []byte
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Article{}, err
	}

	// Faz o unmarshal do bodyBytes, salva na struct, e trata possiveis erros
	var articles []Article
	err = json.Unmarshal(bodyBytes, &articles)
	if err != nil {
		return Article{}, err
	}

	// tratamento para retorno de 0 artigos
	if len(articles) == 0 {
		return Article{}, errors.New("articles not found")
	}

	return articles[0], nil
}

// FetchMorningNews Noticias que serão mandadas as 9 horas (assuntos frontend)
func FetchMorningNews() (Article, error) {
	// faz a solicitação para o endpoint e faz o tratamento de erros
	article, err := FetchArticles("https://dev.to/api/articles?tags=javascript,typescript,react,nextjs,vue,angular,svelte,tailwindcss,vite&top=1")
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

// FetchAfternoonNews Noticias que serão mandadas as 14 horas (assuntos backend)
func FetchAfternoonNews() (Article, error) {
	article, err := FetchArticles("https://dev.to/api/articles?tags=node,python,go,java,rust,springboot,django,fastapi,laravel,graphql,postgres,redis,backend,architecture")
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

// FetchNightNews Noticias que serão mandadas as 19 horas (assuntos devops/cloud)
func FetchNightNews() (Article, error) {
	article, err := FetchArticles("https://dev.to/api/articles?tags=devops,aws,azure,gcp,docker,kubernetes,terraform,ansible,githubactions,cicd,prometheus,grafana,observability,security,linux")
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

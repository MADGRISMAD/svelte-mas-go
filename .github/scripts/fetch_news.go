package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	newsAPIURL = "https://newsapi.org/v2/top-headlines?country=mx&pageSize=1"
	contentDir = "content"
)

type NewsResponse struct {
	Articles []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Content     string `json:"content"`
		URL         string `json:"url"`
		Source      struct {
			Name string `json:"name"`
		} `json:"source"`
		PublishedAt string `json:"publishedAt"`
	} `json:"articles"`
}

func main() {
	apiKey := os.Getenv("NEWSAPI_KEY")
	if apiKey == "" {
		panic("NEWSAPI_KEY no está definido")
	}

	req, _ := http.NewRequest("GET", newsAPIURL, nil)
	req.Header.Set("Authorization", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var news NewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&news); err != nil {
		panic(err)
	}

	if len(news.Articles) == 0 {
		fmt.Println("No se encontraron noticias.")
		return
	}

	article := news.Articles[0]
	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s/news-%s.md", contentDir, date)

	mdContent := fmt.Sprintf(`# %s

**Fuente:** [%s](%s)  
**Fecha:** %s

%s

%s
`, article.Title, article.Source.Name, article.URL, article.PublishedAt[:10], article.Description, article.Content)

	if err := ioutil.WriteFile(filename, []byte(mdContent), 0644); err != nil {
		panic(err)
	}

	fmt.Println("Noticia guardada en:", filename)
}

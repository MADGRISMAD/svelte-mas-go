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
	newsAPIURL = "https://newsapi.org/v2/top-headlines?country=mx&category=general&language=es&pageSize=5"
	contentDir = "content"
)

type NewsResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
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

type ErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	apiKey := os.Getenv("NEWSAPI_KEY")
	if apiKey == "" {
		panic("NEWSAPI_KEY no está definido")
	}

	// Construir URL con API key
	url := fmt.Sprintf("%s&apiKey=%s", newsAPIURL, apiKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(fmt.Sprintf("Error creando request: %v", err))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("Error haciendo request: %v", err))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Error leyendo respuesta: %v", err))
	}

	// Intentar decodificar como error primero
	var errorResp ErrorResponse
	if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Status == "error" {
		panic(fmt.Sprintf("Error de API: %s - %s", errorResp.Code, errorResp.Message))
	}

	// Decodificar respuesta normal
	var news NewsResponse
	if err := json.Unmarshal(body, &news); err != nil {
		panic(fmt.Sprintf("Error decodificando respuesta: %v", err))
	}

	if news.Status != "ok" {
		panic(fmt.Sprintf("Estado de respuesta no OK: %s", news.Status))
	}

	if len(news.Articles) == 0 {
		fmt.Println("🟡 La API no devolvió artículos hoy. Nada que guardar.")
		return
	}

	// Asegurar que el directorio exista
	_ = os.MkdirAll(contentDir, os.ModePerm)

	date := time.Now().Format("2006-01-02")
	for i, article := range news.Articles {
		filename := fmt.Sprintf("%s/news-%s-%d.md", contentDir, date, i+1)

		mdContent := fmt.Sprintf(`# %s

**Fuente:** [%s](%s)  
**Fecha:** %s

%s

%s
`, article.Title, article.Source.Name, article.URL, article.PublishedAt[:10], article.Description, article.Content)

		if err := ioutil.WriteFile(filename, []byte(mdContent), 0644); err != nil {
			fmt.Printf("❌ Error guardando archivo %s: %v\n", filename, err)
			continue
		}

		fmt.Println("✅ Noticia guardada en:", filename)
	}
}

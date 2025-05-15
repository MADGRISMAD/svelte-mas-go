package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Estructura para la respuesta del contenido
type PageContent struct {
	Content string `json:"content"`
}

// Estructura para la creación de páginas
type CreatePageRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Middleware para manejar CORS
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	// Crear el directorio content si no existe
	contentDir := "../content"
	if err := createContentDir(contentDir); err != nil {
		log.Fatal(err)
	}

	// Endpoint para listar todas las páginas y crear nuevas
	http.HandleFunc("/api/pages", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			files, err := ioutil.ReadDir(contentDir)
			if err != nil {
				http.Error(w, "Error al leer el directorio", http.StatusInternalServerError)
				return
			}

			var pages []string
			for _, file := range files {
				if filepath.Ext(file.Name()) == ".md" {
					pages = append(pages, strings.TrimSuffix(file.Name(), ".md"))
				}
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pages)

		case http.MethodPost:
			var req CreatePageRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
				return
			}

			// Validar el nombre del archivo
			if req.Name == "" {
				http.Error(w, "El nombre de la página es requerido", http.StatusBadRequest)
				return
			}

			// Crear el archivo
			filePath := filepath.Join(contentDir, req.Name+".md")
			if err := ioutil.WriteFile(filePath, []byte(req.Content), 0644); err != nil {
				http.Error(w, "Error al crear el archivo", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "Página creada exitosamente"})

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}))

	// Endpoint para obtener el contenido de una página específica
	http.HandleFunc("/api/pages/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		slug := strings.TrimPrefix(r.URL.Path, "/api/pages/")
		if slug == "" {
			http.Error(w, "Slug no proporcionado", http.StatusBadRequest)
			return
		}

		content, err := ioutil.ReadFile(filepath.Join(contentDir, slug+".md"))
		if err != nil {
			http.Error(w, "Página no encontrada", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PageContent{Content: string(content)})
	}))

	// Servir archivos estáticos para el frontend
	fs := http.FileServer(http.Dir("../frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createContentDir(dir string) error {
	// Crear el directorio si no existe
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Crear el archivo welcome.md si no existe
	welcomeFile := filepath.Join(dir, "welcome.md")
	if _, err := os.Stat(welcomeFile); os.IsNotExist(err) {
		return ioutil.WriteFile(welcomeFile, []byte("# Bienvenido\n\nEsta es la página de bienvenida."), 0644)
	}
	return nil
}

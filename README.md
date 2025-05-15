# 📰 Noticias Diarias

Una aplicación web moderna para visualizar noticias diarias, con actualización automática y una interfaz elegante.

## 🚀 Características

- Visualización de noticias en formato Markdown
- Actualización automática diaria de noticias
- Interfaz moderna y responsive
- Backend en Go para servir las noticias
- Scripts de automatización para obtener noticias
- Integración con GitHub Actions para actualizaciones automáticas

## 🏗️ Arquitectura

El proyecto está dividido en tres partes principales:

### Frontend (Vue 3 + Vite)
- Interfaz de usuario moderna
- Visualización de noticias en Markdown
- Diseño responsive con Tailwind CSS
- Componentes Vue 3 con Composition API

### Backend (Go)
- API REST para servir noticias
- Manejo de archivos Markdown
- Servidor HTTP eficiente
- Estructura de carpetas organizada

### Scripts y Automatización
- Scripts Python para obtener noticias
- GitHub Actions para actualización automática
- Integración con APIs de noticias
- Generación automática de archivos Markdown

## 🛠️ Tecnologías

### Frontend
- Vue 3 con Composition API
- Vite como bundler
- Tailwind CSS para estilos
- Marked para renderizado de Markdown

### Backend
- Go (Golang)
- Gin Web Framework
- Estructura de carpetas modular
- Manejo eficiente de archivos

### Automatización
- Python 3.x
- Requests para APIs
- GitHub Actions
- Markdown generación

## 📋 Prerrequisitos

- Node.js (versión 16 o superior)
- Go 1.16 o superior
- Python 3.x
- npm o yarn
- Git

## 🔧 Instalación

1. Clona el repositorio:
```bash
git clone [URL_DEL_REPOSITORIO]
cd [NOMBRE_DEL_REPOSITORIO]
```

2. Configura el Frontend:
```bash
cd frontend
npm install
# o
yarn install
```

3. Configura el Backend:
```bash
cd backend
go mod download
```

4. Configura las variables de entorno:
Crea un archivo `.env` en la raíz del proyecto con:
```env
# Frontend
VITE_API_URL=http://localhost:8080

# Backend
PORT=8080
NEWSAPI_KEY=tu_api_key_aqui

# GitHub Actions
GH_PAT=tu_token_de_github_aqui
```

## 🚀 Desarrollo

### Frontend
```bash
cd frontend
npm run dev
# o
yarn dev
```
La aplicación estará disponible en `http://localhost:5173`

### Backend
```bash
cd backend
go run main.go
```
El servidor estará disponible en `http://localhost:8080`

## 📝 Estructura del Proyecto

```
.
├── frontend/                 # Aplicación Vue
│   ├── src/
│   │   ├── components/      # Componentes Vue
│   │   ├── assets/         # Recursos estáticos
│   │   ├── App.vue         # Componente principal
│   │   └── main.js         # Punto de entrada
│   └── package.json
│
├── backend/                 # Servidor Go
│   ├── main.go             # Punto de entrada
│   ├── handlers/           # Manejadores HTTP
│   ├── models/             # Modelos de datos
│   └── go.mod
│
├── content/                # Archivos de noticias
│   └── news-*.md          # Noticias en Markdown
│
├── .github/
│   └── workflows/         # GitHub Actions
│       └── daily-news.yml # Workflow de actualización
│
└── scripts/               # Scripts de automatización
    └── fetch_news.py      # Script para obtener noticias
```

## 🤖 Automatización

### GitHub Actions
El proyecto incluye un workflow de GitHub Actions que:
- Se ejecuta diariamente a las 7:00 UTC
- Obtiene noticias actualizadas
- Genera archivos Markdown
- Hace commit y push de los cambios

### Scripts
El script `fetch_news.py`:
- Se conecta a APIs de noticias
- Procesa y formatea el contenido
- Genera archivos Markdown
- Maneja errores y reintentos

## 🎨 Personalización

### Frontend
- Estilos en `frontend/src/App.vue`
- Configuración de Tailwind en `tailwind.config.js`
- Componentes personalizables en `src/components/`

### Backend
- Configuración en `backend/config/`
- Manejadores personalizables en `backend/handlers/`
- Modelos de datos en `backend/models/`

### Scripts
- Configuración de APIs en `scripts/config.py`
- Formato de noticias en `scripts/templates/`
- Lógica de procesamiento en `scripts/fetch_news.py`

## 🤝 Contribución

1. Haz fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE.md](LICENSE.md) para más detalles.

## 📞 Soporte

Si tienes alguna pregunta o sugerencia, por favor abre un issue en el repositorio.

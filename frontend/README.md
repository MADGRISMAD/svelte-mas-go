# 📰 Noticias Diarias

Una aplicación web moderna para visualizar noticias diarias, construida con Vue 3 y Vite.

## 🚀 Características

- Visualización de noticias en formato Markdown
- Diseño moderno y responsive
- Interfaz intuitiva y fácil de usar
- Actualización automática de noticias
- Soporte para múltiples fuentes de noticias

## 🛠️ Tecnologías

- Vue 3 con Composition API
- Vite como bundler
- Tailwind CSS para estilos
- Marked para renderizado de Markdown

## 📋 Prerrequisitos

- Node.js (versión 16 o superior)
- npm o yarn

## 🔧 Instalación

1. Clona el repositorio:
```bash
git clone [URL_DEL_REPOSITORIO]
cd frontend
```

2. Instala las dependencias:
```bash
npm install
# o
yarn install
```

3. Configura las variables de entorno:
Crea un archivo `.env` en la raíz del proyecto con:
```env
VITE_API_URL=http://localhost:8080
```

## 🚀 Desarrollo

Para iniciar el servidor de desarrollo:

```bash
npm run dev
# o
yarn dev
```

La aplicación estará disponible en `http://localhost:5173`

## 🏗️ Construcción

Para construir la aplicación para producción:

```bash
npm run build
# o
yarn build
```

Los archivos generados estarán en el directorio `dist/`

## 📝 Estructura del Proyecto

```
frontend/
├── src/
│   ├── components/     # Componentes Vue
│   ├── assets/        # Recursos estáticos
│   ├── App.vue        # Componente principal
│   └── main.js        # Punto de entrada
├── public/            # Archivos públicos
└── index.html         # Template HTML
```

## 🎨 Personalización

### Estilos
Los estilos están construidos con Tailwind CSS. Puedes personalizar el tema en:
- `tailwind.config.js` para configuración global
- `src/App.vue` para estilos específicos de la aplicación

### Markdown
El renderizado de Markdown se configura en `App.vue` usando la librería Marked.

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

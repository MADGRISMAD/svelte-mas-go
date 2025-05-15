<script setup>
import { ref, onMounted } from 'vue'
import { marked } from 'marked'

// Configurar marked para que use clases de Tailwind
marked.setOptions({
  gfm: true,
  breaks: true,
  headerIds: false,
  mangle: false
})

const pages = ref([])
const currentPage = ref(null)
const renderedContent = ref('')
const showNewPageModal = ref(false)
const newPageName = ref('')
const newPageContent = ref('')
const isLoading = ref(false)
const error = ref(null)

const formatPageTitle = (pageName) => {
  // Convertir el nombre del archivo en un título legible
  return pageName
    .replace('news-', '')
    .replace('.md', '')
    .split('-')
    .map(word => word.charAt(0).toUpperCase() + word.slice(1))
    .join(' ')
}

const loadPages = async () => {
  try {
    isLoading.value = true
    error.value = null
    const response = await fetch('http://localhost:8080/api/pages')
    pages.value = await response.json()
  } catch (error) {
    error.value = 'Error al cargar las páginas'
    console.error('Error al cargar las páginas:', error)
  } finally {
    isLoading.value = false
  }
}

const loadPage = async (pageName) => {
  try {
    isLoading.value = true
    error.value = null
    const response = await fetch(`http://localhost:8080/api/pages/${pageName}`)
    const data = await response.json()
    currentPage.value = pageName
    renderedContent.value = marked(data.content)
  } catch (error) {
    error.value = 'Error al cargar la página'
    console.error('Error al cargar la página:', error)
  } finally {
    isLoading.value = false
  }
}

const createPage = async () => {
  if (!newPageName.value.trim()) {
    error.value = 'El nombre de la página es requerido'
    return
  }

  try {
    isLoading.value = true
    error.value = null
    const response = await fetch('http://localhost:8080/api/pages', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: newPageName.value,
        content: newPageContent.value,
      }),
    })

    if (response.ok) {
      showNewPageModal.value = false
      newPageName.value = ''
      newPageContent.value = ''
      await loadPages()
      await loadPage(newPageName.value)
    } else {
      error.value = 'Error al crear la página'
    }
  } catch (error) {
    error.value = 'Error al crear la página'
    console.error('Error al crear la página:', error)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  loadPages()
})
</script>

<template>
  <div class="min-h-screen bg-gray-50 flex flex-col">
    <!-- Header -->
    <header class="bg-white shadow z-10 sticky top-0">
      <div class="max-w-7xl mx-auto px-4 py-3 flex justify-between items-center">
        <h1 class="text-xl font-bold text-gray-900">📰 Noticias Diarias</h1>
        <div class="flex items-center space-x-4">
          <span class="text-sm text-gray-500">{{ new Date().toLocaleDateString('es-ES', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' }) }}</span>
        </div>
      </div>
    </header>

    <!-- Body Layout -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar -->
      <aside class="w-64 bg-white border-r overflow-y-auto">
        <div class="p-4">
          <h2 class="text-sm font-semibold text-gray-500 uppercase tracking-wider mb-4">Noticias</h2>
          <ul class="space-y-1">
            <li
              v-for="page in pages"
              :key="page"
              @click="loadPage(page)"
              :class="[
                'block px-3 py-2 rounded-md text-sm font-medium cursor-pointer hover:bg-gray-50',
                currentPage === page ? 'bg-blue-50 text-blue-700 border-l-4 border-blue-500' : 'text-gray-700'
              ]"
            >
              {{ formatPageTitle(page) }}
            </li>
          </ul>
        </div>
      </aside>

      <!-- Main Content -->
      <main class="flex-1 p-6 overflow-y-auto">
        <div v-if="isLoading" class="flex justify-center items-center h-full">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        </div>
        <div v-else>
          <div v-if="currentPage" class="max-w-4xl mx-auto">
            <article class="bg-white rounded-lg shadow-sm p-6">
              <div v-html="renderedContent" class="markdown-content"></div>
            </article>
          </div>
          <div v-else class="text-center text-gray-500">
            <p>Selecciona una noticia para leer</p>
          </div>
        </div>
      </main>
    </div>

    <!-- New Page Modal -->
    <div v-if="showNewPageModal" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-lg w-full max-w-lg p-6 space-y-6">
        <h3 class="text-lg font-semibold text-gray-800">Crear Nueva Página</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">Nombre de la página</label>
            <input v-model="newPageName" class="w-full border border-gray-300 rounded-md p-2 text-sm" placeholder="Ej: nueva-guía" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Contenido (Markdown)</label>
            <textarea v-model="newPageContent" rows="8" class="w-full border border-gray-300 rounded-md p-2 font-mono text-sm" placeholder="# Mi Página"></textarea>
          </div>
        </div>
        <div class="flex justify-end space-x-2">
          <button @click="showNewPageModal = false" class="px-4 py-2 text-sm bg-gray-200 hover:bg-gray-300 rounded-md">Cancelar</button>
          <button @click="createPage" class="px-4 py-2 text-sm bg-blue-600 text-white hover:bg-blue-700 rounded-md">Crear</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
/* Estilos base */
.markdown-content {
  color: #1f2937;
  font-size: 1.125rem;
  line-height: 1.75;
  font-family: system-ui, -apple-system, sans-serif;
}

/* Encabezados */
.markdown-content h1 {
  font-size: 2.5rem;
  line-height: 1.2;
  font-weight: 800;
  margin-bottom: 1.5rem;
  color: #111827;
  letter-spacing: -0.025em;
}

/* Párrafos y texto */
.markdown-content p {
  margin-bottom: 1.5rem;
  line-height: 1.8;
  color: #374151;
}

.markdown-content strong {
  font-weight: 700;
  color: #111827;
}

/* Enlaces */
.markdown-content a {
  color: #2563eb;
  text-decoration: none;
  border-bottom: 1px solid #93c5fd;
  transition: all 0.2s;
}

.markdown-content a:hover {
  color: #1d4ed8;
  border-bottom-color: #1d4ed8;
}

/* Citas */
.markdown-content blockquote {
  border-left: 4px solid #3b82f6;
  padding: 1rem 1.5rem;
  margin: 1.5rem 0;
  background-color: #f8fafc;
  border-radius: 0.375rem;
  font-style: italic;
  color: #4b5563;
}

/* Código */
.markdown-content code {
  background-color: #f1f5f9;
  border-radius: 0.25rem;
  padding: 0.2rem 0.4rem;
  font-size: 0.875rem;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
  color: #0f172a;
}

/* Imágenes */
.markdown-content img {
  max-width: 100%;
  height: auto;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
  margin: 1.5rem 0;
}

/* Líneas horizontales */
.markdown-content hr {
  margin: 2rem 0;
  border: 0;
  border-top: 2px solid #e5e7eb;
}

/* Estilos para la fecha y fuente */
.markdown-content p:first-of-type {
  color: #6b7280;
  font-size: 0.875rem;
  margin-bottom: 2rem;
}

.markdown-content p:first-of-type strong {
  color: #4b5563;
  font-weight: 600;
}

/* Estilos para el contenido principal */
.markdown-content p:not(:first-of-type) {
  font-size: 1.125rem;
  line-height: 1.8;
  color: #1f2937;
}
</style>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>

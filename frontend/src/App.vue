<script setup>
import { ref, onMounted, computed } from 'vue'
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

// Estado para controlar qué fechas están expandidas
const expandedDates = ref(new Set())

const toggleDate = (date) => {
  if (expandedDates.value.has(date)) {
    expandedDates.value.delete(date)
  } else {
    expandedDates.value.add(date)
  }
}

// Agrupar noticias por fecha
const groupedNews = computed(() => {
  const groups = {}
  console.log('Procesando páginas:', pages.value) // Debug log
  pages.value.forEach(page => {
    console.log('Procesando página:', page) // Debug log
    // Extraer la fecha del nombre del archivo (formato: news-2025-05-15-1.md)
    const parts = page.split('-')
    if (parts[0] === 'news') {
      const date = `${parts[1]}-${parts[2]}-${parts[3]}`
      console.log('Fecha extraída:', date) // Debug log
      if (!groups[date]) {
        groups[date] = []
      }
      groups[date].push(page)
    }
  })
  console.log('Noticias agrupadas:', groups) // Debug log
  return groups
})

// Ordenar fechas de más reciente a más antigua
const sortedDates = computed(() => {
  return Object.keys(groupedNews.value).sort((a, b) => new Date(b) - new Date(a))
})

const formatPageTitle = (pageName) => {
  console.log('Formateando título para:', pageName) // Debug log
  // Extraer el título del nombre del archivo (formato: news-2025-05-15-1.md)
  const parts = pageName.split('-')
  if (parts[0] === 'news') {
    return `Noticia ${parts[4].replace('.md', '')}`
  }
  return pageName
}

const formatDate = (dateStr) => {
  console.log('Formateando fecha:', dateStr) // Debug log
  // La fecha ya viene en formato YYYY-MM-DD
  const [year, month, day] = dateStr.split('-')
  return new Date(year, month - 1, day).toLocaleDateString('es-ES', {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const loadPages = async () => {
  try {
    isLoading.value = true
    error.value = null
    const response = await fetch('http://localhost:8080/api/pages')
    pages.value = await response.json()
    console.log('Páginas cargadas:', pages.value) // Debug log
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
  <div class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 flex flex-col">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b border-gray-200 z-10 sticky top-0">
      <div class="max-w-7xl mx-auto px-4 py-4 flex justify-between items-center">
        <div class="flex items-center space-x-3">
          <span class="text-2xl">📰</span>
          <h1 class="text-xl font-bold bg-gradient-to-r from-blue-600 to-blue-800 bg-clip-text text-transparent">
            Noticias Diarias
          </h1>
        </div>
        <div class="flex items-center space-x-4">
          <span class="text-sm text-gray-600 bg-gray-50 px-3 py-1.5 rounded-full">
            {{ new Date().toLocaleDateString('es-ES', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' }) }}
          </span>
        </div>
      </div>
    </header>

    <!-- Body Layout -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar -->
      <aside class="w-96 bg-white border-r border-gray-200 overflow-y-auto">
        <div class="p-6">
          <h2 class="text-sm font-semibold text-gray-500 uppercase tracking-wider mb-6">Noticias por Fecha</h2>
          <div class="space-y-2">
            <div v-for="date in sortedDates" :key="date" class="border border-gray-200 rounded-lg overflow-hidden">
              <!-- Encabezado del dropdown -->
              <button
                @click="toggleDate(date)"
                class="w-full px-4 py-3 flex items-center justify-between bg-gray-50 hover:bg-gray-100 transition-colors duration-200"
              >
                <span class="text-sm font-medium text-gray-900">{{ formatDate(date) }}</span>
                <span 
                  class="transform transition-transform duration-200"
                  :class="expandedDates.has(date) ? 'rotate-180' : ''"
                >
                  <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </span>
              </button>
              
              <!-- Contenido del dropdown -->
              <div
                v-show="expandedDates.has(date)"
                class="bg-white"
              >
                <ul class="divide-y divide-gray-100">
                  <li
                    v-for="page in groupedNews[date]"
                    :key="page"
                    @click="loadPage(page)"
                    :class="[
                      'group relative px-4 py-3 text-sm font-medium cursor-pointer transition-all duration-200',
                      currentPage === page 
                        ? 'bg-blue-50 text-blue-700' 
                        : 'text-gray-700 hover:bg-gray-50'
                    ]"
                  >
                    <div class="flex items-center justify-between">
                      <span>{{ formatPageTitle(page) }}</span>
                      <span 
                        class="opacity-0 group-hover:opacity-100 transition-opacity duration-200"
                        :class="currentPage === page ? 'text-blue-500' : 'text-gray-400'"
                      >
                        →
                      </span>
                    </div>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </aside>

      <!-- Main Content -->
      <main class="flex-1 p-8 overflow-y-auto">
        <div v-if="isLoading" class="flex justify-center items-center h-full">
          <div class="animate-spin rounded-full h-12 w-12 border-4 border-blue-200 border-t-blue-600"></div>
        </div>
        <div v-else>
          <div v-if="currentPage" class="max-w-4xl mx-auto">
            <article class="bg-white rounded-xl shadow-sm p-8">
              <div v-html="renderedContent" class="markdown-content"></div>
            </article>
          </div>
          <div v-else class="text-center py-12">
            <div class="text-6xl mb-4">📰</div>
            <p class="text-gray-500 text-lg">Selecciona una noticia para comenzar a leer</p>
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
  background: linear-gradient(to right, #1e40af, #3b82f6);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

/* Párrafos y texto */
.markdown-content p {
  margin-bottom: 1.5rem;
  line-height: 1.8;
  color: #374151;
}

.markdown-content p:first-of-type {
  color: #6b7280;
  font-size: 0.875rem;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid #e5e7eb;
}

.markdown-content p:first-of-type strong {
  color: #4b5563;
  font-weight: 600;
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
  padding-bottom: 1px;
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
  border-radius: 0.5rem;
  font-style: italic;
  color: #4b5563;
  box-shadow: 0 1px 3px 0 rgb(0 0 0 / 0.1);
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
  border-radius: 0.75rem;
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
  margin: 1.5rem 0;
  transition: transform 0.2s;
}

.markdown-content img:hover {
  transform: scale(1.01);
}

/* Líneas horizontales */
.markdown-content hr {
  margin: 2rem 0;
  border: 0;
  border-top: 2px solid #e5e7eb;
}

/* Estilos para el contenido principal */
.markdown-content p:not(:first-of-type) {
  font-size: 1.125rem;
  line-height: 1.8;
  color: #1f2937;
}

/* Animaciones */
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.markdown-content > * {
  animation: fadeIn 0.3s ease-out forwards;
}

/* Estilos para listas */
.markdown-content ul,
.markdown-content ol {
  margin: 1.5rem 0;
  padding-left: 1.5rem;
}

.markdown-content li {
  margin-bottom: 0.75rem;
  color: #374151;
}

.markdown-content li::marker {
  color: #3b82f6;
}

/* Estilos para tablas */
.markdown-content table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0;
  margin: 1.5rem 0;
  border-radius: 0.5rem;
  overflow: hidden;
  box-shadow: 0 1px 3px 0 rgb(0 0 0 / 0.1);
}

.markdown-content th {
  background-color: #f8fafc;
  padding: 0.75rem 1rem;
  text-align: left;
  font-weight: 600;
  color: #1f2937;
  border-bottom: 2px solid #e5e7eb;
}

.markdown-content td {
  padding: 0.75rem 1rem;
  border-bottom: 1px solid #e5e7eb;
  color: #374151;
}

.markdown-content tr:last-child td {
  border-bottom: none;
}

.markdown-content tr:hover td {
  background-color: #f8fafc;
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

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
  <div class="min-h-screen bg-gray-100 flex flex-col">
    <!-- Header -->
    <header class="bg-white shadow z-10 sticky top-0">
      <div class="max-w-7xl mx-auto px-4 py-3 flex justify-between items-center">
        <h1 class="text-xl font-bold text-gray-900">📘 Wiki Console</h1>
        <button
          @click="showNewPageModal = true"
          class="px-4 py-2 text-sm text-white bg-blue-600 hover:bg-blue-700 rounded-md"
        >
          + Nueva Página
        </button>
      </div>
    </header>

    <!-- Body Layout -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar -->
      <aside class="w-64 bg-white border-r overflow-y-auto">
        <div class="p-4">
          <h2 class="text-sm font-semibold text-gray-500 uppercase tracking-wider mb-4">Páginas</h2>
          <ul class="space-y-1">
            <li
              v-for="page in pages"
              :key="page"
              @click="loadPage(page)"
              :class="[
                'block px-3 py-2 rounded-md text-sm font-medium cursor-pointer',
                currentPage === page ? 'bg-blue-100 text-blue-700' : 'text-gray-700 hover:bg-gray-200'
              ]"
            >
              {{ page }}
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
          <div v-if="currentPage" class="prose max-w-none">
            <div v-html="renderedContent" class="markdown-content"></div>
          </div>
          <div v-else class="text-center text-gray-500">
            <p>Selecciona o crea una página para comenzar</p>
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
  font-size: 1rem;
  line-height: 1.75;
}

/* Encabezados */
.markdown-content h1 {
  font-size: 2.25rem;
  line-height: 2.5rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  color: #111827;
}

.markdown-content h2 {
  font-size: 1.875rem;
  line-height: 2.25rem;
  font-weight: 600;
  margin-bottom: 1rem;
  margin-top: 2rem;
  color: #111827;
}

.markdown-content h3 {
  font-size: 1.5rem;
  line-height: 2rem;
  font-weight: 600;
  margin-bottom: 0.75rem;
  margin-top: 1.5rem;
  color: #111827;
}

.markdown-content h4 {
  font-size: 1.25rem;
  line-height: 1.75rem;
  font-weight: 600;
  margin-bottom: 0.75rem;
  margin-top: 1.25rem;
  color: #111827;
}

.markdown-content h5 {
  font-size: 1.125rem;
  line-height: 1.75rem;
  font-weight: 600;
  margin-bottom: 0.75rem;
  margin-top: 1.25rem;
  color: #111827;
}

.markdown-content h6 {
  font-size: 1rem;
  line-height: 1.5rem;
  font-weight: 600;
  margin-bottom: 0.75rem;
  margin-top: 1.25rem;
  color: #111827;
}

/* Párrafos y texto */
.markdown-content p {
  margin-bottom: 1rem;
  line-height: 1.625;
}

.markdown-content strong {
  font-weight: 700;
}

.markdown-content em {
  font-style: italic;
}

.markdown-content del,
.markdown-content s {
  text-decoration: line-through;
}

/* Listas */
.markdown-content ul,
.markdown-content ol {
  margin: 1rem 0;
  padding-left: 1.5rem;
}

.markdown-content ul {
  list-style-type: disc;
}

.markdown-content ol {
  list-style-type: decimal;
}

.markdown-content li {
  margin-bottom: 0.5rem;
  color: #374151;
}

.markdown-content li > ul,
.markdown-content li > ol {
  margin-top: 0.5rem;
  margin-bottom: 0;
}

/* Enlaces */
.markdown-content a {
  color: #2563eb;
  text-decoration: underline;
}

.markdown-content a:hover {
  color: #1d4ed8;
}

/* Citas */
.markdown-content blockquote {
  border-left: 4px solid #d1d5db;
  padding: 0.5rem 1rem;
  margin: 1rem 0;
  color: #4b5563;
  font-style: italic;
}

.markdown-content blockquote > * {
  margin: 0.5rem 0;
}

.markdown-content blockquote blockquote {
  border-left-width: 2px;
  margin-left: 1rem;
}

/* Código */
.markdown-content code {
  background-color: #f3f4f6;
  border-radius: 0.25rem;
  padding: 0.125rem 0.25rem;
  font-size: 0.875rem;
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
}

.markdown-content pre {
  background-color: #f3f4f6;
  border-radius: 0.25rem;
  padding: 1rem;
  margin: 1rem 0;
  overflow-x: auto;
}

.markdown-content pre code {
  background-color: transparent;
  padding: 0;
  font-size: 0.875rem;
  line-height: 1.5;
}

/* Tablas */
.markdown-content table {
  width: 100%;
  border-collapse: collapse;
  margin: 1rem 0;
}

.markdown-content th {
  background-color: #f3f4f6;
  border: 1px solid #d1d5db;
  padding: 0.5rem 1rem;
  text-align: left;
  font-weight: 600;
}

.markdown-content td {
  border: 1px solid #d1d5db;
  padding: 0.5rem 1rem;
}

.markdown-content th[style*="text-align:right"],
.markdown-content td[style*="text-align:right"] {
  text-align: right;
}

/* Imágenes */
.markdown-content img {
  max-width: 100%;
  height: auto;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1);
  margin: 1rem 0;
}

/* Líneas horizontales */
.markdown-content hr {
  margin: 2rem 0;
  border: 0;
  border-top: 1px solid #d1d5db;
}

/* Caracteres especiales */
.markdown-content sup {
  vertical-align: super;
  font-size: smaller;
}

.markdown-content sub {
  vertical-align: sub;
  font-size: smaller;
}

/* Definiciones */
.markdown-content dl {
  margin: 1rem 0;
}

.markdown-content dt {
  font-weight: 600;
  margin-top: 1rem;
}

.markdown-content dd {
  margin-left: 1.5rem;
  margin-bottom: 0.5rem;
}

/* Notas al pie */
.markdown-content .footnotes {
  margin-top: 2rem;
  padding-top: 1rem;
  border-top: 1px solid #d1d5db;
  font-size: 0.875rem;
}

.markdown-content .footnotes ol {
  padding-left: 1.5rem;
}

.markdown-content .footnotes li {
  margin-bottom: 0.5rem;
}

.markdown-content .footnote-ref {
  text-decoration: none;
  font-size: 0.75rem;
  vertical-align: super;
}

.markdown-content .footnote-backref {
  text-decoration: none;
  margin-left: 0.25rem;
}

/* Contenedores personalizados */
.markdown-content .warning {
  background-color: #fef2f2;
  border-left: 4px solid #ef4444;
  padding: 1rem;
  margin: 1rem 0;
  border-radius: 0.25rem;
}

/* Abreviaturas */
.markdown-content abbr {
  border-bottom: 1px dotted #6b7280;
  cursor: help;
}

/* Texto insertado y marcado */
.markdown-content ins {
  text-decoration: underline;
  background-color: #dcfce7;
}

.markdown-content mark {
  background-color: #fef08a;
  padding: 0.125rem 0.25rem;
  border-radius: 0.125rem;
}

/* Emojis */
.markdown-content img.emoji {
  height: 1em;
  width: 1em;
  margin: 0 0.05em 0 0.1em;
  vertical-align: -0.1em;
  box-shadow: none;
}

/* Separador de notas al pie */
.markdown-content .footnotes-sep {
  margin-top: 2rem;
  margin-bottom: 1rem;
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

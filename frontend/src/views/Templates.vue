<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold text-white">Configuration Templates</h1>
      <p class="text-gray-400 mt-1">Terraform, Ansible and firewall configuration templates</p>
    </div>

    <div class="flex space-x-2 flex-wrap gap-2">
      <button v-for="cat in categories" :key="cat"
        @click="selectedCategory = cat"
        class="px-4 py-2 rounded-lg text-sm font-medium transition-all border"
        :class="selectedCategory === cat
          ? 'bg-primary-600 border-primary-500 text-white'
          : 'bg-gray-800 border-gray-700 text-gray-300 hover:border-gray-600'">
        {{ cat }}
      </button>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="space-y-2">
        <button v-for="tmpl in filteredTemplates" :key="tmpl.id"
          @click="selectedTemplate = tmpl"
          class="w-full text-left p-3 rounded-lg border transition-all"
          :class="selectedTemplate?.id === tmpl.id
            ? 'border-primary-600 bg-primary-900/20 text-primary-400'
            : 'border-gray-700 hover:border-gray-600 text-gray-300 bg-gray-900'">
          <p class="font-medium text-sm">{{ tmpl.name }}</p>
          <p class="text-xs text-gray-500 mt-0.5">{{ tmpl.type }}</p>
        </button>
      </div>

      <div class="lg:col-span-2">
        <div v-if="selectedTemplate" class="card">
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-lg font-semibold text-white">{{ selectedTemplate.name }}</h2>
            <button @click="copyContent" class="btn-secondary text-xs">📋 Copy</button>
          </div>
          <p class="text-sm text-gray-400 mb-4">{{ selectedTemplate.description }}</p>
          <div class="bg-gray-950 rounded-lg border border-gray-700 p-4">
            <pre class="text-sm text-green-400 overflow-x-auto whitespace-pre-wrap font-mono">{{ selectedTemplate.content }}</pre>
          </div>
        </div>
        <div v-else class="card text-center py-12">
          <p class="text-gray-400">Select a template to view its content</p>
        </div>
      </div>
    </div>

    <div class="card">
      <h2 class="text-lg font-semibold text-white mb-4">Quick Config Generator</h2>
      <ConfigGenerator />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import ConfigGenerator from '@/components/ConfigGenerator.vue'
import templatesData from '@/assets/data/templates.json'

const templates = templatesData.templates || []
const categories = [...new Set(templates.map(t => t.type))]
const selectedCategory = ref(categories[0] || '')
const selectedTemplate = ref(null)

const filteredTemplates = computed(() => {
  if (!selectedCategory.value) return templates
  return templates.filter(t => t.type === selectedCategory.value)
})

async function copyContent() {
  if (selectedTemplate.value?.content) {
    try {
      await navigator.clipboard.writeText(selectedTemplate.value.content)
    } catch (e) {
      console.error('Copy failed', e)
    }
  }
}
</script>

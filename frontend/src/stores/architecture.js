import { defineStore } from 'pinia'
import { ref } from 'vue'
import architecturesData from '@/assets/data/architectures.json'

export const useArchitectureStore = defineStore('architecture', () => {
  const templates = ref(architecturesData.templates || [])
  const selectedTemplate = ref(null)

  function selectTemplate(template) {
    selectedTemplate.value = template
  }

  return { templates, selectedTemplate, selectTemplate }
})

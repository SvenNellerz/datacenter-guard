import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import toolsData from '@/assets/data/tools.json'

export const useToolsStore = defineStore('tools', () => {
  const tools = ref(toolsData.tools || [])
  const searchQuery = ref('')
  const selectedCategory = ref('')
  const selectedPricing = ref('')
  const loading = ref(false)

  const categories = computed(() => {
    const cats = [...new Set(tools.value.map(t => t.category))]
    return cats.sort()
  })

  const filteredTools = computed(() => {
    return tools.value.filter(tool => {
      const matchesSearch = !searchQuery.value ||
        tool.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        tool.description.toLowerCase().includes(searchQuery.value.toLowerCase())
      const matchesCategory = !selectedCategory.value || tool.category === selectedCategory.value
      const matchesPricing = !selectedPricing.value || tool.pricing === selectedPricing.value
      return matchesSearch && matchesCategory && matchesPricing
    })
  })

  function getToolById(id) {
    return tools.value.find(t => t.id === id)
  }

  return { tools, searchQuery, selectedCategory, selectedPricing, loading, categories, filteredTools, getToolById }
})

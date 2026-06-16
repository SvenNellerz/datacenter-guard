<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold text-white">Compliance Frameworks</h1>
      <p class="text-gray-400 mt-1">Framework-based checklist builder with progress tracking</p>
    </div>

    <div class="card">
      <div class="flex items-center space-x-4 flex-wrap gap-2">
        <button v-for="fw in frameworks" :key="fw.id"
          @click="selectFramework(fw)"
          class="px-4 py-2 rounded-lg text-sm font-medium transition-all border"
          :class="selectedFramework?.id === fw.id
            ? 'bg-primary-600 border-primary-500 text-white'
            : 'bg-gray-800 border-gray-700 text-gray-300 hover:border-gray-600'">
          {{ fw.name }}
        </button>
      </div>
    </div>

    <div v-if="selectedFramework" class="card">
      <div class="flex items-center justify-between mb-3">
        <h2 class="text-lg font-semibold text-white">{{ selectedFramework.name }} - Overall Progress</h2>
        <span class="text-2xl font-bold text-primary-400">{{ overallProgress }}%</span>
      </div>
      <div class="w-full bg-gray-700 rounded-full h-3">
        <div class="bg-primary-500 h-3 rounded-full transition-all"
          :style="{ width: overallProgress + '%' }"></div>
      </div>
    </div>

    <div v-if="checklist.length > 0">
      <ChecklistBuilder :categories="checklist" />
    </div>

    <div v-else class="card text-center py-12">
      <p class="text-gray-400">Select a compliance framework above to generate a checklist</p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import ChecklistBuilder from '@/components/ChecklistBuilder.vue'
import frameworksData from '@/assets/data/frameworks.json'

const frameworks = frameworksData.frameworks || []
const selectedFramework = ref(null)
const checklist = ref([])

function selectFramework(fw) {
  selectedFramework.value = fw
  checklist.value = generateChecklist(fw)
}

function generateChecklist(fw) {
  return (fw.categories || fw.controls || []).map((cat, catIdx) => ({
    name: cat.name || cat.title || `Category ${catIdx + 1}`,
    items: (cat.items || cat.controls || cat.requirements || []).map((item, idx) => ({
      id: `${fw.id}-${catIdx}-${idx}`,
      title: typeof item === 'string' ? item : (item.title || item.name || item.description || item),
      description: typeof item === 'object' ? item.description : null,
      checked: false
    }))
  }))
}

const overallProgress = computed(() => {
  if (!checklist.value.length) return 0
  const allItems = checklist.value.flatMap(c => c.items)
  if (!allItems.length) return 0
  return Math.round(allItems.filter(i => i.checked).length / allItems.length * 100)
})
</script>

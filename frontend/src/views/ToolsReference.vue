<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold text-white">Tools Reference</h1>
      <p class="text-gray-400 mt-1">Searchable database of 50+ security tools</p>
    </div>

    <div class="card">
      <div class="flex flex-wrap gap-4">
        <input v-model="toolsStore.searchQuery" type="text" placeholder="Search tools..."
          class="flex-1 min-w-48 bg-gray-800 border border-gray-700 rounded-lg px-4 py-2 text-sm text-white placeholder-gray-500 focus:outline-none focus:border-primary-500" />
        <select v-model="toolsStore.selectedCategory"
          class="bg-gray-800 border border-gray-700 rounded-lg px-3 py-2 text-sm text-white focus:outline-none focus:border-primary-500">
          <option value="">All Categories</option>
          <option v-for="cat in toolsStore.categories" :key="cat" :value="cat">{{ cat }}</option>
        </select>
        <select v-model="toolsStore.selectedPricing"
          class="bg-gray-800 border border-gray-700 rounded-lg px-3 py-2 text-sm text-white focus:outline-none focus:border-primary-500">
          <option value="">All Pricing</option>
          <option value="free">Free</option>
          <option value="freemium">Freemium</option>
          <option value="paid">Paid</option>
          <option value="enterprise">Enterprise</option>
        </select>
      </div>
    </div>

    <p class="text-sm text-gray-400">
      Showing {{ toolsStore.filteredTools.length }} of {{ toolsStore.tools.length }} tools
    </p>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <div v-for="tool in toolsStore.filteredTools" :key="tool.id"
        class="card hover:border-gray-700 transition-all cursor-pointer"
        @click="selectedTool = tool">
        <div class="flex items-start justify-between mb-3">
          <div>
            <h3 class="font-semibold text-white">{{ tool.name }}</h3>
            <p class="text-xs text-gray-500">{{ tool.vendor }}</p>
          </div>
          <span class="badge" :class="pricingClass(tool.pricing)">{{ tool.pricing }}</span>
        </div>
        <p class="text-sm text-gray-400 mb-3">{{ truncate(tool.description, 100) }}</p>
        <div class="flex items-center justify-between">
          <span class="badge bg-gray-700 text-gray-300 text-xs">{{ tool.category }}</span>
          <div class="flex items-center space-x-1">
            <span class="text-xs text-gray-500">Complexity:</span>
            <span class="text-xs" :class="complexityLabel(tool.complexity_score).class">
              {{ complexityLabel(tool.complexity_score).label }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <div v-if="selectedTool" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/60"
      @click.self="selectedTool = null">
      <div class="bg-gray-900 rounded-xl border border-gray-700 p-6 max-w-2xl w-full max-h-screen overflow-y-auto">
        <div class="flex items-start justify-between mb-4">
          <div>
            <h2 class="text-xl font-bold text-white">{{ selectedTool.name }}</h2>
            <p class="text-sm text-gray-400">{{ selectedTool.vendor }} · {{ selectedTool.category }}</p>
          </div>
          <button @click="selectedTool = null" class="text-gray-400 hover:text-white">✕</button>
        </div>
        <p class="text-gray-300 mb-4">{{ selectedTool.description }}</p>
        <div class="grid grid-cols-2 gap-4 mb-4">
          <div>
            <h4 class="text-sm font-medium text-gray-400 mb-2">Key Features</h4>
            <ul class="space-y-1">
              <li v-for="f in selectedTool.key_features" :key="f" class="text-sm text-gray-300">• {{ f }}</li>
            </ul>
          </div>
          <div>
            <h4 class="text-sm font-medium text-gray-400 mb-2">Pros</h4>
            <ul class="space-y-1">
              <li v-for="p in selectedTool.pros" :key="p" class="text-sm text-green-400">+ {{ p }}</li>
            </ul>
            <h4 class="text-sm font-medium text-gray-400 mt-3 mb-2">Cons</h4>
            <ul class="space-y-1">
              <li v-for="c in selectedTool.cons" :key="c" class="text-sm text-red-400">- {{ c }}</li>
            </ul>
          </div>
        </div>
        <div class="flex items-center space-x-4">
          <span class="badge" :class="pricingClass(selectedTool.pricing)">{{ selectedTool.pricing }}</span>
          <span class="text-sm text-gray-400">Cost: {{ selectedTool.cost_estimate }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useToolsStore } from '@/stores/tools'
import { pricingBadgeClass, complexityLabel, truncate } from '@/utils/helpers'

const toolsStore = useToolsStore()
const selectedTool = ref(null)
const pricingClass = pricingBadgeClass
</script>

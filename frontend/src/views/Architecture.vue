<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold text-white">Network Architecture</h1>
      <p class="text-gray-400 mt-1">Interactive topology builder and architecture templates</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="card">
        <h2 class="text-lg font-semibold text-white mb-4">Architecture Templates</h2>
        <div class="space-y-2">
          <button v-for="tmpl in architectureStore.templates" :key="tmpl.id"
            @click="selectTemplate(tmpl)"
            class="w-full text-left p-3 rounded-lg border transition-all"
            :class="selectedTemplate?.id === tmpl.id
              ? 'border-primary-600 bg-primary-900/20 text-primary-400'
              : 'border-gray-700 hover:border-gray-600 text-gray-300'">
            <p class="font-medium text-sm">{{ tmpl.name }}</p>
            <p class="text-xs text-gray-500 mt-0.5">{{ tmpl.size }}</p>
          </button>
        </div>
      </div>

      <div class="lg:col-span-2 card">
        <h2 class="text-lg font-semibold text-white mb-4">Network Topology</h2>
        <NetworkVisualizer />
        <div v-if="selectedTemplate" class="mt-4 p-4 bg-gray-800/50 rounded-lg">
          <h3 class="font-medium text-white mb-2">{{ selectedTemplate.name }}</h3>
          <p class="text-sm text-gray-400">{{ selectedTemplate.description }}</p>
          <div class="mt-3 flex flex-wrap gap-2">
            <span v-for="component in selectedTemplate.components" :key="component"
              class="badge bg-gray-700 text-gray-300">{{ component }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="card">
      <h2 class="text-lg font-semibold text-white mb-4">DMZ Design Patterns</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div v-for="pattern in dmzPatterns" :key="pattern.name" class="p-4 bg-gray-800/50 rounded-lg">
          <h3 class="font-medium text-white mb-2">{{ pattern.name }}</h3>
          <p class="text-sm text-gray-400">{{ pattern.description }}</p>
          <div class="mt-3">
            <p class="text-xs text-gray-500 mb-1">Use cases:</p>
            <ul class="space-y-1">
              <li v-for="uc in pattern.use_cases" :key="uc" class="text-xs text-gray-400">• {{ uc }}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import NetworkVisualizer from '@/components/NetworkVisualizer.vue'
import { useArchitectureStore } from '@/stores/architecture'

const architectureStore = useArchitectureStore()
const selectedTemplate = ref(null)

function selectTemplate(tmpl) {
  selectedTemplate.value = tmpl
  architectureStore.selectTemplate(tmpl)
}

const dmzPatterns = [
  {
    name: 'Single-Tier DMZ',
    description: 'One firewall with a single DMZ segment. Simple but suitable for small deployments.',
    use_cases: ['Small offices', 'Simple web hosting', 'Low-risk environments']
  },
  {
    name: 'Multi-Tier DMZ',
    description: 'Multiple firewalls creating distinct security zones. Recommended for enterprises.',
    use_cases: ['Enterprise data centres', 'E-commerce', 'Financial services']
  },
  {
    name: 'Screened Subnet',
    description: 'Router-based perimeter with dual firewall architecture for maximum protection.',
    use_cases: ['High-security environments', 'Government', 'Critical infrastructure']
  }
]
</script>

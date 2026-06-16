<template>
  <div class="space-y-6">
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div v-for="tool in selectedTools" :key="tool.id" class="card">
        <h3 class="font-semibold text-white mb-2">{{ tool.name }}</h3>
        <div class="space-y-3">
          <div>
            <div class="flex justify-between text-xs text-gray-400 mb-1">
              <span>Complexity</span>
              <span>{{ tool.complexity_score }}/10</span>
            </div>
            <div class="bg-gray-700 rounded-full h-2">
              <div class="bg-amber-500 h-2 rounded-full"
                :style="{ width: (tool.complexity_score * 10) + '%' }"></div>
            </div>
          </div>
          <div>
            <p class="text-xs text-gray-400 mb-1">Pricing</p>
            <span class="badge" :class="pricingClass(tool.pricing)">{{ tool.pricing }}</span>
          </div>
          <div>
            <p class="text-xs text-gray-400 mb-1">Cost Estimate</p>
            <p class="text-sm text-white">{{ tool.cost_estimate || 'Contact vendor' }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  selectedTools: { type: Array, default: () => [] }
})

function pricingClass(pricing) {
  const classes = {
    free: 'bg-green-900/50 text-green-400',
    freemium: 'bg-blue-900/50 text-blue-400',
    paid: 'bg-amber-900/50 text-amber-400',
    enterprise: 'bg-purple-900/50 text-purple-400'
  }
  return classes[pricing] || 'bg-gray-700 text-gray-400'
}
</script>

<template>
  <div class="space-y-4">
    <div v-for="category in categories" :key="category.name" class="card">
      <h3 class="font-semibold text-white mb-3">{{ category.name }}</h3>
      <div class="space-y-2">
        <label v-for="item in category.items" :key="item.id"
          class="flex items-start space-x-3 cursor-pointer group">
          <input type="checkbox" v-model="item.checked"
            class="mt-0.5 rounded border-gray-600 bg-gray-700 text-primary-600 focus:ring-primary-500" />
          <div>
            <p class="text-sm text-gray-300 group-hover:text-white transition-colors">{{ item.title }}</p>
            <p v-if="item.description" class="text-xs text-gray-500 mt-0.5">{{ item.description }}</p>
          </div>
        </label>
      </div>
      <div class="mt-3 flex items-center space-x-2">
        <div class="flex-1 bg-gray-700 rounded-full h-1.5">
          <div class="bg-primary-500 h-1.5 rounded-full transition-all"
            :style="{ width: categoryProgress(category) + '%' }"></div>
        </div>
        <span class="text-xs text-gray-400">{{ categoryProgress(category) }}%</span>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  categories: { type: Array, default: () => [] }
})

function categoryProgress(category) {
  if (!category.items?.length) return 0
  const checked = category.items.filter(i => i.checked).length
  return Math.round((checked / category.items.length) * 100)
}
</script>

<template>
  <div class="overflow-x-auto">
    <table class="w-full text-sm">
      <thead>
        <tr class="border-b border-gray-700">
          <th class="text-left py-3 px-4 text-gray-400 font-medium">Feature</th>
          <th v-for="tool in tools" :key="tool.id" class="text-left py-3 px-4 text-white font-medium">
            {{ tool.name }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="feature in features" :key="feature" class="border-b border-gray-800 hover:bg-gray-800/50">
          <td class="py-3 px-4 text-gray-400">{{ feature }}</td>
          <td v-for="tool in tools" :key="tool.id" class="py-3 px-4">
            <span v-if="toolHasFeature(tool, feature)" class="text-green-400">✓</span>
            <span v-else class="text-gray-600">—</span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
const props = defineProps({
  tools: { type: Array, default: () => [] },
  features: { type: Array, default: () => [] }
})

function toolHasFeature(tool, feature) {
  return tool.key_features?.some(f => f.toLowerCase().includes(feature.toLowerCase()))
}
</script>

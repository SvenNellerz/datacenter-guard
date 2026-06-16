<template>
  <div class="space-y-4">
    <div class="flex items-center space-x-2 flex-wrap gap-2">
      <button v-for="zone in zones" :key="zone.id"
        @click="addNode(zone)"
        class="btn-secondary text-xs">
        + {{ zone.label }}
      </button>
      <button @click="clearGraph" class="text-xs text-red-400 hover:text-red-300 px-3 py-2 rounded-lg border border-red-800 hover:border-red-600 transition-colors">
        Clear
      </button>
    </div>
    <div ref="cyContainer" class="w-full h-96 bg-gray-950 rounded-xl border border-gray-700"></div>
    <p class="text-xs text-gray-500">Click zone buttons to add nodes. Drag to reposition. Click nodes to select.</p>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const cyContainer = ref(null)
let cy = null
let nodeCount = 0

const zones = [
  { id: 'internet', label: 'Internet', color: '#ef4444', type: 'external' },
  { id: 'dmz', label: 'DMZ', color: '#f59e0b', type: 'dmz' },
  { id: 'firewall', label: 'Firewall', color: '#3b82f6', type: 'firewall' },
  { id: 'internal', label: 'Internal Network', color: '#10b981', type: 'internal' },
  { id: 'management', label: 'Management', color: '#8b5cf6', type: 'management' },
  { id: 'server', label: 'Server Zone', color: '#06b6d4', type: 'server' }
]

onMounted(async () => {
  try {
    const cytoscape = (await import('cytoscape')).default
    cy = cytoscape({
      container: cyContainer.value,
      style: [
        {
          selector: 'node',
          style: {
            'background-color': 'data(color)',
            'label': 'data(label)',
            'color': '#ffffff',
            'text-valign': 'center',
            'font-size': '11px',
            'width': 80,
            'height': 40,
            'shape': 'roundrectangle',
            'text-wrap': 'wrap'
          }
        },
        {
          selector: 'edge',
          style: {
            'width': 2,
            'line-color': '#4b5563',
            'target-arrow-color': '#4b5563',
            'target-arrow-shape': 'triangle',
            'curve-style': 'bezier'
          }
        },
        {
          selector: ':selected',
          style: {
            'border-width': 3,
            'border-color': '#3b82f6'
          }
        }
      ],
      layout: { name: 'preset' }
    })
  } catch (e) {
    console.warn('Cytoscape not available:', e.message)
  }
})

onUnmounted(() => {
  if (cy) cy.destroy()
})

function addNode(zone) {
  if (!cy) return
  nodeCount++
  const x = 100 + (nodeCount % 5) * 130
  const y = 80 + Math.floor(nodeCount / 5) * 100
  cy.add({
    group: 'nodes',
    data: { id: `${zone.id}-${nodeCount}`, label: zone.label, color: zone.color },
    position: { x, y }
  })
}

function clearGraph() {
  if (cy) {
    cy.elements().remove()
    nodeCount = 0
  }
}
</script>

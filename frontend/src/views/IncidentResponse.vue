<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold text-white">Incident Response</h1>
      <p class="text-gray-400 mt-1">IR playbooks, procedures and communication templates</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="space-y-2">
        <h2 class="text-sm font-medium text-gray-400 uppercase tracking-wider mb-3">Playbooks</h2>
        <button v-for="pb in playbooks" :key="pb.id"
          @click="selectedPlaybook = pb"
          class="w-full text-left p-3 rounded-lg border transition-all"
          :class="selectedPlaybook?.id === pb.id
            ? 'border-red-700 bg-red-900/20 text-red-400'
            : 'border-gray-700 hover:border-gray-600 text-gray-300 bg-gray-900'">
          <div class="flex items-center space-x-2">
            <span>{{ pb.icon || '🚨' }}</span>
            <div>
              <p class="font-medium text-sm">{{ pb.name }}</p>
              <span class="text-xs badge mt-1" :class="severityClass(pb.severity)">{{ pb.severity }}</span>
            </div>
          </div>
        </button>
      </div>

      <div class="lg:col-span-2">
        <div v-if="selectedPlaybook" class="space-y-4">
          <div class="card">
            <h2 class="text-xl font-bold text-white mb-2">{{ selectedPlaybook.name }}</h2>
            <p class="text-gray-400">{{ selectedPlaybook.description }}</p>
          </div>

          <div class="card">
            <h3 class="font-semibold text-white mb-4">Response Steps</h3>
            <div class="space-y-3">
              <div v-for="(step, idx) in selectedPlaybook.steps" :key="idx"
                class="flex items-start space-x-3 p-3 bg-gray-800/50 rounded-lg">
                <div class="flex-shrink-0 w-6 h-6 bg-primary-600 rounded-full flex items-center justify-center text-xs font-bold text-white">
                  {{ idx + 1 }}
                </div>
                <div>
                  <p class="text-sm font-medium text-white">{{ step.title || step }}</p>
                  <p v-if="step.description" class="text-xs text-gray-400 mt-0.5">{{ step.description }}</p>
                </div>
              </div>
            </div>
          </div>

          <div v-if="selectedPlaybook.contacts" class="card">
            <h3 class="font-semibold text-white mb-3">Escalation Contacts</h3>
            <div class="space-y-2">
              <div v-for="contact in selectedPlaybook.contacts" :key="contact.role"
                class="flex items-center justify-between p-3 bg-gray-800/50 rounded-lg">
                <span class="text-sm text-gray-300">{{ contact.role }}</span>
                <span class="text-sm text-primary-400">{{ contact.contact }}</span>
              </div>
            </div>
          </div>
        </div>
        <div v-else class="card text-center py-12">
          <p class="text-gray-400">Select a playbook to view the response procedure</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import playbooksData from '@/assets/data/playbooks.json'

const playbooks = playbooksData.playbooks || []
const selectedPlaybook = ref(null)

function severityClass(severity) {
  const classes = {
    critical: 'bg-red-900/50 text-red-400',
    high: 'bg-orange-900/50 text-orange-400',
    medium: 'bg-amber-900/50 text-amber-400',
    low: 'bg-green-900/50 text-green-400'
  }
  return classes[severity] || 'bg-gray-700 text-gray-400'
}
</script>

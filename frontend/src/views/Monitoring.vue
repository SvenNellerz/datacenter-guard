<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold text-white">Monitoring & Logging</h1>
      <p class="text-gray-400 mt-1">SIEM solutions, alert rules, and monitoring architecture</p>
    </div>

    <div class="card">
      <h2 class="text-lg font-semibold text-white mb-4">SIEM Solutions Comparison</h2>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-700">
              <th class="text-left py-3 px-4 text-gray-400">Tool</th>
              <th class="text-left py-3 px-4 text-gray-400">Pricing</th>
              <th class="text-left py-3 px-4 text-gray-400">Complexity</th>
              <th class="text-left py-3 px-4 text-gray-400">Best For</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="tool in siemTools" :key="tool.id" class="border-b border-gray-800 hover:bg-gray-800/50">
              <td class="py-3 px-4 font-medium text-white">{{ tool.name }}</td>
              <td class="py-3 px-4">
                <span class="badge" :class="pricingClass(tool.pricing)">{{ tool.pricing }}</span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2">
                  <div class="w-20 bg-gray-700 rounded-full h-1.5">
                    <div class="h-1.5 rounded-full bg-amber-500" :style="{ width: (tool.complexity_score * 10) + '%' }"></div>
                  </div>
                  <span class="text-xs text-gray-400">{{ tool.complexity_score }}/10</span>
                </div>
              </td>
              <td class="py-3 px-4 text-gray-400 text-xs">{{ tool.use_cases?.[0] }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div class="card">
      <h2 class="text-lg font-semibold text-white mb-4">Alert Threshold Templates</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div v-for="alert in alertTemplates" :key="alert.name" class="p-4 bg-gray-800/50 rounded-lg border border-gray-700">
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-medium text-white text-sm">{{ alert.name }}</h3>
            <span class="badge text-xs" :class="severityClass(alert.severity)">{{ alert.severity }}</span>
          </div>
          <p class="text-xs text-gray-400 font-mono">{{ alert.rule }}</p>
        </div>
      </div>
    </div>

    <div class="card">
      <h2 class="text-lg font-semibold text-white mb-4">Log Retention Calculator</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div>
          <label class="block text-sm text-gray-400 mb-2">Daily Log Volume (GB)</label>
          <input v-model.number="logVolume" type="number" min="1"
            class="w-full bg-gray-800 border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-primary-500" />
        </div>
        <div>
          <label class="block text-sm text-gray-400 mb-2">Retention Period</label>
          <select v-model="retentionDays" class="w-full bg-gray-800 border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-primary-500">
            <option :value="30">30 days</option>
            <option :value="90">90 days</option>
            <option :value="180">180 days</option>
            <option :value="365">1 year</option>
            <option :value="730">2 years</option>
          </select>
        </div>
        <div class="flex items-end">
          <div class="w-full p-3 bg-primary-900/30 rounded-lg border border-primary-800">
            <p class="text-xs text-gray-400">Estimated Storage</p>
            <p class="text-2xl font-bold text-primary-400">{{ storageEstimate }} TB</p>
            <p class="text-xs text-gray-500">with compression ~{{ Math.round(storageEstimate * 0.3 * 10) / 10 }} TB</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { pricingBadgeClass } from '@/utils/helpers'
import toolsData from '@/assets/data/tools.json'

const siemTools = toolsData.tools?.filter(t => t.category === 'SIEM').slice(0, 8) || []
const logVolume = ref(10)
const retentionDays = ref(90)

const storageEstimate = computed(() => {
  return Math.round(logVolume.value * retentionDays.value / 1024 * 10) / 10
})

const pricingClass = pricingBadgeClass

function severityClass(severity) {
  const classes = {
    critical: 'bg-red-900/50 text-red-400',
    high: 'bg-orange-900/50 text-orange-400',
    medium: 'bg-amber-900/50 text-amber-400',
    low: 'bg-green-900/50 text-green-400'
  }
  return classes[severity] || 'bg-gray-700 text-gray-400'
}

const alertTemplates = [
  { name: 'Brute Force Detection', severity: 'high', rule: 'failed_logins > 5 in 60s from same IP' },
  { name: 'Privilege Escalation', severity: 'critical', rule: 'sudo | su events with unusual user' },
  { name: 'Data Exfiltration', severity: 'critical', rule: 'outbound_bytes > 1GB from single host' },
  { name: 'After-Hours Access', severity: 'medium', rule: 'login_time outside 06:00-22:00 local' },
  { name: 'New Admin Account', severity: 'high', rule: 'useradd -G sudo or wheel group change' },
  { name: 'Firewall Rule Change', severity: 'high', rule: 'iptables | ufw | firewall-cmd change' }
]
</script>

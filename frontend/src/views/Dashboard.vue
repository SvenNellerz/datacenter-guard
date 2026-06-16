<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold text-white">Security Design Reference</h1>
      <p class="text-gray-400 mt-1">Comprehensive guide for secure data centre implementation</p>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <div v-for="stat in stats" :key="stat.label" class="card text-center">
        <div class="text-3xl mb-2">{{ stat.icon }}</div>
        <div class="text-2xl font-bold text-white">{{ stat.value }}</div>
        <div class="text-sm text-gray-400 mt-1">{{ stat.label }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      <RouterLink v-for="item in navCards" :key="item.path" :to="item.path"
        class="card hover:border-primary-700 transition-all hover:bg-gray-800/50 group cursor-pointer">
        <div class="text-3xl mb-3">{{ item.icon }}</div>
        <h3 class="font-semibold text-white group-hover:text-primary-400 transition-colors">{{ item.title }}</h3>
        <p class="text-sm text-gray-400 mt-1">{{ item.description }}</p>
      </RouterLink>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="card">
        <h2 class="text-lg font-semibold text-white mb-4">🔥 Featured Tools</h2>
        <div class="space-y-3">
          <div v-for="tool in featuredTools" :key="tool.id" class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium text-white">{{ tool.name }}</p>
              <p class="text-xs text-gray-500">{{ tool.category }}</p>
            </div>
            <span class="badge" :class="pricingClass(tool.pricing)">{{ tool.pricing }}</span>
          </div>
        </div>
      </div>
      <div class="card">
        <h2 class="text-lg font-semibold text-white mb-4">📋 Compliance Frameworks</h2>
        <div class="space-y-3">
          <div v-for="fw in frameworks" :key="fw.id"
            class="flex items-center justify-between p-3 bg-gray-800/50 rounded-lg">
            <span class="text-sm font-medium text-white">{{ fw.name }}</span>
            <RouterLink to="/compliance" class="text-xs text-primary-400 hover:text-primary-300">View →</RouterLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { pricingBadgeClass } from '@/utils/helpers'
import toolsData from '@/assets/data/tools.json'
import frameworksData from '@/assets/data/frameworks.json'
import templatesData from '@/assets/data/templates.json'
import playbooksData from '@/assets/data/playbooks.json'

const toolCount = toolsData.tools?.length ?? 0
const frameworkCount = frameworksData.frameworks?.length ?? 0
const templateCount = templatesData.templates?.length ?? 0
const playbookCount = playbooksData.playbooks?.length ?? 0

const stats = [
  { icon: '🛠️', value: toolCount, label: 'Security Tools' },
  { icon: '📋', value: frameworkCount, label: 'Frameworks' },
  { icon: '📦', value: templateCount, label: 'Templates' },
  { icon: '🚨', value: playbookCount, label: 'IR Playbooks' }
]

const navCards = [
  { path: '/architecture', icon: '🏗️', title: 'Architecture', description: 'Visual network topology builder with DMZ patterns' },
  { path: '/pam', icon: '🔐', title: 'PAM', description: 'Privileged access management tools and workflows' },
  { path: '/monitoring', icon: '📈', title: 'Monitoring', description: 'SIEM solutions and alert configuration' },
  { path: '/compliance', icon: '✅', title: 'Compliance', description: 'Framework-based checklist builder' },
  { path: '/tools', icon: '🛠️', title: 'Tools Reference', description: `Searchable database of ${toolCount} security tools` },
  { path: '/templates', icon: '📋', title: 'Templates', description: 'Terraform, Ansible and firewall templates' },
  { path: '/incident-response', icon: '🚨', title: 'Incident Response', description: 'IR playbooks and procedures' }
]

const featuredTools = toolsData.tools?.slice(0, 5) || []
const frameworks = frameworksData.frameworks?.slice(0, 5) || []

const pricingClass = pricingBadgeClass
</script>

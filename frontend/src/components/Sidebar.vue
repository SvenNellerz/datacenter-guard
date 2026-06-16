<template>
  <aside class="fixed left-0 top-16 bottom-0 w-64 bg-gray-900 border-r border-gray-800 overflow-y-auto z-40">
    <nav class="p-4 space-y-1">
      <RouterLink
        v-for="item in navItems"
        :key="item.path"
        :to="item.path"
        class="flex items-center space-x-3 px-3 py-2.5 rounded-lg text-gray-400 hover:text-white hover:bg-gray-800 transition-all group"
        active-class="bg-primary-900/50 text-primary-400 border border-primary-800/50"
      >
        <span class="text-lg">{{ item.icon }}</span>
        <span class="font-medium text-sm">{{ item.label }}</span>
      </RouterLink>
    </nav>
    <div class="p-4 border-t border-gray-800 mt-4">
      <p class="text-xs text-gray-600 text-center">DataCenter Guard v1.0</p>
    </div>
  </aside>
</template>

<script setup>
const navItems = [
  { path: '/', icon: '📊', label: 'Dashboard' },
  { path: '/architecture', icon: '🏗️', label: 'Architecture' },
  { path: '/pam', icon: '🔐', label: 'PAM' },
  { path: '/monitoring', icon: '📈', label: 'Monitoring & Logging' },
  { path: '/compliance', icon: '✅', label: 'Compliance' },
  { path: '/tools', icon: '🛠️', label: 'Tools Reference' },
  { path: '/templates', icon: '📋', label: 'Config Templates' },
  { path: '/incident-response', icon: '🚨', label: 'Incident Response' }
]
</script>

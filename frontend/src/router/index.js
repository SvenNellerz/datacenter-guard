import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', name: 'Dashboard', component: () => import('@/views/Dashboard.vue') },
  { path: '/architecture', name: 'Architecture', component: () => import('@/views/Architecture.vue') },
  { path: '/pam', name: 'PAM', component: () => import('@/views/PAM.vue') },
  { path: '/monitoring', name: 'Monitoring', component: () => import('@/views/Monitoring.vue') },
  { path: '/compliance', name: 'Compliance', component: () => import('@/views/Compliance.vue') },
  { path: '/tools', name: 'ToolsReference', component: () => import('@/views/ToolsReference.vue') },
  { path: '/templates', name: 'Templates', component: () => import('@/views/Templates.vue') },
  { path: '/incident-response', name: 'IncidentResponse', component: () => import('@/views/IncidentResponse.vue') }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router

<template>
  <div class="space-y-6">
    <div>
      <h1 class="text-3xl font-bold text-white">Privileged Access Management</h1>
      <p class="text-gray-400 mt-1">Tools, workflows and best practices for PAM implementation</p>
    </div>

    <div class="card">
      <h2 class="text-lg font-semibold text-white mb-4">PAM Tool Comparison</h2>
      <ToolComparison :tools="pamTools" :features="pamFeatures" />
    </div>

    <div class="card">
      <h2 class="text-lg font-semibold text-white mb-4">PAM Implementation Workflow</h2>
      <div class="flex flex-wrap gap-2 items-center">
        <div v-for="(step, idx) in workflowSteps" :key="step.title"
          class="flex items-center">
          <div class="flex items-center p-3 bg-gray-800 rounded-lg border border-gray-700 min-w-max">
            <span class="text-lg mr-2">{{ step.icon }}</span>
            <div>
              <p class="text-xs text-gray-400">Step {{ idx + 1 }}</p>
              <p class="text-sm font-medium text-white">{{ step.title }}</p>
            </div>
          </div>
          <span v-if="idx < workflowSteps.length - 1" class="mx-2 text-gray-600">→</span>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="card">
        <h2 class="text-lg font-semibold text-white mb-4">SSH Key Management</h2>
        <ul class="space-y-2">
          <li v-for="practice in sshPractices" :key="practice" class="flex items-start space-x-2">
            <span class="text-green-400 mt-0.5">✓</span>
            <span class="text-sm text-gray-300">{{ practice }}</span>
          </li>
        </ul>
      </div>
      <div class="card">
        <h2 class="text-lg font-semibold text-white mb-4">Session Recording Configuration</h2>
        <div class="bg-gray-950 rounded-lg p-4">
          <pre class="text-sm text-green-400 overflow-x-auto font-mono">{{ sessionRecordingConfig }}</pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import ToolComparison from '@/components/ToolComparison.vue'
import toolsData from '@/assets/data/tools.json'

const pamTools = toolsData.tools?.filter(t => t.category === 'PAM').slice(0, 4) || []

const pamFeatures = [
  'MFA Support', 'Session Recording', 'Password Vaulting', 'SSH Key Management',
  'API Access', 'LDAP Integration', 'RBAC', 'Audit Logging'
]

const workflowSteps = [
  { icon: '📋', title: 'Discovery' },
  { icon: '🏷️', title: 'Classification' },
  { icon: '🔐', title: 'Onboarding' },
  { icon: '🔄', title: 'Rotation' },
  { icon: '📊', title: 'Monitoring' },
  { icon: '📝', title: 'Audit' }
]

const sshPractices = [
  'Use Ed25519 or RSA-4096 key pairs',
  'Enforce passphrase on all private keys',
  'Store keys in a dedicated secrets manager',
  'Rotate keys every 90 days',
  'Disable password authentication on SSH servers',
  'Use certificate-based SSH where possible',
  'Maintain an inventory of all authorized keys',
  'Remove stale/orphaned keys immediately'
]

const sessionRecordingConfig = `# HashiCorp Vault SSH Helper Config
vault_addr = "https://vault.example.com:8200"
ssh_mount_point = "ssh"
allowed_roles = "admin,operator"
tty_name = "pts/0"

# Enable session recording
[session_recording]
  enabled = true
  storage_path = "/var/log/vault/sessions"
  format = "asciicast"
  max_ttl = "4h"`
</script>

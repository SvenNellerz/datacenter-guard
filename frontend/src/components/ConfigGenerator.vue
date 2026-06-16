<template>
  <div class="space-y-4">
    <div class="flex items-center space-x-4">
      <select v-model="selectedType"
        class="bg-gray-800 border border-gray-700 rounded-lg px-3 py-2 text-sm text-white focus:outline-none focus:border-primary-500">
        <option value="terraform">Terraform</option>
        <option value="ansible">Ansible</option>
        <option value="firewall">Firewall Rules</option>
      </select>
      <button @click="generate" class="btn-primary text-sm">Generate</button>
      <button @click="copyToClipboard" class="btn-secondary text-sm">Copy</button>
    </div>
    <div class="bg-gray-950 rounded-lg border border-gray-700 p-4">
      <pre class="text-sm text-green-400 overflow-x-auto whitespace-pre-wrap font-mono">{{ generatedConfig }}</pre>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  template: { type: Object, default: null }
})

const selectedType = ref('terraform')
const generatedConfig = ref('# Select a type and click Generate to create a configuration template')

function generate() {
  if (!props.template) {
    generatedConfig.value = getDefaultTemplate(selectedType.value)
    return
  }
  generatedConfig.value = props.template[selectedType.value] || getDefaultTemplate(selectedType.value)
}

function getDefaultTemplate(type) {
  const templates = {
    terraform: `# Terraform Configuration\nterraform {\n  required_version = ">= 1.0"\n  required_providers {\n    aws = {\n      source  = "hashicorp/aws"\n      version = "~> 5.0"\n    }\n  }\n}\n\nprovider "aws" {\n  region = var.aws_region\n}\n`,
    ansible: `---\n- name: Security Hardening Playbook\n  hosts: all\n  become: yes\n  tasks:\n    - name: Update all packages\n      package:\n        name: '*'\n        state: latest\n`,
    firewall: `# Firewall Rules\n# Allow established connections\n-A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT\n# Allow SSH (restrict to management network)\n-A INPUT -p tcp --dport 22 -s 10.0.0.0/8 -j ACCEPT\n# Drop all other input\n-A INPUT -j DROP\n`
  }
  return templates[type] || '# No template available'
}

async function copyToClipboard() {
  try {
    await navigator.clipboard.writeText(generatedConfig.value)
  } catch (e) {
    console.error('Copy failed', e)
  }
}
</script>

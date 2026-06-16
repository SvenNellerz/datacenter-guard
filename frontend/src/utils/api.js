import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: { 'Content-Type': 'application/json' }
})

api.interceptors.response.use(
  response => response.data,
  error => {
    console.error('API Error:', error.message)
    return Promise.reject(error)
  }
)

export const toolsApi = {
  getAll: (params) => api.get('/tools', { params }),
  getById: (id) => api.get(`/tools/${id}`)
}

export const frameworksApi = {
  getAll: () => api.get('/frameworks'),
  getChecklist: (id) => api.get(`/frameworks/${id}/checklist`)
}

export const configsApi = {
  getTemplates: () => api.get('/configs/templates'),
  generate: (data) => api.post('/configs/generate', data)
}

export const playbooksApi = {
  getAll: () => api.get('/playbooks')
}

export const architecturesApi = {
  getAll: () => api.get('/architectures')
}

export default api

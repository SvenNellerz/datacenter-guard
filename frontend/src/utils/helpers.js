export function formatDate(date) {
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric', month: 'short', day: 'numeric'
  })
}

export function truncate(str, length = 100) {
  if (!str) return ''
  return str.length > length ? str.substring(0, length) + '...' : str
}

export function slugify(str) {
  return str.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/(^-|-$)/g, '')
}

export function complexityLabel(score) {
  if (score <= 3) return { label: 'Low', class: 'text-green-400' }
  if (score <= 6) return { label: 'Medium', class: 'text-amber-400' }
  return { label: 'High', class: 'text-red-400' }
}

export function pricingBadgeClass(pricing) {
  const map = {
    free: 'bg-green-900/50 text-green-400 border border-green-800',
    freemium: 'bg-blue-900/50 text-blue-400 border border-blue-800',
    paid: 'bg-amber-900/50 text-amber-400 border border-amber-800',
    enterprise: 'bg-purple-900/50 text-purple-400 border border-purple-800'
  }
  return map[pricing] || 'bg-gray-700 text-gray-400'
}

export function groupBy(array, key) {
  return array.reduce((groups, item) => {
    const group = item[key]
    groups[group] = groups[group] || []
    groups[group].push(item)
    return groups
  }, {})
}

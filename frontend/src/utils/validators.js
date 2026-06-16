export function validateIpAddress(ip) {
  const regex = /^(\d{1,3}\.){3}\d{1,3}(\/\d{1,2})?$/
  return regex.test(ip)
}

export function validateCidr(cidr) {
  const regex = /^(\d{1,3}\.){3}\d{1,3}\/\d{1,2}$/
  return regex.test(cidr)
}

export function validatePort(port) {
  const num = parseInt(port)
  return !isNaN(num) && num >= 1 && num <= 65535
}

export function validateRequired(value) {
  return value !== null && value !== undefined && value !== ''
}

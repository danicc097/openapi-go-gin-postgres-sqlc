export function removePrefix(str: string, prefix: string) {
  if (!str) {
    return ''
  }
  if (str.startsWith(prefix)) {
    return str.slice(prefix.length)
  }
  return str
}

export function sentenceCase(str) {
  if (!str) {
    return ''
  }
  return str.charAt(0).toUpperCase() + str.substr(1).toLowerCase()
}

export function nameInitials(name: string) {
  return name
    .split(' ')
    .map((n) => (n[0] || '').toUpperCase())
    .join('')
}

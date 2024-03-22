import _ from 'lodash'

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
  return _.upperFirst(_.lowerCase(str))
}

export function nameInitials(name: string) {
  return name
    .split(' ')
    .map((n) => (n[0] || '').toUpperCase())
    .join('')
}

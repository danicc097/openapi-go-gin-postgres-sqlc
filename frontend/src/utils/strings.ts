export function removePrefix(str: string, prefix: string) {
  if (str.startsWith(prefix)) {
    return str.slice(prefix.length)
  }
  return str
}

export function sentenceCase(str) {
  return str.charAt(0).toUpperCase() + str.substr(1).toLowerCase()
}

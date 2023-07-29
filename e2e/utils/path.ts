import CONFIG from '../config.json'

export function apiPath(path: string | null) {
  const port = CONFIG.API_PORT?.length > 0 ? `:${CONFIG.API_PORT}` : ''
  return `https://${CONFIG.DOMAIN}${port}${CONFIG.API_PREFIX}${CONFIG.API_VERSION}${path ?? ''}`
}

export function frontendPath(path: string | null) {
  const port = CONFIG.API_PORT?.length > 0 ? `:${CONFIG.API_PORT}` : ''
  return `https://${CONFIG.DOMAIN}${port}${path ?? ''}`
}

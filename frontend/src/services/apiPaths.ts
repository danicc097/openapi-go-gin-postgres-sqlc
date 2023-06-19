import CONFIG from 'src/config'
import type { paths } from 'src/types/schema'

export function apiPath(path: keyof paths | null) {
  const port = CONFIG.API_PORT?.length > 0 ? `:${CONFIG.API_PORT}` : ''
  return `https://${CONFIG.DOMAIN}${port}${CONFIG.API_PREFIX}${CONFIG.API_VERSION}${path ?? ''}`
}

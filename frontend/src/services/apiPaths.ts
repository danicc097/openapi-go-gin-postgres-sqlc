import config from '@config'
import type { paths } from 'src/types/schema'

export function apiPath(path: keyof paths) {
  const port = config.API_PORT?.length > 0 ? ':' + config.API_PORT : ''
  return `https://${config.DOMAIN}${port}${config.API_PREFIX}${config.API_VERSION}${path}`
}

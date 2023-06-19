import ROLES from '../roles.json'
import type { Role } from 'src/gen/model'

export default ROLES as unknown as {
  [key in Role]: typeof ROLES[keyof typeof ROLES]
}

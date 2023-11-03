import OPERATION_AUTH from '../operationAuth.gen.json'
import type { Role, Scopes } from 'src/gen/model'
import { operations } from 'src/types/schema'

export default OPERATION_AUTH as unknown as {
  [key in keyof operations]: {
    scopes: Scopes
    role: Role
    requiresAuthentication: boolean
  }
}

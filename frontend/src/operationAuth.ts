import OPERATION_AUTH from '../operationAuth.gen.json'
import type { Role, Scopes } from 'src/gen/model'
import { operations } from 'src/types/schema'

export type OperationAuth = {
  scopes: Scopes
  role: Role
  requiresAuthentication: boolean
}

export default OPERATION_AUTH as unknown as {
  [key in keyof operations]: OperationAuth
}

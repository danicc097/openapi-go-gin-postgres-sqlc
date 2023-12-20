import { ROLES } from 'src/config'
import { WorkItemRole, type Role, type Scopes, type User } from 'src/gen/model'
import { keys } from 'src/utils/object'

interface IsAuthorizedParams {
  user: User
  requiredRole?: Role | null
  requiredScopes?: Scopes | null
}

// TODO: isAuthorized mapped against @operationAuth. would need to generate a wrapper
// for orval's react query per operation id that checks the current user state and its
// scopes, role and requiresAuthentication before making the request.

export function isAuthorized({ user, requiredRole = null, requiredScopes = null }: IsAuthorizedParams): boolean {
  if (requiredRole !== null) {
    if (ROLES[user.role].rank < ROLES[requiredRole].rank) {
      return false
    }
  }
  if (requiredScopes !== null) {
    for (const scope of requiredScopes) {
      if (!user.scopes.includes(scope)) {
        return false
      }
    }
  }

  return true
}

export const WORK_ITEM_ROLES: WorkItemRole[] = keys(WorkItemRole)

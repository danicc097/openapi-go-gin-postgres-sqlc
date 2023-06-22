import ROLES from 'src/roles'
import SCOPES from 'src/scopes'
import type { Role, Scopes, UserResponse } from 'src/gen/model'

interface IsAuthorizedParams {
  user: UserResponse
  requiredRole?: Role
  requiredScopes?: Scopes
}

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

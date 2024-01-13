import { ROLES } from 'src/config'
import { WorkItemRole, type Role, type Scopes, type User } from 'src/gen/model'
import { apiPath } from 'src/services/apiPaths'
import { joinWithAnd } from 'src/utils/format'
import { keys } from 'src/utils/object'

interface IsAuthorizedParams {
  user?: User
  requiredRole?: Role | null
  requiredScopes?: Scopes | null
}

export interface IsAuthorizedResult {
  isAuthorized: boolean
  missingScopes?: Scopes
  missingRole?: Role
  errorMessage?: string
}
/* TODO: isAuthorized mapped against @operationAuth. would need to generate a wrapper
   for orval's react query per operation id that checks the current user state and its
  scopes, role and requiresAuthentication before making the request. */
export function isAuthorized({
  user,
  requiredRole = null,
  requiredScopes = null,
}: IsAuthorizedParams): IsAuthorizedResult {
  const result: IsAuthorizedResult = {
    isAuthorized: false,
  }

  if (!user) {
    result.errorMessage = 'User not authenticated. Please log in.'
    return result
  }

  if (requiredRole !== null) {
    if (ROLES[user.role].rank < ROLES[requiredRole].rank) {
      result.missingRole = requiredRole
    }
  }

  if (requiredScopes !== null) {
    const missingScopes: Scopes = []
    for (const scope of requiredScopes) {
      if (!user.scopes.includes(scope)) {
        missingScopes.push(scope)
      }
    }

    if (missingScopes.length > 0) {
      result.missingScopes = missingScopes
    }
  }

  if (result.missingRole || result.missingScopes) {
    result.errorMessage = getUnauthorizedMessage(result)
    result.isAuthorized = false
    return result
  }

  result.isAuthorized = true
  return result
}

export const WORK_ITEM_ROLES: WorkItemRole[] = keys(WorkItemRole)

export const redirectToAuthLogin = () => {
  window.location.replace(
    `${apiPath('/auth/myprovider/login')}?auth-redirect-uri=${encodeURIComponent(window.location.href)}`,
  )
}

const getUnauthorizedMessage = (authResult: IsAuthorizedResult): string => {
  if (!authResult.isAuthorized) {
    const messages: string[] = []

    if (authResult.missingRole) {
      messages.push(`missing role ${authResult.missingRole}`)
    }

    if (authResult.missingScopes && authResult.missingScopes.length > 0) {
      const scopeMessage = joinWithAnd(authResult.missingScopes)
      messages.push(`missing scopes ${scopeMessage}`)
    }

    return messages.join(', ')
  }

  return ''
}

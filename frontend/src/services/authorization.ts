import { OperationAuth, ROLES } from 'src/config'
import { WorkItemRole, type Role, type Scopes, type UserResponse } from 'src/gen/model'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { apiPath } from 'src/services/apiPaths'
import { joinWithAnd } from 'src/utils/format'
import { keys } from 'src/utils/object'

interface CheckAuthorizationParams {
  user?: UserResponse
  requiredRole?: Role | null
  requiredScopes?: Scopes | null
}

export interface Authorization {
  authorized: boolean
  missingScopes?: Scopes
  missingRole?: Role
  errorMessage?: string
}
/* TODO: checkAuthorization mapped against @operationAuth in . would need to generate a wrapper
   for orval's react query per operation id that checks the current user state and its
  scopes, role and requiresAuthentication before making the request to prevent useless calls
  in case frontend does not reimplement all auth logic in client. */
export function checkAuthorization({
  user,
  requiredRole = null,
  requiredScopes = null,
}: CheckAuthorizationParams): Authorization {
  const result: Authorization = {
    authorized: false,
  }

  if (!user) {
    result.errorMessage = 'User not authenticated. Please log in'
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
    result.authorized = false
    return result
  }

  result.authorized = true
  return result
}

export const WORK_ITEM_ROLES: WorkItemRole[] = keys(WorkItemRole)

export const redirectToAuthLogin = () => {
  window.location.replace(
    `${apiPath('/auth/myprovider/login')}?auth-redirect-uri=${encodeURIComponent(window.location.href)}`,
  )
}

const getUnauthorizedMessage = (authResult: Authorization): string => {
  if (!authResult.authorized) {
    const messages: string[] = []

    if (authResult.missingRole) {
      messages.push(`missing role ${authResult.missingRole}`)
    }

    if (authResult.missingScopes && authResult.missingScopes.length > 0) {
      const quotedScopes = authResult.missingScopes.map((s) => `"${s}"`)
      const scopeMessage = joinWithAnd(quotedScopes)
      messages.push(`missing scopes ${scopeMessage}`)
    }

    return joinWithAnd(messages)
  }

  return ''
}

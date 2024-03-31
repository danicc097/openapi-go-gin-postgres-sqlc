import { OperationAuth } from 'src/config'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { checkAuthorization } from 'src/services/authorization'

// TODO: frontend.gen will generate an orval wrapper that makes use of this hook
// for each call to api
export const useOperationAuth = (operationAuth: OperationAuth): boolean => {
  const { user } = useAuthenticatedUser()

  return checkAuthorization({ user, requiredRole: operationAuth.role, requiredScopes: operationAuth.scopes }).authorized
}

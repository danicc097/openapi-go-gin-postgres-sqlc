import roles from '@roles'
import scopes from '@scopes'
import type { UserResponse } from 'src/gen/model'
import { getGetCurrentUserMock, getUserMSW } from 'src/gen/user/user.msw'
import { isAuthorized } from 'src/services/authorization'
import { describe, expect, it, test } from 'vitest'

describe('roles and scopes', async () => {
  const user = getGetCurrentUserMock() as UserResponse
  test('role', () => {
    user.role = 'user'
    expect(isAuthorized({ user, requiredRole: 'admin' })).toBe(false)
    expect(isAuthorized({ user, requiredRole: 'user' })).toBe(true)
  })
  test('scopes', () => {
    user.scopes = ['scopes:write']
    expect(isAuthorized({ user, requiredScopes: ['team-settings:write'] })).toBe(false)
    expect(isAuthorized({ user, requiredScopes: ['team-settings:write', 'scopes:write'] })).toBe(false)
    expect(isAuthorized({ user, requiredScopes: ['scopes:write'] })).toBe(true)
  })
  test('roles and scopes', () => {
    user.role = 'user'
    user.scopes = ['scopes:write']
    expect(isAuthorized({ user, requiredScopes: ['team-settings:write'], requiredRole: 'user' })).toBe(false)
    expect(isAuthorized({ user, requiredScopes: ['scopes:write'], requiredRole: 'admin' })).toBe(false)
    expect(isAuthorized({ user, requiredScopes: ['scopes:write'], requiredRole: 'user' })).toBe(true)
  })
  test('default authorized', () => {
    expect(isAuthorized({ user })).toBe(true)
  })
})

import type { User } from 'src/gen/model'
import { isAuthorized } from 'src/services/authorization'
import { describe, expect, it, test } from 'vitest'

describe('roles and scopes', async () => {
  const user = {} as User

  test('role', () => {
    user.role = 'user'
    const resultAdmin = isAuthorized({ user, requiredRole: 'admin' })
    const resultUser = isAuthorized({ user, requiredRole: 'user' })

    expect(resultAdmin.isAuthorized).toBe(false)
    expect(resultAdmin.missingRole).toBe('admin')

    expect(resultUser.isAuthorized).toBe(true)
    expect(resultUser.missingRole).toBeUndefined()
  })

  test('scopes', () => {
    user.scopes = ['scopes:write']

    const resultWrite = isAuthorized({ user, requiredScopes: ['team-settings:write'] })
    const resultBoth = isAuthorized({ user, requiredScopes: ['team-settings:write', 'scopes:write'] })
    const resultValid = isAuthorized({ user, requiredScopes: ['scopes:write'] })

    expect(resultWrite.isAuthorized).toBe(false)
    expect(resultWrite.missingScopes).toEqual(['team-settings:write'])

    expect(resultBoth.isAuthorized).toBe(false)
    expect(resultBoth.missingScopes).toEqual(['team-settings:write'])

    expect(resultValid.isAuthorized).toBe(true)
    expect(resultValid.missingScopes).toBeUndefined()
  })

  test('roles and scopes', () => {
    user.role = 'user'
    user.scopes = ['scopes:write']

    const resultUserAdmin = isAuthorized({ user, requiredScopes: ['team-settings:write'], requiredRole: 'user' })
    const resultAdmin = isAuthorized({ user, requiredScopes: ['scopes:write'], requiredRole: 'admin' })
    const resultUserValid = isAuthorized({ user, requiredScopes: ['scopes:write'], requiredRole: 'user' })

    expect(resultUserAdmin.isAuthorized).toBe(false)
    expect(resultUserAdmin.missingRole).toBeUndefined()
    expect(resultUserAdmin.missingScopes).toEqual(['team-settings:write'])

    expect(resultAdmin.isAuthorized).toBe(false)
    expect(resultAdmin.missingRole).toBe('admin')
    expect(resultAdmin.missingScopes).toBeUndefined()

    expect(resultUserValid.isAuthorized).toBe(true)
    expect(resultUserValid.missingRole).toBeUndefined()
    expect(resultUserValid.missingScopes).toBeUndefined()
  })

  test('default authorized', () => {
    const result = isAuthorized({ user })

    expect(result.isAuthorized).toBe(true)
    expect(result.missingRole).toBeUndefined()
    expect(result.missingScopes).toBeUndefined()
  })

  test('no user unauthorized', () => {
    const result = isAuthorized({ user: undefined })

    expect(result.isAuthorized).toBe(false)
    expect(result.missingRole).toBeUndefined()
    expect(result.missingScopes).toBeUndefined()
  })
})

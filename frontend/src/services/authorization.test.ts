import type { UserResponse } from 'src/gen/model'
import { checkAuthorization } from 'src/services/authorization'
import { describe, expect, it, test } from 'vitest'

describe('roles and scopes', async () => {
  const user = {} as UserResponse

  test('role', () => {
    user.role = 'user'
    const resultAdmin = checkAuthorization({ user, requiredRole: 'admin' })
    const resultUser = checkAuthorization({ user, requiredRole: 'user' })

    expect(resultAdmin.authorized).toBe(false)
    expect(resultAdmin.missingRole).toBe('admin')

    expect(resultUser.authorized).toBe(true)
    expect(resultUser.missingRole).toBeUndefined()
  })

  test('scopes', () => {
    user.scopes = ['scopes:write']

    const resultWrite = checkAuthorization({ user, requiredScopes: ['team-settings:write'] })
    const resultBoth = checkAuthorization({ user, requiredScopes: ['team-settings:write', 'scopes:write'] })
    const resultValid = checkAuthorization({ user, requiredScopes: ['scopes:write'] })

    expect(resultWrite.authorized).toBe(false)
    expect(resultWrite.missingScopes).toEqual(['team-settings:write'])

    expect(resultBoth.authorized).toBe(false)
    expect(resultBoth.missingScopes).toEqual(['team-settings:write'])

    expect(resultValid.authorized).toBe(true)
    expect(resultValid.missingScopes).toBeUndefined()
  })

  test('roles and scopes', () => {
    user.role = 'user'
    user.scopes = ['scopes:write']

    const resultUserAdmin = checkAuthorization({ user, requiredScopes: ['team-settings:write'], requiredRole: 'user' })
    const resultAdmin = checkAuthorization({ user, requiredScopes: ['scopes:write'], requiredRole: 'admin' })
    const resultUserValid = checkAuthorization({ user, requiredScopes: ['scopes:write'], requiredRole: 'user' })

    expect(resultUserAdmin.authorized).toBe(false)
    expect(resultUserAdmin.missingRole).toBeUndefined()
    expect(resultUserAdmin.missingScopes).toEqual(['team-settings:write'])

    expect(resultAdmin.authorized).toBe(false)
    expect(resultAdmin.missingRole).toBe('admin')
    expect(resultAdmin.missingScopes).toBeUndefined()

    expect(resultUserValid.authorized).toBe(true)
    expect(resultUserValid.missingRole).toBeUndefined()
    expect(resultUserValid.missingScopes).toBeUndefined()
  })

  test('default authorized', () => {
    const result = checkAuthorization({ user })

    expect(result.authorized).toBe(true)
    expect(result.missingRole).toBeUndefined()
    expect(result.missingScopes).toBeUndefined()
  })

  test('no user unauthorized', () => {
    const result = checkAuthorization({ user: undefined })

    expect(result.authorized).toBe(false)
    expect(result.missingRole).toBeUndefined()
    expect(result.missingScopes).toBeUndefined()
  })
})

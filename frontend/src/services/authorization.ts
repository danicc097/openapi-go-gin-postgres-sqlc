import type { components } from '../types/schema'

type RolePermissions = {
  [key in components['schemas']['User']['role']]: components['schemas']['User']['role'][]
}

const ROLE_PERMISSIONS: RolePermissions = {
  user: ['user'],
  manager: ['manager', 'user'],
  admin: ['admin', 'manager', 'user'],
}
Object.freeze(ROLE_PERMISSIONS)

/**
 * Returns the roles allowed to view content for `role`.
 * @example getImplicitRoles('user') // ['user', 'manager', 'admin']
 */
const getImplicitRoles = (role: components['schemas']['User']['role']) => {
  return Object.keys(ROLE_PERMISSIONS).filter((r) => ROLE_PERMISSIONS[r].includes(role))
}

/**
 * Returns the access levels for a given role.
 * @example getAccessibleRoles('manager') // ['manager', 'user']
 */
const getRolePermissions = (role: components['schemas']['User']['role']) => {
  return ROLE_PERMISSIONS[role]
}

export { getImplicitRoles, getRolePermissions }

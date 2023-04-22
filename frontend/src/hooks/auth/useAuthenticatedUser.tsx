import { capitalize } from 'lodash'
import { useMemo } from 'react'
import { COLOR_BLIND_PALETTE, generateColor } from 'src/utils/colors'
import roles from '@roles'
import scopes from '@scopes'
import type { UserResponse } from 'src/gen/model'
import { getGetCurrentUserMock } from 'src/gen/user/user.msw'

export const useAuthenticatedUser = () => {
  // TODO for app_env dev, remove Authorization header and comes from backend via x-api-key header
  // or have fallthorugh if authentication failed instead - would need multierror
  // const user: UserResponse = {
  //   hasGlobalNotifications: true,
  //   hasPersonalNotifications: true,
  //   role: 'admin',
  //   userID: crypto.randomUUID(),
  //   email: 'admin@mail.com',
  //   firstName: 'John',
  //   lastName: 'Doe',
  //   fullName: 'John Doe',
  //   username: 'john.doe',
  //   scopes: ['users:read', 'test-scope', 'scopes:write'],
  //   createdAt: new Date(),
  //   deletedAt: null,
  // }

  const user: UserResponse = {
    role: 'user',
    scopes: ['users:read'],
    apiKey: null,
    teams: null,
    projects: null,
    user: {
      userID: 'c7fd2433-dbb7-4612-ab13-ddb0d3404728',
      username: 'user_2',
      email: 'user_2@email.com',
      firstName: 'Name 2',
      lastName: 'Surname 2',
      fullName: 'Name 2 Surname 2',
      hasPersonalNotifications: false,
      hasGlobalNotifications: true,
      createdAt: new Date('2023-04-01T06:24:22.390699Z'),
      deletedAt: null,
      timeEntries: null,
      userAPIKey: null,
      teams: null,
      workItems: null,
    },
  }

  const logUserOut = () => {
    null
  }

  return {
    user,
    logUserOut,
  }
}

import * as React from 'react'
import { BrowserRouter } from 'react-router-dom'

import CollapsibleNav from './CollapsibleNav'
import { test } from 'vitest'
// import { getGetCurrentUserMock } from 'src/gen/user/user.msw'
import type { UserResponse } from 'src/gen/model'

test('Renders content', async () => {
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
  return (
    <BrowserRouter>
      <CollapsibleNav user={user} />
    </BrowserRouter>
  )
})

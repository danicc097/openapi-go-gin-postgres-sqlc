import UserAvatar, { UserAvatarProps } from './UserAvatar'
import { BrowserRouter } from 'react-router-dom'
import React from 'react' // fix vitest
import { test } from 'vitest'
import { getGetCurrentUserMock } from 'src/gen/user/user.msw'
import type { UserResponse } from 'src/gen/model'

test('Renders content', async () => {
  const props: UserAvatarProps = {
    user: getGetCurrentUserMock(),
    size: 'l',
    initialsLength: 2,
    type: 'user',
  }
  return <UserAvatar {...props} />
})

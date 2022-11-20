import * as React from 'react'
import { screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import UserAvatar, { UserAvatarProps } from './UserAvatar'
import '@testing-library/jest-dom'
import { BrowserRouter } from 'react-router-dom'

import { render as renderWithStore } from 'src/test/test-utils'
import { testInitialState } from 'src/test/test-state'

test('Renders content', async () => {
  const props: UserAvatarProps = {
    user: testInitialState['auth']['user'],
    size: 'l',
    initialsLength: 2,
    type: 'user',
    color: '#eee',
  }
  renderWithStore(<UserAvatar {...props} />)
})

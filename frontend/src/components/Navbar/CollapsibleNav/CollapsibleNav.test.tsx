import * as React from 'react'
import { BrowserRouter } from 'react-router-dom'

import CollapsibleNav from './CollapsibleNav'
import { test } from 'vitest'
import { getGetCurrentUserMock } from 'src/gen/user/user.msw'

test('Renders content', async () => {
  const user = getGetCurrentUserMock()
  return (
    <BrowserRouter>
      <CollapsibleNav user={user} />
    </BrowserRouter>
  )
})

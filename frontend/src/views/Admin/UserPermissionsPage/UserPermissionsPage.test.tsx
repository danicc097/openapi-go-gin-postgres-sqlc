import * as React from 'react'
import { BrowserRouter } from 'react-router-dom'

import { test } from 'vitest'
import UserPermissionsPage from 'src/views/Admin/UserPermissionsPage/UserPermissionsPage'

test('Renders content', async () => {
  return new Promise((resolve) => {
    resolve(<UserPermissionsPage />)
  })
})

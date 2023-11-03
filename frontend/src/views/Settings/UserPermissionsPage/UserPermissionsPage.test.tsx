import * as React from 'react'
import { BrowserRouter } from 'react-router-dom'

import { test } from 'vitest'
import UserPermissionsPage from 'src/views/Settings/UserPermissionsPage/UserPermissionsPage'
import { render } from 'src/test-utils/render'

test('Renders content', async () => {
  return render(<UserPermissionsPage />)
})

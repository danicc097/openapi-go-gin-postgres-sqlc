import * as React from 'react'
import { BrowserRouter } from 'react-router-dom'

import { test } from 'vitest'
import UserPermissionsPage from 'src/views/Settings/UserPermissionsPage/UserPermissionsPage'
import { render } from 'src/test-utils/render'

test('Renders content', async () => {
  // FIXME:  import UserPermissionsPage breaks tests (error in src/TraceProvider)
  // Method Promise.prototype.then called on incompatible receiver [object Object]
  return render(<UserPermissionsPage />)
})

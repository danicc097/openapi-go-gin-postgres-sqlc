import * as React from 'react'
import { BrowserRouter } from 'react-router-dom'

import { test } from 'vitest'
import { setup } from 'src/test-utils/render'
// import ProjectManagementPage from 'src/views/Admin/ProjectManagementPage/ProjectManagementPage'

test('Renders content', async () => {
  return setup(<></>)
})

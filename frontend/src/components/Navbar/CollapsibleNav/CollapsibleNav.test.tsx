import * as React from 'react'
import { screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Navbar from '../Navbar'
import '@testing-library/jest-dom'
import { BrowserRouter } from 'react-router-dom'

import { render as renderWithStore } from 'src/test/test-utils'
import { testInitialState } from 'src/test/test-state'
import CollapsibleNav from './CollapsibleNav'

test('Renders content', async () => {
  renderWithStore(
    <BrowserRouter>
      <CollapsibleNav user={testInitialState.auth.user} />
    </BrowserRouter>,
    { initialState: testInitialState },
  )
})

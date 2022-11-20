import * as React from 'react'
import { screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Navbar from './Navbar'
import '@testing-library/jest-dom'
import { BrowserRouter } from 'react-router-dom'
import { vi } from 'vitest'
import { render as renderWithStore } from 'src/test/test-utils'
import { testInitialState } from 'src/test/test-state'

// import EventSourceSetup from 'src/setup'

const mEventSourceInstance = {
  addEventListener: vi.fn(),
}
const mEventSource: any = jest.fn(() => mEventSourceInstance)

global.EventSource = mEventSource

test('Renders content', async () => {
  renderWithStore(
    <BrowserRouter>
      <Navbar />
    </BrowserRouter>,
    { initialState: testInitialState },
  )
})

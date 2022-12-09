import Navbar from './Navbar'
import { BrowserRouter } from 'react-router-dom'
import { test, vi, vitest } from 'vitest'
import React from 'react' // fix vitest

// import EventSourceSetup from 'src/setup'

const mEventSourceInstance = {
  addEventListener: vi.fn(),
}
const mEventSource: any = vitest.fn(() => mEventSourceInstance)

global.EventSource = mEventSource

test('Renders content', async () => {
  return (
    <BrowserRouter>
      <Navbar />
    </BrowserRouter>
  )
})

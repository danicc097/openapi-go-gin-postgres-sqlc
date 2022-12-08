import * as React from 'react'
import LandingPage from './LandingPage'
import '@testing-library/jest-dom'
import { BrowserRouter } from 'react-router-dom'

import { test } from 'vitest'

test('Renders content', async () => {
  return <LandingPage />
})

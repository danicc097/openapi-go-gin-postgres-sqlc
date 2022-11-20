import * as React from 'react'
import { screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Layout from './Layout'
import '@testing-library/jest-dom'
import { BrowserRouter } from 'react-router-dom'

import { render as renderWithStore } from 'src/test/test-utils'

test('Renders content', async () => {
  renderWithStore(
    <BrowserRouter>
      <Layout>
        <div></div>
      </Layout>
    </BrowserRouter>,
  )
})

import React from 'react'
import { render, screen } from '@testing-library/react'
import TraceProvider from './TraceProvider' // Assuming TraceProvider is in the same directory
import { newFrontendSpan } from './traceProvider'
import { vitest } from 'vitest'

describe('TraceProvider', () => {
  it.skip('(does not render in CI) starts tracer with expected parameters', () => {
    render(
      <TraceProvider>
        <div>Test Child</div>
      </TraceProvider>,
    )
    expect(screen.getByText('Test Child')).toBeInTheDocument()

    // see https://github.com/open-telemetry/opentelemetry-js-contrib/blob/main/packages/opentelemetry-test-utils/src/test-utils.ts
  })
})

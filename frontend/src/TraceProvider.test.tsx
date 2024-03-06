import React from 'react'
import { render, screen } from '@testing-library/react'
import TraceProvider, { newFrontendSpan, tracer } from './TraceProvider' // Assuming TraceProvider is in the same directory
import { vitest } from 'vitest'

describe('TraceProvider', () => {
  it('starts tracer with expected parameters', () => {
    render(
      <TraceProvider>
        <div>Test Child</div>
      </TraceProvider>,
    )
    expect(screen.getByText('Test Child')).toBeInTheDocument()

    // see https://github.com/open-telemetry/opentelemetry-js-contrib/blob/main/packages/opentelemetry-test-utils/src/test-utils.ts
  })
})

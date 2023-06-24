import 'regenerator-runtime/runtime'
import { expect, afterEach, vi } from 'vitest'
import { cleanup } from '@testing-library/react'

import type { TestingLibraryMatchers } from '@testing-library/jest-dom/matchers'
import matchers from '@testing-library/jest-dom/matchers'

declare module 'vitest' {
  interface Assertion<T = any> extends jest.Matchers<void, T>, TestingLibraryMatchers<T, void> {}
}

expect.extend(matchers)

// runs a cleanup after each test case
afterEach(() => {
  cleanup() // clean jsdom
})

// eslint-disable-next-line @typescript-eslint/no-empty-function
window.URL.createObjectURL = (() => {}) as any

window.matchMedia = (query) => ({
  matches: false,
  media: query,
  onchange: null,
  addListener: vi.fn(), // deprecated
  removeListener: vi.fn(), // deprecated
  addEventListener: vi.fn(),
  removeEventListener: vi.fn(),
  dispatchEvent: vi.fn(),
})

export default class EventSourceSetup {
  eventSource

  constructor() {
    const eventSource = new EventSource('http://localhost')
    this.eventSource = eventSource

    eventSource.addEventListener('loading', function (event) {
      console.log('loading')
    })

    eventSource.addEventListener('loaded', function (event) {
      console.log('loaded')
    })

    eventSource.addEventListener('error', function (event) {
      console.log('error')
    })

    eventSource.onerror = (error) => {
      console.error('EventSource failed: ', error)
    }
  }
}

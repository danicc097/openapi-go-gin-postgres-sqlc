import 'regenerator-runtime/runtime'
import { expect, afterEach } from 'vitest'
import { cleanup } from '@testing-library/react'
import matchers from '@testing-library/jest-dom/matchers'

// extends Vitest's expect method with methods from react-testing-library
expect.extend(matchers)

// runs a cleanup after each test case (e.g. clearing jsdom)
afterEach(() => {
  cleanup()
})

// TODO find equivalent in vitest
// configure({ testIdAttribute: 'data-test-subj' })

// eslint-disable-next-line @typescript-eslint/no-empty-function
window.URL.createObjectURL = (() => {}) as any

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

import 'regenerator-runtime/runtime'
import { expect, afterEach, vi } from 'vitest'
import { cleanup } from '@testing-library/react'
import indexeddb from 'fake-indexeddb'

import type { TestingLibraryMatchers } from '@testing-library/jest-dom/matchers'
import matchers from '@testing-library/jest-dom/matchers'
import '@testing-library/jest-dom'
import '@testing-library/jest-dom/extend-expect'

import AxiosInterceptors from 'src/utils/axios'
import { AXIOS_INSTANCE } from 'src/api/mutator'

// its set in hook with a token, so we have to set explicitly in tests
AxiosInterceptors.setupAxiosInstance(AXIOS_INSTANCE, '')
// some calls in App.tsx should be abstracted as init.ts or the like
// so we just import that both in setupTests and App.tsx
import 'src/utils/dayjs'

// runs a cleanup after each test case
afterEach(() => {
  cleanup() // clean jsdom
})

// types
declare module 'vitest' {
  interface Assertion<T = any> extends jest.Matchers<void, T>, TestingLibraryMatchers<T, void> {}
}

expect.extend(matchers)

globalThis.indexedDB = indexeddb

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

export class PointerEvent extends Event {
  button: number
  ctrlKey: boolean

  constructor(type, props) {
    super(type, props)
    if (props.button != null) {
      this.button = props.button
    }
    if (props.ctrlKey != null) {
      this.ctrlKey = props.ctrlKey
    }
  }
}

window.PointerEvent = PointerEvent as any

import * as ResizeObserverModule from 'resize-observer-polyfill'
;(global as any).ResizeObserver = ResizeObserverModule.default
;(global as any).DOMRect = {
  fromRect: () => ({ top: 0, left: 0, bottom: 0, right: 0, width: 0, height: 0 }),
}

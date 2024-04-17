import { setupServer } from 'msw/node'

/* Vitest doesn't make vars in setupTests available, like jest does. */
export function setupMSW() {
  const server = setupServer()

  beforeAll(() => {
    // Start the interception.
    server.listen()
  })

  afterEach(() => {
    // Remove any handlers you may have added
    // in individual tests (runtime handlers).
    server.resetHandlers()
  })

  afterAll(() => {
    // Disable request interception and clean up.
    server.close()
  })

  return server
}

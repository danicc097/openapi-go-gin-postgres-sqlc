import { setupServer } from 'msw/node'

/* Vitest doesn't make vars in setupTests available, like jest does. */
export function setupMSW() {
  const server = setupServer()

  beforeAll(() => server.listen())
  afterAll(() => server.close())

  return server
}

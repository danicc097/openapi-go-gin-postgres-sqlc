// TODO: msw getGetPaginatedUsersMockHandler
import { setupServer } from 'msw/node'
import { getGetPaginatedUsersMockHandler } from 'src/gen/user/user.msw'

const server = setupServer()
server.use(getGetPaginatedUsersMockHandler())

test('', () => {
  null
})

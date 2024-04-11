// TODO: msw getGetPaginatedUsersMockHandler
import { setupServer } from 'msw/node'
import { getGetPaginatedUsersMockHandler } from 'src/gen/user/user.msw'
import { render } from 'src/test-utils'
import DemoMantineReactTable from 'src/views/DemoMantineReactTable/DemoMantineReactTable'

const server = setupServer()
server.use(getGetPaginatedUsersMockHandler())

test('Renders content', async () => {
  return render(<DemoMantineReactTable></DemoMantineReactTable>)
})

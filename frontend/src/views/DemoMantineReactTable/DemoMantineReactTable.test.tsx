// TODO: msw getGetPaginatedUsersMockHandler
import dayjs from 'dayjs'
import { HttpResponse, http } from 'msw'
import { setupServer } from 'msw/node'
import { UserID } from 'src/gen/entity-ids'
import { getGetPaginatedUsersMockHandler } from 'src/gen/user/user.msw'
import { render, screen } from 'src/test-utils'
import DemoMantineReactTable from 'src/views/DemoMantineReactTable/DemoMantineReactTable'

const server = setupServer()
server.use(
  getGetPaginatedUsersMockHandler({
    items: [
      {
        userID: '4af90297-125c-47db-a8ca-7e7d32d3a0b1' as UserID,
        username: 'user_49',
        email: 'user_49@mail.com',
        age: null,
        firstName: 'Reo',
        lastName: 'Michael',
        fullName: 'Reo Michael',
        scopes: ['users:read'],
        hasPersonalNotifications: true,
        hasGlobalNotifications: true,
        createdAt: dayjs('2024-04-07T11:10:48.123456+02:00').toDate(),
        updatedAt: dayjs('2024-04-08T11:10:48.123456+02:00').toDate(),
        deletedAt: null,
        role: 'user',
        teams: [],
        projects: [],
      },
    ],
    page: {
      nextCursor: 'next-cursor',
    },
  }),
)
// establish API mocking before all tests
beforeAll(() => server.listen())
// reset any request handlers that are declared as a part of our tests
// (i.e. for testing one-time error scenarios)
afterEach(() => server.resetHandlers())
// clean up once the tests are done
afterAll(() => server.close())

test('Renders content', async () => {
  render(<DemoMantineReactTable></DemoMantineReactTable>)

  const el = await screen.findByText('user_49@mail.com', {}, { timeout: 5000 })
  const allRows = screen.queryAllByRole('row')
  const firstRow = allRows.filter((row) => row.getAttribute('data-index') === '0')
})

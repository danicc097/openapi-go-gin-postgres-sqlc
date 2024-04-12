import { faker } from '@faker-js/faker'
import dayjs from 'dayjs'
import { HttpResponse, http } from 'msw'
import { setupServer } from 'msw/node'
import { UserID } from 'src/gen/entity-ids'
import { PaginatedUsersResponse, User } from 'src/gen/model'
import { getGetPaginatedUsersMockHandler } from 'src/gen/user/user.msw'
import { render, screen } from 'src/test-utils'
import DemoMantineReactTable from 'src/views/DemoMantineReactTable/DemoMantineReactTable'

function usersForPage(page: number): User[] {
  return [...Array(15)].map((x, i) => ({
    userID: faker.string.uuid() as UserID,
    username: `user_page${page}_${i}`,
    email: `user_page${page}_${i}@mail.com`,
    age: null,
    firstName: `${page}-${i}-A`,
    lastName: `${page}-${i}-B`,
    fullName: `${page}-${i}-A B`,
    scopes: ['users:read'],
    hasPersonalNotifications: true,
    hasGlobalNotifications: true,
    createdAt: dayjs('2024-04-07T11:10:48.123456+02:00').toDate(),
    updatedAt: dayjs('2024-04-08T11:10:48.123456+02:00').toDate(),
    deletedAt: null,
    role: 'user',
    teams: [],
    projects: [],
  }))
}

const server = setupServer()
const firstPage: PaginatedUsersResponse = {
  items: usersForPage(1),
  page: {
    nextCursor: 'next-cursor-1',
  },
}
const secondPage: PaginatedUsersResponse = {
  items: usersForPage(2),
  page: {
    nextCursor: 'next-cursor-2',
  },
}
// establish API mocking before all tests
beforeAll(() => server.listen())
// reset any request handlers that are declared as a part of our tests
// (i.e. for testing one-time error scenarios)
afterEach(() => server.resetHandlers())
// clean up once the tests are done
afterAll(() => server.close())

test('Renders content', async () => {
  render(<DemoMantineReactTable></DemoMantineReactTable>)

  server.use(getGetPaginatedUsersMockHandler(firstPage))

  console.log({ handlers: server.listHandlers() })
  console.log(document.body.innerHTML)

  const el = await screen.findByText(firstPage.items![0]!.email, {}, { timeout: 50000 })
  const allRows = screen.queryAllByRole('row')
  const firstRow = allRows.filter((row) => row.getAttribute('data-index') === '0')

  // TODO: scroll down container -Infinity
  // we should generate pages of length > 10 so that testing scroll on end reached works.

  // TODO: should test it was called with cursor=next-cursor-1
  server.use(getGetPaginatedUsersMockHandler(secondPage))
})

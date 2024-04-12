import { faker } from '@faker-js/faker'
import axios from 'axios'
import dayjs from 'dayjs'
import { HttpResponse, http } from 'msw'
import { setupServer } from 'msw/node'
import { UserID } from 'src/gen/entity-ids'
import { PaginatedUsersResponse, User } from 'src/gen/model'
import { getGetPaginatedUsersMockHandler } from 'src/gen/user/user.msw'
import { apiPath } from 'src/services/apiPaths'
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

console.log('debug 01')
const server = setupServer()
console.log('debug 02')
// establish API mocking before all tests
beforeAll(() => server.listen())
console.log('debug 03')
// reset any request handlers that are declared as a part of our tests
// (i.e. for testing one-time error scenarios)
console.log('debug 04')
afterEach(() => server.resetHandlers())
// clean up once the tests are done
console.log('debug 05')
afterAll(() => server.close())
console.log('debug 0')

test('mrt-table-tests-render', async () => {
  render(<DemoMantineReactTable></DemoMantineReactTable>)

  console.log('debug 1')
  server.use(getGetPaginatedUsersMockHandler(firstPage))

  console.log('debug 2')
  const p = apiPath()
  console.log(p)
  const res = await (await axios.get(`https://test.com/user/page`)).data
  console.log({ res })
  console.log({ handlers: server.listHandlers() })
  console.log(document.body.innerHTML)

  const el = await screen.findByText(firstPage.items![0]!.email, {}, { timeout: 5000 })
  const allRows = screen.queryAllByRole('row')
  const firstRow = allRows.filter((row) => row.getAttribute('data-index') === '0')

  // TODO: scroll down container -Infinity
  // we should generate pages of length > 10 so that testing scroll on end reached works.

  // TODO: should test it was called with cursor=next-cursor-1
  server.use(getGetPaginatedUsersMockHandler(secondPage))
})

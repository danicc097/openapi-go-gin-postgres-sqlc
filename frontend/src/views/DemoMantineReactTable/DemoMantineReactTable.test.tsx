import { faker } from '@faker-js/faker'
import axios from 'axios'
import dayjs from 'dayjs'
import { HttpResponse, http } from 'msw'
import { UserID } from 'src/gen/entity-ids'
import { PaginatedUsersResponse, User } from 'src/gen/model'
import { getGetPaginatedUsersMockHandler } from 'src/gen/user/user.msw'
import { apiPath } from 'src/services/apiPaths'
import { render, screen } from 'src/test-utils'
import { setupMSW } from 'src/test-utils/msw'
import DemoMantineReactTable from 'src/views/DemoMantineReactTable/DemoMantineReactTable'
import { vitest } from 'vitest'

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
    createdAt: dayjs().add(i, 'day').toDate(),
    updatedAt: dayjs().add(i, 'day').toDate(),
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
const server = setupMSW()

test('mrt-table-tests-render', async () => {
  const requestSpy = vitest.fn()
  server.events.on('request:start', requestSpy)

  server.use(getGetPaginatedUsersMockHandler(firstPage))

  render(<DemoMantineReactTable></DemoMantineReactTable>)

  const el = await screen.findByText(firstPage.items![0]!.email, {}, { timeout: 5000 })
  const url = new URL(requestSpy.mock.lastCall[0]['request']['url'])

  //FIXME: has actually fetched more (logging Fetching more...)
  expect(url.searchParams.get('cursor')).toBe('next-cursor-1')
  const allRows = screen.queryAllByRole('row')
  const firstRow = allRows.filter((row) => row.getAttribute('data-index') === '0')

  // TODO: scroll down container -Infinity
  // we should generate pages of length > 10 so that testing scroll on end reached works.

  // TODO: should test it was called with cursor=next-cursor-1
  server.use(getGetPaginatedUsersMockHandler(secondPage))
})

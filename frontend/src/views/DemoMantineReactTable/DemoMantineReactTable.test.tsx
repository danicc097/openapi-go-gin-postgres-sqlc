import { faker } from '@faker-js/faker'
import axios from 'axios'
import dayjs from 'dayjs'
import { HttpResponse, http } from 'msw'
import { UserID } from 'src/gen/entity-ids'
import { PaginatedUsersResponse, User } from 'src/gen/model'
import { getGetPaginatedUsersMockHandler } from 'src/gen/user/user.msw'
import { apiPath } from 'src/services/apiPaths'
import { fireEvent, render, screen } from 'src/test-utils'
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

  render(<DemoMantineReactTable></DemoMantineReactTable>)

  const table = document.getElementById('users-table')!
  vitest.spyOn(table, 'scrollHeight', 'get').mockImplementation(() => 1147)
  vitest.spyOn(table, 'clientHeight', 'get').mockImplementation(() => 585)
  // don't intercept until bounding rect mock is set up (let it retry network error - doesn't affect request spy calls)
  server.use(getGetPaginatedUsersMockHandler(firstPage))

  await screen.findByText(firstPage.items![0]!.email, {}, { timeout: 5000 })
  expect(requestSpy.mock.calls).toHaveLength(1)
  const firstPageUrl = new URL(requestSpy.mock.calls[0][0]['request']['url'])

  expect(firstPageUrl.searchParams.get('cursor')).toBe(null)
  const allRows = screen.queryAllByRole('row')
  const firstRow = allRows.filter((row) => row.getAttribute('data-index') === '0')
  // TODO: maybe should test rendering

  server.use(getGetPaginatedUsersMockHandler(secondPage))
  vitest.spyOn(table, 'scrollTop', 'get').mockImplementation(() => 500)
  fireEvent.scroll(table, { target: { scrollY: 100 } })
  await screen.findByText(secondPage.items![0]!.email, {}, { timeout: 5000 })

  const secondPageUrl = new URL(requestSpy.mock.calls[1][0]['request']['url'])
  expect(secondPageUrl.searchParams.get('cursor')).toBe('next-cursor-1')
})

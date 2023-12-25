/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { faker } from '@faker-js/faker'
import { HttpResponse, delay, http } from 'msw'

export const getCreateTeamMock = () => ({
  createdAt: (() => faker.date.past())(),
  description: faker.word.sample(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
  teamID: faker.number.int({ min: undefined, max: undefined }),
  updatedAt: (() => faker.date.past())(),
})

export const getGetTeamMock = () => ({
  createdAt: (() => faker.date.past())(),
  description: faker.word.sample(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
  teamID: faker.number.int({ min: undefined, max: undefined }),
  updatedAt: (() => faker.date.past())(),
})

export const getUpdateTeamMock = () => ({
  createdAt: (() => faker.date.past())(),
  description: faker.word.sample(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
  teamID: faker.number.int({ min: undefined, max: undefined }),
  updatedAt: (() => faker.date.past())(),
})

export const getTeamMock = () => [
  http.post('*/project/:projectName/team/', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getCreateTeamMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.get('*/project/:projectName/team/:id/', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getGetTeamMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.patch('*/project/:projectName/team/:id/', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getUpdateTeamMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.delete('*/project/:projectName/team/:id/', async () => {
    await delay(1000)
    return new HttpResponse(null, {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
]

/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { faker } from '@faker-js/faker'
import { HttpResponse, delay, http } from 'msw'

export const getCreateWorkItemTypeMock = () => ({
  color: faker.word.sample(),
  description: faker.word.sample(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
  workItemTypeID: faker.number.int({ min: undefined, max: undefined }),
})

export const getGetWorkItemTypeMock = () => ({
  color: faker.word.sample(),
  description: faker.word.sample(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
  workItemTypeID: faker.number.int({ min: undefined, max: undefined }),
})

export const getUpdateWorkItemTypeMock = () => ({
  color: faker.word.sample(),
  description: faker.word.sample(),
  name: faker.word.sample(),
  projectID: faker.number.int({ min: undefined, max: undefined }),
  workItemTypeID: faker.number.int({ min: undefined, max: undefined }),
})

export const getWorkItemTypeMock = () => [
  http.post('*/project/:projectName/workItemType/', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getCreateWorkItemTypeMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.get('*/workItemType/:id', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getGetWorkItemTypeMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.patch('*/workItemType/:id', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getUpdateWorkItemTypeMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.delete('*/workItemType/:id', async () => {
    await delay(1000)
    return new HttpResponse(null, {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
]

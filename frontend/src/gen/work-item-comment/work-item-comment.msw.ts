/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { faker } from '@faker-js/faker'
import { HttpResponse, delay, http } from 'msw'

export const getCreateWorkItemCommentMock = () => ({
  createdAt: (() => faker.date.past())(),
  message: faker.word.sample(),
  updatedAt: (() => faker.date.past())(),
  userID: (() => faker.datatype.uuid())(),
  workItemCommentID: faker.number.int({ min: undefined, max: undefined }),
  workItemID: faker.number.int({ min: undefined, max: undefined }),
})

export const getGetWorkItemCommentMock = () => ({
  createdAt: (() => faker.date.past())(),
  message: faker.word.sample(),
  updatedAt: (() => faker.date.past())(),
  userID: (() => faker.datatype.uuid())(),
  workItemCommentID: faker.number.int({ min: undefined, max: undefined }),
  workItemID: faker.number.int({ min: undefined, max: undefined }),
})

export const getUpdateWorkItemCommentMock = () => ({
  createdAt: (() => faker.date.past())(),
  message: faker.word.sample(),
  updatedAt: (() => faker.date.past())(),
  userID: (() => faker.datatype.uuid())(),
  workItemCommentID: faker.number.int({ min: undefined, max: undefined }),
  workItemID: faker.number.int({ min: undefined, max: undefined }),
})

export const getWorkItemCommentMock = () => [
  http.post('*/work-item/:workItemID/comment/', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getCreateWorkItemCommentMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.get('*/work-item/:workItemID/comment/:workItemCommentID', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getGetWorkItemCommentMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.patch('*/work-item/:workItemID/comment/:workItemCommentID', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getUpdateWorkItemCommentMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.delete('*/work-item/:workItemID/comment/:workItemCommentID', async () => {
    await delay(1000)
    return new HttpResponse(null, {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
]

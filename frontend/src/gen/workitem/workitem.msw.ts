/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { rest } from 'msw'
import { faker } from '@faker-js/faker'

export const getCreateWorkitemMock = () =>
  faker.helpers.arrayElement([
    {
      closed: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      createdAt: (() => faker.date.past())(),
      deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      description: faker.random.word(),
      kanbanStepID: faker.datatype.number({ min: undefined, max: undefined }),
      metadata: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() =>
        faker.datatype.number({ min: 0, max: undefined }),
      ),
      targetDate: (() => faker.date.past())(),
      teamID: faker.datatype.number({ min: undefined, max: undefined }),
      title: faker.random.word(),
      updatedAt: (() => faker.date.past())(),
      workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
    },
  ])

export const getGetWorkitemMock = () =>
  faker.helpers.arrayElement([
    {
      closed: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      createdAt: (() => faker.date.past())(),
      deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      description: faker.random.word(),
      kanbanStepID: faker.datatype.number({ min: undefined, max: undefined }),
      metadata: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() =>
        faker.datatype.number({ min: 0, max: undefined }),
      ),
      targetDate: (() => faker.date.past())(),
      teamID: faker.datatype.number({ min: undefined, max: undefined }),
      title: faker.random.word(),
      updatedAt: (() => faker.date.past())(),
      workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
    },
  ])

export const getUpdateWorkitemMock = () =>
  faker.helpers.arrayElement([
    {
      closed: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      createdAt: (() => faker.date.past())(),
      deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      description: faker.random.word(),
      kanbanStepID: faker.datatype.number({ min: undefined, max: undefined }),
      metadata: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() =>
        faker.datatype.number({ min: 0, max: undefined }),
      ),
      targetDate: (() => faker.date.past())(),
      teamID: faker.datatype.number({ min: undefined, max: undefined }),
      title: faker.random.word(),
      updatedAt: (() => faker.date.past())(),
      workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
    },
  ])

export const getCreateWorkitemCommentMock = () =>
  faker.helpers.arrayElement([
    {
      createdAt: (() => faker.date.past())(),
      message: faker.random.word(),
      updatedAt: (() => faker.date.past())(),
      userID: faker.random.word(),
      workItemCommentID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemID: faker.datatype.number({ min: undefined, max: undefined }),
    },
  ])

export const getWorkitemMSW = () => [
  rest.post('*/workitem/', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getCreateWorkitemMock()))
  }),
  rest.get('*/workitem/:id/', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getGetWorkitemMock()))
  }),
  rest.patch('*/workitem/:id/', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getUpdateWorkitemMock()))
  }),
  rest.delete('*/workitem/:id/', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'))
  }),
  rest.post('*/workitem/:id/comments/', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getCreateWorkitemCommentMock()))
  }),
]
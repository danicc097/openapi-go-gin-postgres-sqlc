/**
 * Generated by orval v6.19.1 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { faker } from '@faker-js/faker'
import { rest } from 'msw'

export const getPingMock = () => faker.random.word()

export const getOpenapiYamlGetMock = () => faker.word.sample()

export const getDefaultMSW = () => [
  rest.get('*/ping', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.text(getPingMock()))
  }),
  rest.get('*/openapi.yaml', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getOpenapiYamlGetMock()))
  }),
]

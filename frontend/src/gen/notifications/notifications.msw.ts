/**
 * Generated by orval v6.19.1 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { rest } from 'msw'

export const getNotificationsMSW = () => [
  rest.get('*/notifications/user/page', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'))
  }),
]

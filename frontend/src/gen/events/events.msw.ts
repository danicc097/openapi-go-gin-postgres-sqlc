/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { faker } from '@faker-js/faker'
import { HttpResponse, delay, http } from 'msw'

export const getEventsMock = () => faker.word.sample()

export const getEventsMock = () => [
  http.get('*/events', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getEventsMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
]

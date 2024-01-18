import type { Branded } from 'src/types/utils'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { faker } from '@faker-js/faker'
import { HttpResponse, delay, http } from 'msw'

export const getPingMock = () => faker.word.sample()

export const getOpenapiYamlGetMock = () => faker.word.sample()

export const getDefaultMock = () => [
  http.get('*/ping', async () => {
    await delay(1000)
    return new HttpResponse(getPingMock(), {
      status: 200,
      headers: {
        'Content-Type': 'text/plain',
      },
    })
  }),
  http.get('*/openapi.yaml', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getOpenapiYamlGetMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
]

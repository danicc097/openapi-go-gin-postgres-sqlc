import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import {
  faker
} from '@faker-js/faker'
import {
  HttpResponse,
  delay,
  http
} from 'msw'

export const getPingResponseMock = (): string => (faker.word.sample())

export const getOpenapiYamlGetResponseMock = (): Blob => (faker.word.sample())


export const getPingMockHandler = () => {
  return http.get('*/ping', async () => {
    await delay(200);
    return new HttpResponse(getPingResponseMock(),
      {
        status: 200,
        headers: {
          'Content-Type': 'text/plain',
        }
      }
    )
  })
}

export const getOpenapiYamlGetMockHandler = (overrideResponse?: Blob) => {
  return http.get('*/openapi.yaml', async () => {
    await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getOpenapiYamlGetResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}
export const getDefaultMock = () => [
  getPingMockHandler(),
  getOpenapiYamlGetMockHandler()
]

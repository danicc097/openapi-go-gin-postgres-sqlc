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
import type {
  Team
} from '.././model'

export const getCreateTeamResponseMock = (overrideResponse: any = {}): Team => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID, updatedAt: (() => faker.date.past())(), ...overrideResponse})

export const getGetTeamResponseMock = (overrideResponse: any = {}): Team => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID, updatedAt: (() => faker.date.past())(), ...overrideResponse})

export const getUpdateTeamResponseMock = (overrideResponse: any = {}): Team => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID, updatedAt: (() => faker.date.past())(), ...overrideResponse})


export const getCreateTeamMockHandler = (overrideResponse?: Team) => {
  return http.post('*/project/:projectName/team/', async () => {
    await delay(1000);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getCreateTeamResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetTeamMockHandler = (overrideResponse?: Team) => {
  return http.get('*/team/:teamID', async () => {
    await delay(1000);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getGetTeamResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateTeamMockHandler = (overrideResponse?: Team) => {
  return http.patch('*/team/:teamID', async () => {
    await delay(1000);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getUpdateTeamResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getDeleteTeamMockHandler = () => {
  return http.delete('*/team/:teamID', async () => {
    await delay(1000);
    return new HttpResponse(null,
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}
export const getTeamMock = () => [
  getCreateTeamMockHandler(),
  getGetTeamMockHandler(),
  getUpdateTeamMockHandler(),
  getDeleteTeamMockHandler()
]

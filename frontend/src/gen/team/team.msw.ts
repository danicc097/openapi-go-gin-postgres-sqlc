import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
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
  TeamResponse
} from '.././model'

export const getCreateTeamResponseMock = (overrideResponse: Partial< TeamResponse > = {}): TeamResponse => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID as EntityIDs.TeamID, updatedAt: (() => faker.date.past())(), ...overrideResponse})

export const getGetTeamResponseMock = (overrideResponse: Partial< TeamResponse > = {}): TeamResponse => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID as EntityIDs.TeamID, updatedAt: (() => faker.date.past())(), ...overrideResponse})

export const getUpdateTeamResponseMock = (overrideResponse: Partial< TeamResponse > = {}): TeamResponse => ({createdAt: (() => faker.date.past())(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID as EntityIDs.TeamID, updatedAt: (() => faker.date.past())(), ...overrideResponse})


export const getCreateTeamMockHandler = (overrideResponse?: TeamResponse | ((info: Parameters<Parameters<typeof http.post>[1]>[0]) => Promise<TeamResponse> | TeamResponse)) => {
  return http.post('*/project/:projectName/team/', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getCreateTeamResponseMock()),
      {
        status: 201,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetTeamMockHandler = (overrideResponse?: TeamResponse | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<TeamResponse> | TeamResponse)) => {
  return http.get('*/team/:teamID', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getGetTeamResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateTeamMockHandler = (overrideResponse?: TeamResponse | ((info: Parameters<Parameters<typeof http.patch>[1]>[0]) => Promise<TeamResponse> | TeamResponse)) => {
  return http.patch('*/team/:teamID', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getUpdateTeamResponseMock()),
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
  return http.delete('*/team/:teamID', async () => {await delay(200);
    return new HttpResponse(null,
      {
        status: 204,
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

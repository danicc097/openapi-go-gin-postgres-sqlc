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
  ActivityResponse
} from '.././model'

export const getCreateActivityResponseMock = (overrideResponse: any = {}): ActivityResponse => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID, deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), isProductive: faker.datatype.boolean(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, ...overrideResponse})

export const getGetActivityResponseMock = (overrideResponse: any = {}): ActivityResponse => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID, deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), isProductive: faker.datatype.boolean(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, ...overrideResponse})

export const getUpdateActivityResponseMock = (overrideResponse: any = {}): ActivityResponse => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID, deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), isProductive: faker.datatype.boolean(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, ...overrideResponse})


export const getCreateActivityMockHandler = (overrideResponse?: ActivityResponse) => {
  return http.post('*/project/:projectName/activity/', async () => {
    await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getCreateActivityResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetActivityMockHandler = (overrideResponse?: ActivityResponse) => {
  return http.get('*/activity/:activityID', async () => {
    await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getGetActivityResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateActivityMockHandler = (overrideResponse?: ActivityResponse) => {
  return http.patch('*/activity/:activityID', async () => {
    await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse ? overrideResponse : getUpdateActivityResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getDeleteActivityMockHandler = () => {
  return http.delete('*/activity/:activityID', async () => {
    await delay(200);
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
export const getActivityMock = () => [
  getCreateActivityMockHandler(),
  getGetActivityMockHandler(),
  getUpdateActivityMockHandler(),
  getDeleteActivityMockHandler()
]

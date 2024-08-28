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
  WorkItemCommentResponse
} from '.././model'

export const getCreateWorkItemCommentResponseMock = (overrideResponse: Partial< WorkItemCommentResponse > = {}): WorkItemCommentResponse => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.helpers.fromRegExp('^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}'), workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID, ...overrideResponse})

export const getGetWorkItemCommentResponseMock = (overrideResponse: Partial< WorkItemCommentResponse > = {}): WorkItemCommentResponse => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.helpers.fromRegExp('^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}'), workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID, ...overrideResponse})

export const getUpdateWorkItemCommentResponseMock = (overrideResponse: Partial< WorkItemCommentResponse > = {}): WorkItemCommentResponse => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.helpers.fromRegExp('^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}'), workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID, ...overrideResponse})


export const getCreateWorkItemCommentMockHandler = (overrideResponse?: WorkItemCommentResponse | ((info: Parameters<Parameters<typeof http.post>[1]>[0]) => Promise<WorkItemCommentResponse> | WorkItemCommentResponse)) => {
  return http.post('*/work-item/:workItemID/comment/', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getCreateWorkItemCommentResponseMock()),
      {
        status: 201,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetWorkItemCommentMockHandler = (overrideResponse?: WorkItemCommentResponse | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<WorkItemCommentResponse> | WorkItemCommentResponse)) => {
  return http.get('*/work-item/:workItemID/comment/:workItemCommentID', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getGetWorkItemCommentResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateWorkItemCommentMockHandler = (overrideResponse?: WorkItemCommentResponse | ((info: Parameters<Parameters<typeof http.patch>[1]>[0]) => Promise<WorkItemCommentResponse> | WorkItemCommentResponse)) => {
  return http.patch('*/work-item/:workItemID/comment/:workItemCommentID', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getUpdateWorkItemCommentResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getDeleteWorkItemCommentMockHandler = () => {
  return http.delete('*/work-item/:workItemID/comment/:workItemCommentID', async () => {await delay(200);
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
export const getWorkItemCommentMock = () => [
  getCreateWorkItemCommentMockHandler(),
  getGetWorkItemCommentMockHandler(),
  getUpdateWorkItemCommentMockHandler(),
  getDeleteWorkItemCommentMockHandler()
]

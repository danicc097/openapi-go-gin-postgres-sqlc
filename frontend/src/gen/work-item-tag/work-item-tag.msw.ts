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
  WorkItemTagResponse
} from '.././model'

export const getCreateWorkItemTagResponseMock = (overrideResponse: Partial< WorkItemTagResponse > = {}): WorkItemTagResponse => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID, ...overrideResponse})

export const getGetWorkItemTagResponseMock = (overrideResponse: Partial< WorkItemTagResponse > = {}): WorkItemTagResponse => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID, ...overrideResponse})

export const getUpdateWorkItemTagResponseMock = (overrideResponse: Partial< WorkItemTagResponse > = {}): WorkItemTagResponse => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID, ...overrideResponse})


export const getCreateWorkItemTagMockHandler = (overrideResponse?: WorkItemTagResponse | ((info: Parameters<Parameters<typeof http.post>[1]>[0]) => Promise<WorkItemTagResponse> | WorkItemTagResponse)) => {
  return http.post('*/project/:projectName/work-item-tag/', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getCreateWorkItemTagResponseMock()),
      {
        status: 201,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetWorkItemTagMockHandler = (overrideResponse?: WorkItemTagResponse | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<WorkItemTagResponse> | WorkItemTagResponse)) => {
  return http.get('*/work-item-tag/:workItemTagID', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getGetWorkItemTagResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateWorkItemTagMockHandler = (overrideResponse?: WorkItemTagResponse | ((info: Parameters<Parameters<typeof http.patch>[1]>[0]) => Promise<WorkItemTagResponse> | WorkItemTagResponse)) => {
  return http.patch('*/work-item-tag/:workItemTagID', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getUpdateWorkItemTagResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getDeleteWorkItemTagMockHandler = () => {
  return http.delete('*/work-item-tag/:workItemTagID', async () => {await delay(200);
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
export const getWorkItemTagMock = () => [
  getCreateWorkItemTagMockHandler(),
  getGetWorkItemTagMockHandler(),
  getUpdateWorkItemTagMockHandler(),
  getDeleteWorkItemTagMockHandler()
]

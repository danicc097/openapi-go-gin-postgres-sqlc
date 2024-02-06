import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
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

export const getCreateWorkItemTypeMock = () => ({color: faker.word.sample(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID})

export const getGetWorkItemTypeMock = () => ({color: faker.word.sample(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID})

export const getUpdateWorkItemTypeMock = () => ({color: faker.word.sample(), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID})

export const getWorkItemTypeMock = () => [
http.post('*/project/:projectName/work-item-type/', async () => {
        await delay(1000);
        return new HttpResponse(JSON.stringify(getCreateWorkItemTypeMock()),
          { 
            status: 200,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        )
      }),http.get('*/work-item-type/:workItemTypeID', async () => {
        await delay(1000);
        return new HttpResponse(JSON.stringify(getGetWorkItemTypeMock()),
          { 
            status: 200,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        )
      }),http.patch('*/work-item-type/:workItemTypeID', async () => {
        await delay(1000);
        return new HttpResponse(JSON.stringify(getUpdateWorkItemTypeMock()),
          { 
            status: 200,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        )
      }),http.delete('*/work-item-type/:workItemTypeID', async () => {
        await delay(1000);
        return new HttpResponse(null,
          { 
            status: 200,
            headers: {
              'Content-Type': 'application/json',
            }
          }
        )
      }),]

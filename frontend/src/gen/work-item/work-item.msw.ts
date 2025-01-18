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
import {
  Scope,
  WorkItemRole
} from '.././model'
import type {
  DemoTwoWorkItemResponse,
  DemoWorkItemResponse,
  PaginatedDemoWorkItemsResponse,
  WorkItemResponse
} from '.././model'

export const getCreateWorkitemResponseDemoWorkItemResponseMock = (overrideResponse: Partial<DemoWorkItemResponse> = {}): DemoWorkItemResponse => ({...{closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), demoWorkItem: {lastMessageAt: (() => faker.date.past())(), line: faker.word.sample(), ref: faker.helpers.fromRegExp('^[0-9]{8}$'), reopened: faker.datatype.boolean(), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID}, description: faker.word.sample(), kanbanStepID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.KanbanStepID as EntityIDs.KanbanStepID, members: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({role: faker.helpers.arrayElement(Object.values(WorkItemRole)), user: {age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: faker.word.sample(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), scopes: faker.helpers.arrayElements(Object.values(Scope)), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample()}})), metadata: (() => ({
              key: faker.string.sample()
            }))(), projectName: faker.helpers.arrayElement(['demo'] as const), targetDate: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntries: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID as EntityIDs.ActivityID, comment: faker.word.sample(), durationMinutes: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), start: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntryID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TimeEntryID as EntityIDs.TimeEntryID, userID: faker.string.uuid() as EntityIDs.UserID, workItemID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.WorkItemID | null})), title: faker.word.sample(), updatedAt: (() => faker.date.past())(), workItemComments: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID})), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID, workItemTags: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID as EntityIDs.WorkItemTagID})), workItemType: {color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, ...overrideResponse});

export const getCreateWorkitemResponseDemoTwoWorkItemResponseMock = (overrideResponse: Partial<DemoTwoWorkItemResponse> = {}): DemoTwoWorkItemResponse => ({...{closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), demoTwoWorkItem: {customDateForProject2: faker.helpers.arrayElement([(() => faker.date.past())(), null]), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID}, description: faker.word.sample(), kanbanStepID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.KanbanStepID as EntityIDs.KanbanStepID, members: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({role: faker.helpers.arrayElement(Object.values(WorkItemRole)), user: {age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: faker.word.sample(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), scopes: faker.helpers.arrayElements(Object.values(Scope)), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample()}})), metadata: (() => ({
              key: faker.string.sample()
            }))(), projectName: faker.helpers.arrayElement(['demo_two'] as const), targetDate: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntries: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID as EntityIDs.ActivityID, comment: faker.word.sample(), durationMinutes: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), start: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntryID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TimeEntryID as EntityIDs.TimeEntryID, userID: faker.string.uuid() as EntityIDs.UserID, workItemID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.WorkItemID | null})), title: faker.word.sample(), updatedAt: (() => faker.date.past())(), workItemComments: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID})), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID, workItemTags: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID as EntityIDs.WorkItemTagID})), workItemType: {color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, ...overrideResponse});

export const getCreateWorkitemResponseMock = (): WorkItemResponse => (faker.helpers.arrayElement([{...getCreateWorkitemResponseDemoWorkItemResponseMock()},{...getCreateWorkitemResponseDemoTwoWorkItemResponseMock()}]))

export const getGetWorkItemResponseDemoWorkItemResponseMock = (overrideResponse: Partial<DemoWorkItemResponse> = {}): DemoWorkItemResponse => ({...{closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), demoWorkItem: {lastMessageAt: (() => faker.date.past())(), line: faker.word.sample(), ref: faker.helpers.fromRegExp('^[0-9]{8}$'), reopened: faker.datatype.boolean(), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID}, description: faker.word.sample(), kanbanStepID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.KanbanStepID as EntityIDs.KanbanStepID, members: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({role: faker.helpers.arrayElement(Object.values(WorkItemRole)), user: {age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: faker.word.sample(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), scopes: faker.helpers.arrayElements(Object.values(Scope)), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample()}})), metadata: (() => ({
              key: faker.string.sample()
            }))(), projectName: faker.helpers.arrayElement(['demo'] as const), targetDate: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntries: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID as EntityIDs.ActivityID, comment: faker.word.sample(), durationMinutes: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), start: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntryID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TimeEntryID as EntityIDs.TimeEntryID, userID: faker.string.uuid() as EntityIDs.UserID, workItemID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.WorkItemID | null})), title: faker.word.sample(), updatedAt: (() => faker.date.past())(), workItemComments: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID})), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID, workItemTags: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID as EntityIDs.WorkItemTagID})), workItemType: {color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, ...overrideResponse});

export const getGetWorkItemResponseDemoTwoWorkItemResponseMock = (overrideResponse: Partial<DemoTwoWorkItemResponse> = {}): DemoTwoWorkItemResponse => ({...{closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), demoTwoWorkItem: {customDateForProject2: faker.helpers.arrayElement([(() => faker.date.past())(), null]), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID}, description: faker.word.sample(), kanbanStepID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.KanbanStepID as EntityIDs.KanbanStepID, members: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({role: faker.helpers.arrayElement(Object.values(WorkItemRole)), user: {age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: faker.word.sample(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), scopes: faker.helpers.arrayElements(Object.values(Scope)), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample()}})), metadata: (() => ({
              key: faker.string.sample()
            }))(), projectName: faker.helpers.arrayElement(['demo_two'] as const), targetDate: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntries: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID as EntityIDs.ActivityID, comment: faker.word.sample(), durationMinutes: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), start: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntryID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TimeEntryID as EntityIDs.TimeEntryID, userID: faker.string.uuid() as EntityIDs.UserID, workItemID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.WorkItemID | null})), title: faker.word.sample(), updatedAt: (() => faker.date.past())(), workItemComments: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID})), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID, workItemTags: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID as EntityIDs.WorkItemTagID})), workItemType: {color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, ...overrideResponse});

export const getGetWorkItemResponseMock = (): WorkItemResponse => (faker.helpers.arrayElement([{...getGetWorkItemResponseDemoWorkItemResponseMock()},{...getGetWorkItemResponseDemoTwoWorkItemResponseMock()}]))

export const getUpdateWorkitemResponseDemoWorkItemResponseMock = (overrideResponse: Partial<DemoWorkItemResponse> = {}): DemoWorkItemResponse => ({...{closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), demoWorkItem: {lastMessageAt: (() => faker.date.past())(), line: faker.word.sample(), ref: faker.helpers.fromRegExp('^[0-9]{8}$'), reopened: faker.datatype.boolean(), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID}, description: faker.word.sample(), kanbanStepID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.KanbanStepID as EntityIDs.KanbanStepID, members: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({role: faker.helpers.arrayElement(Object.values(WorkItemRole)), user: {age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: faker.word.sample(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), scopes: faker.helpers.arrayElements(Object.values(Scope)), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample()}})), metadata: (() => ({
              key: faker.string.sample()
            }))(), projectName: faker.helpers.arrayElement(['demo'] as const), targetDate: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntries: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID as EntityIDs.ActivityID, comment: faker.word.sample(), durationMinutes: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), start: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntryID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TimeEntryID as EntityIDs.TimeEntryID, userID: faker.string.uuid() as EntityIDs.UserID, workItemID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.WorkItemID | null})), title: faker.word.sample(), updatedAt: (() => faker.date.past())(), workItemComments: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID})), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID, workItemTags: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID as EntityIDs.WorkItemTagID})), workItemType: {color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, ...overrideResponse});

export const getUpdateWorkitemResponseDemoTwoWorkItemResponseMock = (overrideResponse: Partial<DemoTwoWorkItemResponse> = {}): DemoTwoWorkItemResponse => ({...{closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), demoTwoWorkItem: {customDateForProject2: faker.helpers.arrayElement([(() => faker.date.past())(), null]), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID}, description: faker.word.sample(), kanbanStepID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.KanbanStepID as EntityIDs.KanbanStepID, members: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({role: faker.helpers.arrayElement(Object.values(WorkItemRole)), user: {age: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), email: faker.word.sample(), firstName: faker.helpers.arrayElement([faker.word.sample(), null]), fullName: faker.helpers.arrayElement([faker.word.sample(), null]), hasGlobalNotifications: faker.datatype.boolean(), hasPersonalNotifications: faker.datatype.boolean(), lastName: faker.helpers.arrayElement([faker.word.sample(), null]), scopes: faker.helpers.arrayElements(Object.values(Scope)), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, username: faker.word.sample()}})), metadata: (() => ({
              key: faker.string.sample()
            }))(), projectName: faker.helpers.arrayElement(['demo_two'] as const), targetDate: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntries: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({activityID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ActivityID as EntityIDs.ActivityID, comment: faker.word.sample(), durationMinutes: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]), start: (() => faker.date.past())(), teamID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.TeamID | null, timeEntryID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TimeEntryID as EntityIDs.TimeEntryID, userID: faker.string.uuid() as EntityIDs.UserID, workItemID: faker.helpers.arrayElement([faker.number.int({min: undefined, max: undefined}), null]) as EntityIDs.WorkItemID | null})), title: faker.word.sample(), updatedAt: (() => faker.date.past())(), workItemComments: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({createdAt: (() => faker.date.past())(), message: faker.word.sample(), updatedAt: (() => faker.date.past())(), userID: faker.string.uuid() as EntityIDs.UserID, workItemCommentID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemCommentID as EntityIDs.WorkItemCommentID, workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID})), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID, workItemTags: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTagID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTagID as EntityIDs.WorkItemTagID})), workItemType: {color: faker.helpers.fromRegExp('^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$'), description: faker.word.sample(), name: faker.word.sample(), projectID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.ProjectID as EntityIDs.ProjectID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID}, ...overrideResponse});

export const getUpdateWorkitemResponseMock = (): WorkItemResponse => (faker.helpers.arrayElement([{...getUpdateWorkitemResponseDemoWorkItemResponseMock()},{...getUpdateWorkitemResponseDemoTwoWorkItemResponseMock()}]))

export const getGetPaginatedWorkItemResponseMock = (overrideResponse: Partial< PaginatedDemoWorkItemsResponse > = {}): PaginatedDemoWorkItemsResponse => ({items: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), createdAt: (() => faker.date.past())(), deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]), description: faker.word.sample(), kanbanStepID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.KanbanStepID as EntityIDs.KanbanStepID, lastMessageAt: (() => faker.date.past())(), line: faker.word.sample(), metadata: {
        [faker.string.alphanumeric(5)]: {}
      }, ref: faker.helpers.fromRegExp('^[0-9]{8}$'), reopened: faker.datatype.boolean(), targetDate: (() => faker.date.past())(), teamID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.TeamID as EntityIDs.TeamID, title: faker.word.sample(), updatedAt: (() => faker.date.past())(), workItemID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemID as EntityIDs.WorkItemID, workItemTypeID: faker.number.int({min: undefined, max: undefined}) as EntityIDs.WorkItemTypeID as EntityIDs.WorkItemTypeID})), page: {nextCursor: faker.word.sample()}, ...overrideResponse})


export const getCreateWorkitemMockHandler = (overrideResponse?: WorkItemResponse | ((info: Parameters<Parameters<typeof http.post>[1]>[0]) => Promise<WorkItemResponse> | WorkItemResponse)) => {
  return http.post('*/work-item/', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getCreateWorkitemResponseMock()),
      {
        status: 201,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getGetWorkItemMockHandler = (overrideResponse?: WorkItemResponse | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<WorkItemResponse> | WorkItemResponse)) => {
  return http.get('*/work-item/:workItemID/', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getGetWorkItemResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getUpdateWorkitemMockHandler = (overrideResponse?: WorkItemResponse | ((info: Parameters<Parameters<typeof http.patch>[1]>[0]) => Promise<WorkItemResponse> | WorkItemResponse)) => {
  return http.patch('*/work-item/:workItemID/', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getUpdateWorkitemResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}

export const getDeleteWorkitemMockHandler = () => {
  return http.delete('*/work-item/:workItemID/', async () => {await delay(200);
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

export const getGetPaginatedWorkItemMockHandler = (overrideResponse?: PaginatedDemoWorkItemsResponse | ((info: Parameters<Parameters<typeof http.get>[1]>[0]) => Promise<PaginatedDemoWorkItemsResponse> | PaginatedDemoWorkItemsResponse)) => {
  return http.get('*/work-item/page', async (info) => {await delay(200);
    return new HttpResponse(JSON.stringify(overrideResponse !== undefined 
            ? (typeof overrideResponse === "function" ? await overrideResponse(info) : overrideResponse) 
            : getGetPaginatedWorkItemResponseMock()),
      {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        }
      }
    )
  })
}
export const getWorkItemMock = () => [
  getCreateWorkitemMockHandler(),
  getGetWorkItemMockHandler(),
  getUpdateWorkitemMockHandler(),
  getDeleteWorkitemMockHandler(),
  getGetPaginatedWorkItemMockHandler()
]

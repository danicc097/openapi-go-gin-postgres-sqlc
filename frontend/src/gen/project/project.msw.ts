/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { faker } from '@faker-js/faker'
import { HttpResponse, delay, http } from 'msw'
import { Project, Scope, WorkItemRole } from '.././model'

export const getGetProjectMock = () => ({
  boardConfig: {
    fields: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
      isEditable: faker.datatype.boolean(),
      isVisible: faker.datatype.boolean(),
      name: faker.word.sample(),
      path: faker.word.sample(),
      showCollapsed: faker.datatype.boolean(),
    })),
    header: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() =>
      faker.word.sample(),
    ),
    visualization: {},
  },
  createdAt: (() => faker.date.past())(),
  description: faker.word.sample(),
  name: faker.helpers.arrayElement(Object.values(Project)),
  projectID: faker.number.int({ min: undefined, max: undefined }),
  updatedAt: (() => faker.date.past())(),
})

export const getGetProjectConfigMock = () => ({
  fields: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
    isEditable: faker.datatype.boolean(),
    isVisible: faker.datatype.boolean(),
    name: faker.word.sample(),
    path: faker.word.sample(),
    showCollapsed: faker.datatype.boolean(),
  })),
  header: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => faker.word.sample()),
  visualization: {},
})

export const getGetProjectBoardMock = () => ({ projectName: faker.helpers.arrayElement(Object.values(Project)) })

export const getGetProjectWorkitemsMock = () =>
  faker.helpers.arrayElement([
    {
      closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      createdAt: (() => faker.date.past())(),
      deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      demoWorkItem: {
        lastMessageAt: (() => faker.date.past())(),
        line: faker.word.sample(),
        ref: faker.word.sample(),
        reopened: faker.datatype.boolean(),
        workItemID: faker.number.int({ min: undefined, max: undefined }),
      },
      description: faker.word.sample(),
      kanbanStepID: faker.number.int({ min: undefined, max: undefined }),
      members: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        role: faker.helpers.arrayElement(Object.values(WorkItemRole)),
        user: {
          createdAt: (() => faker.date.past())(),
          deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
          email: faker.word.sample(),
          firstName: faker.helpers.arrayElement([faker.word.sample(), null]),
          fullName: faker.helpers.arrayElement([faker.word.sample(), null]),
          hasGlobalNotifications: faker.datatype.boolean(),
          hasPersonalNotifications: faker.datatype.boolean(),
          lastName: faker.helpers.arrayElement([faker.word.sample(), null]),
          scopes: faker.helpers.arrayElements(Object.values(Scope)),
          userID: faker.word.sample(),
          username: faker.word.sample(),
        },
      })),
      metadata: (() => ({
        key: faker.color.hsl(),
      }))(),
      targetDate: (() => faker.date.past())(),
      teamID: faker.helpers.arrayElement([faker.number.int({ min: undefined, max: undefined }), null]),
      timeEntries: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        activityID: faker.number.int({ min: undefined, max: undefined }),
        comment: faker.word.sample(),
        durationMinutes: faker.helpers.arrayElement([faker.number.int({ min: undefined, max: undefined }), null]),
        start: (() => faker.date.past())(),
        teamID: faker.helpers.arrayElement([faker.number.int({ min: undefined, max: undefined }), null]),
        timeEntryID: faker.number.int({ min: undefined, max: undefined }),
        userID: faker.word.sample(),
        workItemID: faker.helpers.arrayElement([faker.number.int({ min: undefined, max: undefined }), null]),
      })),
      title: faker.word.sample(),
      updatedAt: (() => faker.date.past())(),
      workItemComments: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        createdAt: (() => faker.date.past())(),
        message: faker.word.sample(),
        updatedAt: (() => faker.date.past())(),
        userID: faker.word.sample(),
        workItemCommentID: faker.number.int({ min: undefined, max: undefined }),
        workItemID: faker.number.int({ min: undefined, max: undefined }),
      })),
      workItemID: faker.number.int({ min: undefined, max: undefined }),
      workItemTags: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        color: faker.word.sample(),
        deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
        description: faker.word.sample(),
        name: faker.word.sample(),
        projectID: faker.number.int({ min: undefined, max: undefined }),
        workItemTagID: faker.number.int({ min: undefined, max: undefined }),
      })),
      workItemType: {
        color: faker.word.sample(),
        description: faker.word.sample(),
        name: faker.word.sample(),
        projectID: faker.number.int({ min: undefined, max: undefined }),
        workItemTypeID: faker.number.int({ min: undefined, max: undefined }),
      },
      workItemTypeID: faker.number.int({ min: undefined, max: undefined }),
    },
    {
      closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      createdAt: (() => faker.date.past())(),
      deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      demoTwoWorkItem: {
        customDateForProject2: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
        workItemID: faker.number.int({ min: undefined, max: undefined }),
      },
      description: faker.word.sample(),
      kanbanStepID: faker.number.int({ min: undefined, max: undefined }),
      members: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        role: faker.helpers.arrayElement(Object.values(WorkItemRole)),
        user: {
          createdAt: (() => faker.date.past())(),
          deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
          email: faker.word.sample(),
          firstName: faker.helpers.arrayElement([faker.word.sample(), null]),
          fullName: faker.helpers.arrayElement([faker.word.sample(), null]),
          hasGlobalNotifications: faker.datatype.boolean(),
          hasPersonalNotifications: faker.datatype.boolean(),
          lastName: faker.helpers.arrayElement([faker.word.sample(), null]),
          scopes: faker.helpers.arrayElements(Object.values(Scope)),
          userID: faker.word.sample(),
          username: faker.word.sample(),
        },
      })),
      metadata: (() => ({
        key: faker.color.hsl(),
      }))(),
      targetDate: (() => faker.date.past())(),
      teamID: faker.helpers.arrayElement([faker.number.int({ min: undefined, max: undefined }), null]),
      timeEntries: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        activityID: faker.number.int({ min: undefined, max: undefined }),
        comment: faker.word.sample(),
        durationMinutes: faker.helpers.arrayElement([faker.number.int({ min: undefined, max: undefined }), null]),
        start: (() => faker.date.past())(),
        teamID: faker.helpers.arrayElement([faker.number.int({ min: undefined, max: undefined }), null]),
        timeEntryID: faker.number.int({ min: undefined, max: undefined }),
        userID: faker.word.sample(),
        workItemID: faker.helpers.arrayElement([faker.number.int({ min: undefined, max: undefined }), null]),
      })),
      title: faker.word.sample(),
      updatedAt: (() => faker.date.past())(),
      workItemComments: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        createdAt: (() => faker.date.past())(),
        message: faker.word.sample(),
        updatedAt: (() => faker.date.past())(),
        userID: faker.word.sample(),
        workItemCommentID: faker.number.int({ min: undefined, max: undefined }),
        workItemID: faker.number.int({ min: undefined, max: undefined }),
      })),
      workItemID: faker.number.int({ min: undefined, max: undefined }),
      workItemTags: Array.from({ length: faker.number.int({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        color: faker.word.sample(),
        deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
        description: faker.word.sample(),
        name: faker.word.sample(),
        projectID: faker.number.int({ min: undefined, max: undefined }),
        workItemTagID: faker.number.int({ min: undefined, max: undefined }),
      })),
      workItemType: {
        color: faker.word.sample(),
        description: faker.word.sample(),
        name: faker.word.sample(),
        projectID: faker.number.int({ min: undefined, max: undefined }),
        workItemTypeID: faker.number.int({ min: undefined, max: undefined }),
      },
      workItemTypeID: faker.number.int({ min: undefined, max: undefined }),
    },
  ])

export const getProjectMock = () => [
  http.post('*/project/:projectName/initialize', async () => {
    await delay(1000)
    return new HttpResponse(null, {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.get('*/project/:projectName/', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getGetProjectMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.get('*/project/:projectName/config', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getGetProjectConfigMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.put('*/project/:projectName/config', async () => {
    await delay(1000)
    return new HttpResponse(null, {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.get('*/project/:projectName/board', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getGetProjectBoardMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
  http.get('*/project/:projectName/workitems', async () => {
    await delay(1000)
    return new HttpResponse(JSON.stringify(getGetProjectWorkitemsMock()), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    })
  }),
]

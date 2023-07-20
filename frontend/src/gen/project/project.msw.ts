/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { rest } from 'msw'
import { faker } from '@faker-js/faker'
import { Project, Scope } from '.././model'

export const getGetProjectMock = () => ({
  boardConfig: {
    fields: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
      isEditable: faker.datatype.boolean(),
      isVisible: faker.datatype.boolean(),
      name: faker.random.word(),
      path: faker.random.word(),
      showCollapsed: faker.datatype.boolean(),
    })),
    header: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() =>
      faker.random.word(),
    ),
    visualization: {},
  },
  createdAt: (() => faker.date.past())(),
  description: faker.random.word(),
  name: faker.helpers.arrayElement(Object.values(Project)),
  projectID: faker.datatype.number({ min: undefined, max: undefined }),
  updatedAt: (() => faker.date.past())(),
})

export const getGetProjectConfigMock = () => ({
  fields: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
    isEditable: faker.datatype.boolean(),
    isVisible: faker.datatype.boolean(),
    name: faker.random.word(),
    path: faker.random.word(),
    showCollapsed: faker.datatype.boolean(),
  })),
  header: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() =>
    faker.random.word(),
  ),
  visualization: {},
})

export const getGetProjectBoardMock = () => ({
  activities: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
    activityID: faker.datatype.number({ min: undefined, max: undefined }),
    description: faker.random.word(),
    isProductive: faker.datatype.boolean(),
    name: faker.random.word(),
    projectID: faker.datatype.number({ min: undefined, max: undefined }),
  })),
  boardConfig: {
    fields: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
      isEditable: faker.datatype.boolean(),
      isVisible: faker.datatype.boolean(),
      name: faker.random.word(),
      path: faker.random.word(),
      showCollapsed: faker.datatype.boolean(),
    })),
    header: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() =>
      faker.random.word(),
    ),
    visualization: {},
  },
  createdAt: (() => faker.date.past())(),
  description: faker.random.word(),
  kanbanSteps: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
    color: faker.random.word(),
    description: faker.random.word(),
    kanbanStepID: faker.datatype.number({ min: undefined, max: undefined }),
    name: faker.random.word(),
    projectID: faker.datatype.number({ min: undefined, max: undefined }),
    stepOrder: faker.datatype.number({ min: undefined, max: undefined }),
    timeTrackable: faker.datatype.boolean(),
  })),
  name: faker.helpers.arrayElement(Object.values(Project)),
  projectID: faker.datatype.number({ min: undefined, max: undefined }),
  teams: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
    createdAt: (() => faker.date.past())(),
    description: faker.random.word(),
    name: faker.random.word(),
    projectID: faker.datatype.number({ min: undefined, max: undefined }),
    teamID: faker.datatype.number({ min: undefined, max: undefined }),
    updatedAt: (() => faker.date.past())(),
  })),
  updatedAt: (() => faker.date.past())(),
  workItemTags: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
    color: faker.random.word(),
    description: faker.random.word(),
    name: faker.random.word(),
    projectID: faker.datatype.number({ min: undefined, max: undefined }),
    workItemTagID: faker.datatype.number({ min: undefined, max: undefined }),
  })),
  workItemTypes: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
    color: faker.random.word(),
    description: faker.random.word(),
    name: faker.random.word(),
    projectID: faker.datatype.number({ min: undefined, max: undefined }),
    workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
  })),
})

export const getGetProjectWorkitemsMock = () =>
  faker.helpers.arrayElement([
    {
      closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      createdAt: (() => faker.date.past())(),
      deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      demoWorkItem: {
        lastMessageAt: (() => faker.date.past())(),
        line: faker.random.word(),
        ref: faker.random.word(),
        reopened: faker.datatype.boolean(),
        workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      },
      description: faker.random.word(),
      kanbanStepID: faker.datatype.number({ min: undefined, max: undefined }),
      members: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        createdAt: (() => faker.date.past())(),
        deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
        email: faker.random.word(),
        firstName: faker.helpers.arrayElement([faker.random.word(), null]),
        fullName: faker.helpers.arrayElement([faker.random.word(), null]),
        hasGlobalNotifications: faker.datatype.boolean(),
        hasPersonalNotifications: faker.datatype.boolean(),
        lastName: faker.helpers.arrayElement([faker.random.word(), null]),
        scopes: faker.helpers.arrayElements(Object.values(Scope)),
        userID: faker.random.word(),
        username: faker.random.word(),
      })),
      metadata: faker.helpers.arrayElement([(() => ({ key: faker.color.hsl() }))(), null]),
      targetDate: (() => faker.date.past())(),
      teamID: faker.datatype.number({ min: undefined, max: undefined }),
      timeEntries: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        activityID: faker.datatype.number({ min: undefined, max: undefined }),
        comment: faker.random.word(),
        durationMinutes: faker.helpers.arrayElement([faker.datatype.number({ min: undefined, max: undefined }), null]),
        start: (() => faker.date.past())(),
        teamID: faker.helpers.arrayElement([faker.datatype.number({ min: undefined, max: undefined }), null]),
        timeEntryID: faker.datatype.number({ min: undefined, max: undefined }),
        userID: faker.random.word(),
        workItemID: faker.helpers.arrayElement([faker.datatype.number({ min: undefined, max: undefined }), null]),
      })),
      title: faker.random.word(),
      updatedAt: (() => faker.date.past())(),
      workItemComments: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        createdAt: (() => faker.date.past())(),
        message: faker.random.word(),
        updatedAt: (() => faker.date.past())(),
        userID: faker.random.word(),
        workItemCommentID: faker.datatype.number({ min: undefined, max: undefined }),
        workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      })),
      workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemTags: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        color: faker.random.word(),
        description: faker.random.word(),
        name: faker.random.word(),
        projectID: faker.datatype.number({ min: undefined, max: undefined }),
        workItemTagID: faker.datatype.number({ min: undefined, max: undefined }),
      })),
      workItemType: {
        color: faker.random.word(),
        description: faker.random.word(),
        name: faker.random.word(),
        projectID: faker.datatype.number({ min: undefined, max: undefined }),
        workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
      },
      workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
    },
    {
      closedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      createdAt: (() => faker.date.past())(),
      deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
      demoTwoWorkItem: {
        customDateForProject2: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
        workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      },
      description: faker.random.word(),
      kanbanStepID: faker.datatype.number({ min: undefined, max: undefined }),
      members: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        createdAt: (() => faker.date.past())(),
        deletedAt: faker.helpers.arrayElement([(() => faker.date.past())(), null]),
        email: faker.random.word(),
        firstName: faker.helpers.arrayElement([faker.random.word(), null]),
        fullName: faker.helpers.arrayElement([faker.random.word(), null]),
        hasGlobalNotifications: faker.datatype.boolean(),
        hasPersonalNotifications: faker.datatype.boolean(),
        lastName: faker.helpers.arrayElement([faker.random.word(), null]),
        scopes: faker.helpers.arrayElements(Object.values(Scope)),
        userID: faker.random.word(),
        username: faker.random.word(),
      })),
      metadata: faker.helpers.arrayElement([(() => ({ key: faker.color.hsl() }))(), null]),
      targetDate: (() => faker.date.past())(),
      teamID: faker.datatype.number({ min: undefined, max: undefined }),
      timeEntries: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        activityID: faker.datatype.number({ min: undefined, max: undefined }),
        comment: faker.random.word(),
        durationMinutes: faker.helpers.arrayElement([faker.datatype.number({ min: undefined, max: undefined }), null]),
        start: (() => faker.date.past())(),
        teamID: faker.helpers.arrayElement([faker.datatype.number({ min: undefined, max: undefined }), null]),
        timeEntryID: faker.datatype.number({ min: undefined, max: undefined }),
        userID: faker.random.word(),
        workItemID: faker.helpers.arrayElement([faker.datatype.number({ min: undefined, max: undefined }), null]),
      })),
      title: faker.random.word(),
      updatedAt: (() => faker.date.past())(),
      workItemComments: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        createdAt: (() => faker.date.past())(),
        message: faker.random.word(),
        updatedAt: (() => faker.date.past())(),
        userID: faker.random.word(),
        workItemCommentID: faker.datatype.number({ min: undefined, max: undefined }),
        workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      })),
      workItemID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemTags: Array.from({ length: faker.datatype.number({ min: 1, max: 10 }) }, (_, i) => i + 1).map(() => ({
        color: faker.random.word(),
        description: faker.random.word(),
        name: faker.random.word(),
        projectID: faker.datatype.number({ min: undefined, max: undefined }),
        workItemTagID: faker.datatype.number({ min: undefined, max: undefined }),
      })),
      workItemType: {
        color: faker.random.word(),
        description: faker.random.word(),
        name: faker.random.word(),
        projectID: faker.datatype.number({ min: undefined, max: undefined }),
        workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
      },
      workItemTypeID: faker.datatype.number({ min: undefined, max: undefined }),
    },
  ])

export const getCreateWorkitemTagMock = () =>
  faker.helpers.arrayElement([
    {
      color: faker.random.word(),
      description: faker.random.word(),
      name: faker.random.word(),
      projectID: faker.datatype.number({ min: undefined, max: undefined }),
      workItemTagID: faker.datatype.number({ min: undefined, max: undefined }),
    },
  ])

export const getProjectMSW = () => [
  rest.post('*/project/:projectName/initialize', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'))
  }),
  rest.get('*/project/:projectName/', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getGetProjectMock()))
  }),
  rest.get('*/project/:projectName/config', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getGetProjectConfigMock()))
  }),
  rest.put('*/project/:projectName/config', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'))
  }),
  rest.get('*/project/:projectName/board', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getGetProjectBoardMock()))
  }),
  rest.get('*/project/:projectName/workitems', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getGetProjectWorkitemsMock()))
  }),
  rest.post('*/project/:projectName/tag/', (_req, res, ctx) => {
    return res(ctx.delay(1000), ctx.status(200, 'Mocked status'), ctx.json(getCreateWorkitemTagMock()))
  }),
]

import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import 'src/assets/css/pulsate.css'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
// import 'regenerator-runtime/runtime'
import {
  MantineProvider,
  Title,
  ColorInput,
  Accordion,
  Button,
  Text,
  Flex,
  useMantineTheme,
  Avatar,
  Group,
  Space,
  Box,
  createTheme,
  localStorageColorSchemeManager,
  Textarea,
} from '@mantine/core'
import { QueryClient } from '@tanstack/react-query'
import { PersistQueryClientProvider, type PersistedClient } from '@tanstack/react-query-persist-client'
import axios from 'axios'
import ProtectedRoute from 'src/components/Permissions/ProtectedRoute'
import { useNotificationAPI } from 'src/hooks/ui/useNotificationAPI'
import { responseInterceptor } from 'src/queries/interceptors'
import { ModalsProvider } from '@mantine/modals'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { Notifications } from '@mantine/notifications'
import { ErrorPage } from 'src/components/ErrorPage/ErrorPage'
import HttpStatus from 'src/utils/httpStatus'
import DynamicForm, {
  selectOptionsBuilder,
  type SelectOptions,
  type DynamicFormOptions,
  InputOptions,
} from 'src/utils/formGeneration'
import type { CreateWorkItemTagRequest, DbWorkItemTag, User, WorkItemRole } from 'src/gen/model'
import type { GetKeys, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import { validateField } from 'src/utils/validation'
import { FormProvider, useForm, useFormState, useWatch } from 'react-hook-form'
import { ajvResolver } from '@hookform/resolvers/ajv'
import dayjs from 'dayjs'
import { ErrorBoundary } from 'react-error-boundary'
import { CodeHighlight } from '@mantine/code-highlight'
import _, { initial } from 'lodash'
import { getGetCurrentUserMock } from 'src/gen/user/user.msw'
import { colorBlindPalette } from 'src/utils/colors'
import { validateJson } from 'src/client-validator/validate'
import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { IconCircle, IconTag } from '@tabler/icons'
import useRenders from 'src/hooks/utils/useRenders'
import { fullFormats } from 'ajv-formats/dist/formats'
import { nameInitials } from 'src/utils/strings'
import WorkItemRoleBadge from 'src/components/Badges/WorkItemRoleBadge'
import { WORK_ITEM_ROLES } from 'src/services/authorization'
import { v4 as uuidv4 } from 'uuid'

import '@mantine/core/styles.css'
import '@mantine/notifications/styles.css'
import '@mantine/code-highlight/styles.css'
import '@mantine/dates/styles.css'
import UserComboboxOption from 'src/components/Combobox/UserComboboxOption'
import { useFormSlice } from 'src/slices/form'
import { useCalloutErrors } from 'src/components/Callout/ErrorCallout'
import { persister } from 'src/idb'
import { queryClient } from 'src/react-query'
import { parseSchemaFields } from 'src/utils/jsonSchema'
import { schemaDefinitions } from 'src/client-validator/gen/meta'
import Project from 'src/views/Project/Project'
import { colorSwatchComponentInputOption } from 'src/components/formGeneration/components'
import { AppTourProvider } from 'src/tours/AppTourProvider'

const schema = {
  properties: {
    base: {
      properties: {
        closed: {
          format: 'date-time',
          type: ['object', 'null'],
        },
        description: {
          type: 'string',
          minLength: 1,
        },
        kanbanStepID: {
          type: 'integer',
        },
        targetDate: {
          format: 'date',
          type: 'object',
        },
        teamID: {
          type: 'integer',
        },
        items: {
          items: {
            properties: {
              items: {
                items: {
                  type: 'string',
                  minLength: 1,
                },
                type: ['array', 'null'],
              },
              userId: {
                items: {
                  type: 'string',
                  minLength: 1,
                },
                type: ['array', 'null'],
              },
              name: {
                type: 'string',
                minLength: 1,
              },
            },
            required: ['userId', 'name'],
            type: 'object',
          },
          type: ['array', 'null'],
        },
        workItemTypeID: {
          type: 'integer',
        },
      },
      required: ['items', 'description', 'workItemTypeID', 'teamID', 'kanbanStepID', 'closed', 'targetDate'],
      type: 'object',
    },
    demoProject: {
      properties: {
        lastMessageAt: {
          format: 'date-time',
          type: 'object',
        },
        line: {
          type: 'string',
          minLength: 1,
        },
        ref: {
          pattern: '^[0-9]{8}$',
          type: 'string',
          minLength: 1,
        },
        reopened: {
          type: 'boolean',
        },
        workItemID: {
          type: 'integer',
        },
      },
      required: ['workItemID', 'ref', 'line', 'lastMessageAt', 'reopened'],
      type: 'object',
    },
    members: {
      items: {
        properties: {
          role: {
            title: 'WorkItem role',
            type: 'string',
            minLength: 1,
            'x-generated': '-',
            enum: ['preparer', 'reviewer'],
            description: "represents a database 'work_item_role'",
          },
          userID: {
            type: 'string',
            minLength: 1,
          },
        },
        required: ['userID', 'role'],
        type: 'object',
      },
      type: ['array', 'null'],
    },
    tagIDs: {
      items: {
        type: 'integer',
      },
      type: ['array', 'null'],
    },
  },
  required: ['demoProject', 'base', 'tagIDs', 'members'],
  type: 'object',
  'x-postgen-struct': 'RestDemoWorkItemCreateRequest',
}

function ErrorFallback({ error }: any) {
  return (
    <div role="alert">
      <p>Something went wrong:</p>
      <pre style={{ color: 'red' }}>{error.message}</pre>
    </div>
  )
}

const Layout = React.lazy(() => import('./components/Layout/Layout'))
const LandingPage = React.lazy(() => import('./views/LandingPage/LandingPage'))
const UserPermissionsPage = React.lazy(() => import('src/views/Settings/UserPermissionsPage/UserPermissionsPage'))
const ProjectManagementPage = React.lazy(() => import('src/views/Admin/ProjectManagementPage/ProjectManagementPage'))

const uuids = [
  'fcd252dc-72a4-4514-bdd1-3cac573a5fac',
  '120cb364-2b18-49fb-b505-568834614c5d',
  'bdab07d6-c2b4-44b0-b6d0-e87f62037cc1',
  'ad52daf8-9bad-4671-b3f1-535178b0346e',
  '3e82b3a5-5757-4860-8bf7-2e7962534328',
  'd59d3a5c-b99f-40aa-9419-75a2bbb0fd52',
]

// TODO: /users where deleted_at null
// will be used on generated filterable mantine datatable table as in
const members = [...Array(10)].map((x, i) => {
  const user = getGetCurrentUserMock()
  user.email = `${i}@mail.com`
  user.userID = i < uuids.length ? uuids[i]! : uuidv4()
  return user
})

const tags = [...Array(10)].map((x, i) => {
  const tag: DbWorkItemTag = {
    name: `tag #${i}`,
    color: _.sample(colorBlindPalette)!,
    workItemTagID: i,
    projectID: 1,
    description: `description for tag #${i}`,
  } // TODO: get workitem tags endpoint
  return tag
})

const colorSchemeManager = localStorageColorSchemeManager({ key: 'theme' })

const userIdSelectOption = selectOptionsBuilder({
  type: 'select',
  values: members,
  //  TODO: transformers can be reusable between forms. could simply become
  //  {
  //   type: "select"
  //   values: ...
  //   ...userIdFormTransformers
  // }
  optionTransformer(el) {
    return <UserComboboxOption user={el} />
  },
  formValueTransformer(el) {
    return el.userID
  },
  pillTransformer(el) {
    return <>{el.email}</>
  },
  searchValueTransformer(el) {
    return `${el.email} ${el.fullName} ${el.username}`
  },
})

export default function App() {
  useEffect(() => {
    document.body.style.background = 'none !important'
  }, [])

  const { verifyNotificationPermission } = useNotificationAPI()
  const [notificationWarningSent, setNotificationWarningSent] = useState(false)

  useEffect(() => {
    if (!notificationWarningSent) {
      verifyNotificationPermission()
      setNotificationWarningSent(true)
    }
  }, [verifyNotificationPermission, notificationWarningSent])

  const formInitialValues = {
    base: {
      items: [
        {
          items: ['0001', '0002'],
          userId: ['120cb364-2b18-49fb-b505-568834614c5d', 'fcd252dc-72a4-4514-bdd1-3cac573a5fac'],
          name: 'item-1',
        },
        { items: ['0011', '0012'], userId: ['badid', 'badid2'], name: 'item-2' },
      ],
      // closed: dayjs('2023-03-24T20:42:00.000Z').toDate(),
      targetDate: dayjs('2023-02-22').toDate(),
      description: 'some text',
      kanbanStepID: 1,
      teamID: 1,
      metadata: {},
      // title: {},
      workItemTypeID: 1,
    },
    // TODO: formGeneration must not assume options do exist, else all fails catastrophically.
    // it's not just checking types...
    // 1. move callout errors state to zustand, and create callout warnings too
    // 2.(sol 1) if option not found for initial data, remove from form values
    // and show persistent callout _warning_ that X was deleted since it was not found.
    // it should update the form but show callout error saying ignoring bad type in `formField`, in this case `tagIDs.1`
    // 2. (sol 2 which wont work) leave form as is and validate on first render will not catch errors for options not found, if type is right...
    tagIDs: [1, 2, 'badid'], // {"invalidParams":{"name":"tagIDs.1","reason":"must be integer"} and we can set invalid manually via component id (which will be `input-tagIDs.1` )
    tagIDsMultiselect: null,
    // tagIDs: [0, 5, 8],
    demoProject: {
      lastMessageAt: dayjs('2023-03-24T20:42:00.000Z').toDate(),
      ref: '12341234',
      workItemID: 1,
      reopened: false, // for create will ignore field for form gen
    },
    members: [{ userID: '2ae4bc55-5c26-4b93-8dc7-e2bc0e9e3a65' }, { role: 'preparer', userID: 'bad userID' }],
  } as TestTypes.DemoWorkItemCreateRequest

  const form = useForm<TestTypes.DemoWorkItemCreateRequest>({
    resolver: ajvResolver(schema as any, {
      strict: false,
      formats: fullFormats,
    }),
    mode: 'all',
    reValidateMode: 'onChange',
    defaultValues: formInitialValues ?? {},
    // shouldUnregister: true, // defaultValues will not be merged against submission result.
  })
  const createWorkItemTagForm = useForm<CreateWorkItemTagRequest>({
    resolver: ajvResolver(schema as any, {
      strict: false,
      formats: fullFormats,
    }),
    mode: 'all',
    reValidateMode: 'onChange',
  })
  const { register, handleSubmit, control, formState } = form
  const errors = formState.errors
  const formSLice = useFormSlice()
  const [errorSet, seterrorSet] = useState(false)
  const { extractCalloutErrors, setCalloutErrors, calloutErrors, extractCalloutTitle } =
    useCalloutErrors('demoWorkItemCreateForm')

  useEffect(() => {
    console.log('errors')
    console.log(errors)
    // if (Object.keys(errors).length > 0 && !errorSet) {
    // setCalloutErrors('Validation error')

    // console.log('errors')
    // console.log(errors)

    // setCalloutErrors('Validation error')
    // seterrorSet(true)
    // // console.log(formSLice.callout[formName])
    // // console.log(`form has errors: ${Object.keys(errors).length > 0}`)
    // }
  }, [formState])

  // useEffect(() => {
  //   console.log(demoWorkItemCreateForm.values)
  //   try {
  //     DemoWorkItemCreateRequestDecoder.decode(demoWorkItemCreateForm.values)
  //   } catch (error) {
  //     console.error(JSON.stringify(error.validationErrors.errors))
  //   }
  // }, [demoWorkItemCreateForm])

  type ExcludedFormKeys = 'base.metadata' | 'tagIDsMultiselect'

  return (
    <PersistQueryClientProvider client={queryClient} persistOptions={{ persister }}>
      <MantineProvider
        colorSchemeManager={colorSchemeManager}
        theme={createTheme({
          shadows: {
            md: '1px 1px 3px rgba(0, 0, 0, .25)',
            xl: '5px 5px 3px rgba(0, 0, 0, .25)',
          },
          fontFamily: 'Catamaran, Arial, sans-serif',
        })}
      >
        <ModalsProvider
          labels={{ confirm: 'Submit', cancel: 'Cancel' }}
          modalProps={{ styles: { root: { marginTop: '100px', zIndex: 20000 } } }}
        >
          <Notifications />
          <BrowserRouter basename="">
            <React.Suspense
              fallback={<div style={{ backgroundColor: 'rgb(20, 21, 25)', height: '100vh', width: '100vw' }} />}
            >
              <AppTourProvider>
                <Layout>
                  <Routes>
                    <Route path="/project" element={<Project />} />
                    <Route
                      path="/"
                      element={
                        <React.Suspense fallback={<FallbackLoading />}>
                          {/* <LandingPage /> */}
                          <Title size={20}>This form has been automatically generated from an openapi spec</Title>
                          <Button
                            onClick={(e) => {
                              try {
                                console.log(errors?.demoProject)
                                // const r = demoWorkItemCreateForm.validate()
                                // console.log({ r })
                                // DemoWorkItemCreateRequestDecoder.decode(demoWorkItemCreateForm.values)
                              } catch (error) {
                                console.error(JSON.stringify(error?.validationErrors?.errors))
                              }
                            }}
                          >
                            Validate form
                          </Button>
                          <FormProvider {...form}>
                            <DynamicForm<TestTypes.DemoWorkItemCreateRequest, ExcludedFormKeys>
                              onSubmit={(e) => {
                                e.preventDefault()
                                form.handleSubmit(
                                  (data) => {
                                    console.log({ data })
                                  },
                                  (errors) => {
                                    console.log({ errors })
                                  },
                                )(e)
                              }}
                              formName="demoWorkItemCreateForm"
                              // schemaFields will come from `parseSchemaFields(schema.RestDemo... OR  asConst(jsonSchema.definitions.<...>))`
                              // using this hardcoded for testing purposes
                              schemaFields={{
                                base: { isArray: false, required: true, type: 'object' },
                                'base.closed': { type: 'date-time', required: false, isArray: false },
                                'base.description': { type: 'string', required: true, isArray: false },
                                'base.kanbanStepID': { type: 'integer', required: true, isArray: false },
                                'base.targetDate': { type: 'date', required: true, isArray: false },
                                'base.teamID': { type: 'integer', required: true, isArray: false },
                                'base.items': { type: 'object', required: false, isArray: true },
                                'base.items.name': { type: 'string', required: true, isArray: false },
                                'base.items.userId': { type: 'string', required: false, isArray: true },
                                'base.items.items': { type: 'string', required: false, isArray: true },
                                'base.workItemTypeID': { type: 'integer', required: true, isArray: false },
                                demoProject: { isArray: false, required: true, type: 'object' },
                                'demoProject.lastMessageAt': { type: 'date-time', required: true, isArray: false },
                                'demoProject.line': { type: 'string', required: true, isArray: false },
                                'demoProject.ref': { type: 'string', required: true, isArray: false },
                                'demoProject.reopened': { type: 'boolean', required: true, isArray: false },
                                'demoProject.workItemID': { type: 'integer', required: true, isArray: false },
                                members: { type: 'object', required: false, isArray: true },
                                'members.role': { type: 'string', required: true, isArray: false },
                                'members.userID': { type: 'string', required: true, isArray: false },
                                tagIDs: { type: 'integer', required: false, isArray: true },
                              }}
                              options={{
                                // since labels is mandatory, instead of duplicating with ignore: U[] just
                                // check if labels hasOwnProperty fieldKey and if not exclude from form.
                                labels: {
                                  base: null,
                                  'base.closed': 'Closed',
                                  'base.description': 'Description',
                                  // 'base.metadata': 'metadata', // ignored -> not a key
                                  'base.kanbanStepID': 'Kanban step', // if using KanbanStep transformer, then "Kanban step", "Kanban step name", etc.
                                  'base.targetDate': 'Target date',
                                  'demoProject.reopened': 'Reopened',
                                  'base.teamID': 'Team',
                                  'base.items': 'Items',
                                  'base.items.name': 'Name',
                                  'base.items.items': 'Items',
                                  'base.items.userId': 'User',
                                  'base.workItemTypeID': 'Type',
                                  demoProject: null,
                                  'demoProject.lastMessageAt': 'Last message at',
                                  'demoProject.line': 'Line',
                                  'demoProject.ref': 'Ref',
                                  'demoProject.workItemID': 'Work item',
                                  members: 'Members',
                                  'members.role': 'Role',
                                  'members.userID': 'User',
                                  tagIDs: 'Tags',
                                },
                                renderOrderPriority: ['tagIDs', 'members'],
                                accordion: {
                                  'base.items': {
                                    defaultOpen: true,
                                    title: (
                                      <Flex align="center" gap={10}>
                                        <IconTag size={16} />
                                        <Text fw={700} size={'md'}>
                                          Items
                                        </Text>
                                      </Flex>
                                    ),
                                  },
                                },
                                defaultValues: {
                                  'demoProject.line': '1111',
                                  'members.role': 'preparer',
                                },
                                selectOptions: {
                                  'members.userID': userIdSelectOption,
                                  'base.items.userId': userIdSelectOption,
                                  tagIDs: selectOptionsBuilder({
                                    type: 'multiselect',
                                    searchValueTransformer(el) {
                                      return el.name
                                    },
                                    values: tags,
                                    optionTransformer(el) {
                                      return (
                                        <Group align="center">
                                          <Flex align="center" gap={12} justify="center">
                                            <IconCircle size={12} fill={el.color} />
                                            <div>{el.name}</div>
                                          </Flex>
                                        </Group>
                                      )
                                    },
                                    formValueTransformer(el) {
                                      return el.workItemTagID
                                    },
                                    pillTransformer(el) {
                                      return <div>{el.name}</div>
                                    },
                                    labelColor(el) {
                                      return el.color
                                    },
                                  }),
                                  'members.role': selectOptionsBuilder({
                                    type: 'select',
                                    values: WORK_ITEM_ROLES,
                                    optionTransformer: (el) => {
                                      return <WorkItemRoleBadge role={el} />
                                    },
                                    formValueTransformer(el) {
                                      return el
                                    },
                                    pillTransformer(el) {
                                      return <WorkItemRoleBadge role={el} />
                                    },
                                  }),
                                },
                                input: {
                                  'demoProject.line': {
                                    component: colorSwatchComponentInputOption,
                                  },
                                },
                                // these should probably be all required later, to ensure formField is never used.
                                propsOverride: {
                                  'demoProject.line': {
                                    description: 'This is some help text for a disabled field.',
                                    disabled: true,
                                  },
                                },
                              }} // satisfies DynamicFormOptions<TestTypes.DemoWorkItemCreateRequest, ExcludedFormKeys> // not needed anymore for some reason
                            />
                          </FormProvider>
                        </React.Suspense>
                      }
                    />
                    <Route
                      path="/settings/user-permissions-management"
                      element={
                        <React.Suspense fallback={<FallbackLoading />}>
                          <ProtectedRoute>
                            <UserPermissionsPage />
                          </ProtectedRoute>
                        </React.Suspense>
                      }
                    />
                    <Route
                      path="/admin/project-management"
                      element={
                        <React.Suspense fallback={<FallbackLoading />}>
                          <ProtectedRoute>
                            <ProjectManagementPage />
                          </ProtectedRoute>
                        </React.Suspense>
                      }
                    />
                    <Route
                      path="*"
                      element={
                        <React.Suspense fallback={<FallbackLoading />}>
                          <ProtectedRoute>
                            <ErrorPage status={HttpStatus.NOT_FOUND_404} />
                          </ProtectedRoute>
                        </React.Suspense>
                      }
                    />
                  </Routes>
                </Layout>
              </AppTourProvider>
            </React.Suspense>
          </BrowserRouter>
        </ModalsProvider>
      </MantineProvider>
      {!import.meta.env.PROD && <ReactQueryDevtools initialIsOpen />}
    </PersistQueryClientProvider>
  )
}

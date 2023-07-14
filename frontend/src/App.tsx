import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import 'src/assets/css/pulsate.css'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
// import 'regenerator-runtime/runtime'
import {
  ColorSchemeProvider,
  type ColorScheme,
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
} from '@mantine/core'
import { QueryClient } from '@tanstack/react-query'
import { PersistQueryClientProvider, type PersistedClient, type Persister } from '@tanstack/react-query-persist-client'
import axios from 'axios'
import { del, get, set } from 'idb-keyval'
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
} from 'src/utils/formGeneration'
import type { DbWorkItemTag, RestDemoWorkItemCreateRequest, User } from 'src/gen/model'
import type { GetKeys, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import { RestDemoWorkItemCreateRequestDecoder } from 'src/client-validator/gen/decoders'
import { validateField } from 'src/utils/validation'
import { FormProvider, useForm, useWatch } from 'react-hook-form'
import { ajvResolver } from '@hookform/resolvers/ajv'
import dayjs from 'dayjs'
import { ErrorBoundary } from 'react-error-boundary'
import { Prism } from '@mantine/prism'
import { initial } from 'lodash'
import { getGetCurrentUserMock } from 'src/gen/user/user.msw'
import { colorBlindPalette } from 'src/utils/colors'
import { validateJson } from 'src/client-validator/validate'
import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { IconTag } from '@tabler/icons'
import JSON_SCHEMA from 'src/client-validator/gen/dereferenced-schema.json'
import useRenders from 'src/hooks/utils/useRenders'
import { fullFormats } from 'ajv-formats/dist/formats'
import { nameInitials } from 'src/utils/strings'

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
              name: {
                type: 'string',
                minLength: 1,
                $schema: 'http://json-schema.org/draft-04/schema#',
              },
            },
            required: ['items', 'name'],
            type: 'object',
            $schema: 'http://json-schema.org/draft-04/schema#',
          },
          type: ['array', 'null'],
        },
        workItemTypeID: {
          type: 'integer',
        },
      },
      required: ['items', 'description', 'workItemTypeID', 'teamID', 'kanbanStepID', 'closed', 'targetDate'],
      type: 'object',
      $schema: 'http://json-schema.org/draft-04/schema#',
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
      $schema: 'http://json-schema.org/draft-04/schema#',
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
            $schema: 'http://json-schema.org/draft-04/schema#',
          },
          userID: {
            type: 'string',
            minLength: 1,
            $schema: 'http://json-schema.org/draft-04/schema#',
          },
        },
        required: ['userID', 'role'],
        type: 'object',
        $schema: 'http://json-schema.org/draft-04/schema#',
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

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      cacheTime: 1000 * 60 * 5, // 5 min
      // cacheTime: 0,
      refetchOnWindowFocus: false,
      refetchOnMount: false,
      staleTime: Infinity,
      keepPreviousData: true,
    },
    mutations: {
      cacheTime: 1000 * 60 * 5, // 5 minutes
    },
  },
  // queryCache,
})

// axios.interceptors.request.use(requestInterceptor, function (error) {
//   return Promise.reject(error)
// })
axios.interceptors.response.use(responseInterceptor, function (error) {
  return Promise.reject(error)
})

/**
 * Creates an Indexed DB persister
 * @see https://developer.mozilla.org/en-US/docs/Web/API/IndexedDB_API
 */
function createIDBPersister(idbValidKey: IDBValidKey = 'reactQuery') {
  return {
    persistClient: async (client: PersistedClient) => {
      set(idbValidKey, client)
    },
    restoreClient: async () => {
      return await get<PersistedClient>(idbValidKey)
    },
    removeClient: async () => {
      await del(idbValidKey)
    },
  } as Persister
}

export const persister = createIDBPersister()

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

const members = [...Array(10)].map((x, i) => {
  const user = getGetCurrentUserMock()
  user.email = `${i}@mail.com`
  user.userID = `${i}ae4bc55-5c26-4b93-8dc7-e2bc0e9e3a65`
  return user
})

const tags = [...Array(10)].map((x, i) => {
  const tag: DbWorkItemTag = {
    name: `${i} tag`,
    color: `#${i}34236`,
    workItemTagID: i,
    projectID: 1,
    description: 'description',
  } // TODO: get workitem tags endpoint
  return tag
})

const userIdOptionTransformer = (el: User) => {
  return (
    <Group noWrap spacing="lg" align="center" mr={8}>
      <Flex align={'center'}>
        <Avatar size={'28px'} radius="xl" data-test-id="header-profile-avatar" alt={el?.username}>
          {nameInitials(el?.fullName || '')}
        </Avatar>
        <Space p={5} />
      </Flex>
      <Box ml={'auto'}>{el?.email}</Box>
    </Group>
  )
}

export default function App() {
  const theme = useMantineTheme()
  const [colorScheme, setColorScheme] = useState<ColorScheme>(
    localStorage.getItem('theme') === 'dark' ? 'dark' : 'light',
  )
  const toggleColorScheme = (value?: ColorScheme) => {
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'))
  }

  useEffect(() => {
    document.body.style.background = 'none !important'
  }, [])

  useEffect(() => {
    localStorage.setItem('theme', colorScheme)
  }, [colorScheme])

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
        { items: ['0001', '0002'], name: 'item-1' },
        { items: ['0011', '0012'], name: 'item-2' },
      ],
      // closed: dayjs('2023-03-24T20:42:00.000Z').toDate(),
      // targetDate: dayjs('2023-02-22').toDate(),
      description: 'some text',
      kanbanStepID: 1,
      teamID: 1,
      metadata: {},
      // title: {},
      workItemTypeID: 1,
    },
    // TODO: need to check runtime type, else all fails catastrophically.
    // it should update the form but show callout error saying ignoring bad type in `formField`, in this case `tagIDs.1`
    // tagIDs: [1, 'fsfefes'], // {"invalidParams":{"name":"tagIDs.1","reason":"must be integer"} and we can set invalid manually via component id (which will be `input-tagIDs.1` )
    demoProject: {
      lastMessageAt: dayjs('2023-03-24T20:42:00.000Z').toDate(),
      ref: '12341234',
      workItemID: 1,
      reopened: false, // for create will ignore field for form gen
    },
    tagIDs: [0, 5, 8],
    members: [{ userID: '2ae4bc55-5c26-4b93-8dc7-e2bc0e9e3a65' }, { role: 'preparer', userID: 'user 2' }],
  } as TestTypes.RestDemoWorkItemCreateRequest

  const form = useForm<TestTypes.RestDemoWorkItemCreateRequest>({
    resolver: ajvResolver(schema as any, {
      strict: false,
      formats: fullFormats,
    }),
    mode: 'all',
    defaultValues: formInitialValues ?? {},
    // shouldUnregister: true, // defaultValues will not be merged against submission result.
  })

  const {
    register,
    handleSubmit,
    control,
    formState: { errors, defaultValues },
  } = form

  // useEffect(() => {
  //   console.log(demoWorkItemCreateForm.values)
  //   try {
  //     RestDemoWorkItemCreateRequestDecoder.decode(demoWorkItemCreateForm.values)
  //   } catch (error) {
  //     console.error(JSON.stringify(error.validationErrors.errors))
  //   }
  // }, [demoWorkItemCreateForm])

  type ExcludedFormKeys = 'base.metadata' | 'tagIDsMultiselect'

  return (
    <PersistQueryClientProvider client={queryClient} persistOptions={{ persister }}>
      <ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
        <MantineProvider
          withGlobalStyles
          withNormalizeCSS
          theme={{
            colorScheme,
            shadows: {
              md: '1px 1px 3px rgba(0, 0, 0, .25)',
              xl: '5px 5px 3px rgba(0, 0, 0, .25)',
            },
            fontFamily: 'Catamaran, Arial, sans-serif',
          }}
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
                <Layout>
                  <Routes>
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
                                // RestDemoWorkItemCreateRequestDecoder.decode(demoWorkItemCreateForm.values)
                              } catch (error) {
                                console.error(JSON.stringify(error?.validationErrors?.errors))
                              }
                            }}
                          >
                            Validate form
                          </Button>
                          <FormProvider {...form}>
                            <DynamicForm<TestTypes.RestDemoWorkItemCreateRequest, ExcludedFormKeys>
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
                                  'base.closed': 'closed',
                                  'base.description': 'description',
                                  // 'base.metadata': 'metadata', // ignored -> not a key
                                  'base.kanbanStepID': 'kanbanStepID', // if using KanbanStep transformer, then "Kanban step", "Kanban step name", etc.
                                  'base.targetDate': 'targetDate',
                                  'demoProject.reopened': 'reopened',
                                  'base.teamID': 'teamID',
                                  'base.items': 'items',
                                  'base.items.name': 'name',
                                  'base.items.items': 'items',
                                  'base.workItemTypeID': 'workItemTypeID',
                                  demoProject: null,
                                  'demoProject.lastMessageAt': 'lastMessageAt',
                                  'demoProject.line': 'line',
                                  'demoProject.ref': 'ref',
                                  'demoProject.workItemID': 'workItemID',
                                  members: 'members',
                                  'members.role': 'role',
                                  'members.userID': 'User',
                                  tagIDs: 'tagIDs',
                                },
                                accordion: {
                                  'base.items': {
                                    defaultOpen: true,
                                    title: (
                                      <Flex align="center" gap={10}>
                                        <IconTag size={16} />
                                        <Text weight={700} size={'md'} color={theme.primaryColor}>
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
                                  'members.userID': selectOptionsBuilder({
                                    type: 'select',
                                    values: members,
                                    // transformers can be reusable between forms. could simply become
                                    //  {
                                    //   type: "select"
                                    //   values: ...
                                    //   ...userIdFormTransformers
                                    // }
                                    optionTransformer: userIdOptionTransformer,
                                    formValueTransformer(el) {
                                      return el.userID
                                    },
                                    labelTransformer(el) {
                                      return <>{el.email}</>
                                    },
                                  }),
                                  tagIDs: selectOptionsBuilder({
                                    type: 'multiselect',
                                    values: tags,
                                    optionTransformer(el) {
                                      return (
                                        <Group noWrap spacing="lg" align="center">
                                          <Flex align={'center'}></Flex>
                                          <div>{el?.name}</div>
                                        </Group>
                                      )
                                    },
                                    formValueTransformer(el) {
                                      return el.workItemTagID
                                    },
                                    labelTransformer(el) {
                                      return <>{el.name} label</>
                                    },
                                    labelColor(el) {
                                      return el.color
                                    },
                                  }),
                                },
                                input: {
                                  'demoProject.line': {
                                    component: (
                                      <ColorInput
                                        placeholder="Pick color"
                                        disallowInput
                                        withPicker={false}
                                        swatches={colorBlindPalette}
                                        styles={{ root: { width: '100%' } }}
                                      />
                                    ),
                                  },
                                },
                                // these should probably be all required later, to ensure formField is never used.
                                propsOverride: {
                                  'demoProject.line': {
                                    label: 'Line',
                                    description: 'This is some help text.',
                                  },
                                },
                              }} // satisfies DynamicFormOptions<TestTypes.RestDemoWorkItemCreateRequest, ExcludedFormKeys> // not needed anymore for some reason
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
              </React.Suspense>
            </BrowserRouter>
          </ModalsProvider>
        </MantineProvider>
      </ColorSchemeProvider>
      {!import.meta.env.PROD && <ReactQueryDevtools initialIsOpen />}
    </PersistQueryClientProvider>
  )
}

import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import 'src/assets/css/pulsate.css'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
// import 'regenerator-runtime/runtime'
import { ColorSchemeProvider, type ColorScheme, MantineProvider } from '@mantine/core'
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
import { DynamicForm } from 'src/utils/formGeneration'
import type { RestDemoWorkItemCreateRequest } from 'src/gen/model'
import { type FieldPath } from 'react-hook-form'
import type { RecursiveKeyOfArray } from 'src/types/utils'
import { RestDemoWorkItemCreateRequestDecoder } from 'src/client-validator/gen/decoders'
import { validateField } from 'src/utils/validation'
import { useForm } from '@mantine/form'
import DemoWorkItemForm from 'src/components/forms/DemoProjectWorkItemForm'
import dayjs from 'dayjs'

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

const Layout = React.lazy(() => import('./components/Layout/Layout'))
const LandingPage = React.lazy(() => import('./views/LandingPage/LandingPage'))
const UserPermissionsPage = React.lazy(() => import('src/views/Settings/UserPermissionsPage/UserPermissionsPage'))
const ProjectManagementPage = React.lazy(() => import('src/views/Admin/ProjectManagementPage/ProjectManagementPage'))

export default function App() {
  const [colorScheme, setColorScheme] = useState<ColorScheme>(
    localStorage.getItem('theme') === 'dark' ? 'dark' : 'light',
  )
  const toggleColorScheme = (value?: ColorScheme) => {
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'))
  }

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

  type RestDemoWorkItemCreateRequestFormField =
    // hack to use 'members.role' instead of 'members.??.role'
    FieldPath<RestDemoWorkItemCreateRequest> | RecursiveKeyOfArray<RestDemoWorkItemCreateRequest['members'], 'members'>

  const demoWorkItemCreateForm = useForm<RestDemoWorkItemCreateRequest>({
    // TODO: simple function to initialize top level with empty object if property type === object
    // now that we have json schema dereferenced
    initialValues: {
      base: {
        closed: dayjs().toDate(),
      },
      demoProject: {},
    } as RestDemoWorkItemCreateRequest,
    validateInputOnChange: true,
    validate: {
      // TODO: should be able to validate whole nested objects at once.
      base: (v, vv, path) => validateField(RestDemoWorkItemCreateRequestDecoder, path, vv),
      // members: (v, vv, path) => validateField(RestDemoWorkItemCreateRequestDecoder, path, vv), // TODO: foreach validate
    },
  })

  useEffect(() => {
    console.log(demoWorkItemCreateForm)
  }, [demoWorkItemCreateForm])

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
                          <DynamicForm<RestDemoWorkItemCreateRequestFormField, RestDemoWorkItemCreateRequest>
                            form={demoWorkItemCreateForm}
                            // schemaFields will come from `parseSchemaFields(schema.RestDemo...)`
                            schemaFields={{
                              base: { isArray: false, required: true, type: 'object' },
                              'base.closed': { type: 'date-time', required: true, isArray: false },
                              'base.description': { type: 'string', required: true, isArray: false },
                              'base.kanbanStepID': { type: 'integer', required: true, isArray: false },
                              'base.metadata': { type: 'integer', required: true, isArray: true },
                              'base.targetDate': { type: 'date-time', required: true, isArray: false },
                              'base.teamID': { type: 'integer', required: true, isArray: false },
                              'base.title': { type: 'string', required: true, isArray: false },
                              'base.workItemTypeID': { type: 'integer', required: true, isArray: false },
                              demoProject: { isArray: false, required: true, type: 'object' },
                              'demoProject.lastMessageAt': { type: 'date-time', required: true, isArray: false },
                              'demoProject.line': { type: 'string', required: true, isArray: false },
                              'demoProject.ref': { type: 'string', required: true, isArray: false },
                              'demoProject.reopened': { type: 'boolean', required: true, isArray: false },
                              'demoProject.workItemID': { type: 'integer', required: true, isArray: false },
                              members: { type: 'object', required: true, isArray: true },
                              'members.role': { type: 'string', required: true, isArray: false },
                              'members.userID': { type: 'string', required: true, isArray: false },
                              tagIDs: { type: 'integer', required: true, isArray: true },
                            }}
                            options={{
                              defaultValue: {
                                'demoProject.line': '534543523', // should fail due to TypeOf
                                members: [{ role: 'preparer', userID: 'c446259c-1083-4212-98fe-bd080c41e7d7' }],
                              },
                            }}
                          />
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

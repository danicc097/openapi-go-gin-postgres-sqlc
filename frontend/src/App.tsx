import React, { useEffect, useMemo, useState } from 'react'
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom'
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
import { parseSchemaFields } from 'src/utils/jsonSchema'
import { schemaDefinitions } from 'src/client-validator/gen/meta'
import Project from 'src/views/Project/Project'
import { colorSwatchComponentInputOption } from 'src/components/formGeneration/components'
import { AppTourProvider } from 'src/tours/AppTourProvider'
import { useGetPaginatedUsers } from 'src/gen/user/user'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { reactQueryDefaultAppOptions } from 'src/react-query'
import DemoGeneratedForm from 'src/views/DemoGeneratedForm/DemoGeneratedForm'

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

const colorSchemeManager = localStorageColorSchemeManager({ key: 'theme' })

const routes = {
  '/project': <Project />,
  '/': <h1>Home</h1>,
  '/settings/user-permissions-management': (
    <ProtectedRoute>
      <UserPermissionsPage />
    </ProtectedRoute>
  ),
  '/demo/generated-form': <DemoGeneratedForm />,
  // TODO: update from eui
  // '/admin/project-management': (
  //   <ProtectedRoute>
  //     <ProjectManagementPage />
  //   </ProtectedRoute>
  // ),
  '*': (
    <ProtectedRoute>
      <ErrorPage status={HttpStatus.NOT_FOUND_404} />
    </ProtectedRoute>
  ),
}

export default function App() {
  useEffect(() => {
    document.body.style.background = 'none !important' // body was preventing flashes
  }, [])

  const { verifyNotificationPermission } = useNotificationAPI()
  const [notificationWarningSent, setNotificationWarningSent] = useState(false)

  useEffect(() => {
    if (!notificationWarningSent) {
      verifyNotificationPermission()
      setNotificationWarningSent(true)
    }
  }, [verifyNotificationPermission, notificationWarningSent])

  return (
    <>
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
                    {Object.entries(routes).map(([path, component], index) => (
                      <Route
                        key={index}
                        path={path}
                        element={
                          path === '/' ? (
                            <React.Suspense fallback={<FallbackLoading />}>
                              <h1>Home</h1>
                              <ul>
                                {Object.keys(routes).map((routePath) => (
                                  <li key={routePath}>
                                    <Link to={routePath}>{routePath}</Link>
                                  </li>
                                ))}
                              </ul>
                            </React.Suspense>
                          ) : (
                            <React.Suspense fallback={<FallbackLoading />}>{component}</React.Suspense>
                          )
                        }
                      />
                    ))}
                  </Routes>
                </Layout>
              </AppTourProvider>
            </React.Suspense>
          </BrowserRouter>
        </ModalsProvider>
      </MantineProvider>
      {!import.meta.env.PROD && <ReactQueryDevtools initialIsOpen />}
    </>
  )
}

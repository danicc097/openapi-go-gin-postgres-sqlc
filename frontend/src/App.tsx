import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import 'src/assets/css/pulsate.css'
import '@mantine/core/styles.css'
import '@mantine/notifications/styles.css'
import '@mantine/code-highlight/styles.css'
import '@mantine/dates/styles.css'
import 'mantine-react-table/styles.css' //import MRT styles

import React, { useEffect, useState } from 'react'
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
// import 'regenerator-runtime/runtime'
import { MantineProvider, createTheme, localStorageColorSchemeManager } from '@mantine/core'
import ProtectedRoute from 'src/components/Permissions/ProtectedRoute'
import { useNotificationAPI } from 'src/hooks/ui/useNotificationAPI'
import { ModalsProvider } from '@mantine/modals'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { Notifications } from '@mantine/notifications'
import { ErrorPage } from 'src/components/ErrorPage/ErrorPage'
import HttpStatus from 'src/utils/httpStatus'
import _ from 'lodash'

import Project from 'src/views/Project/Project'
import { AppTourProvider } from 'src/tours/AppTourProvider'
import DemoGeneratedForm from 'src/views/DemoGeneratedForm/DemoGeneratedForm'
import DemoMantineReactTable from 'src/views/DemoMantineReactTable/DemoMantineReactTable'

function ErrorFallback({ error }: any) {
  return (
    <div role="alert">
      <p>Something went wrong:</p>
      <pre style={{ color: 'red' }}>{error.message}</pre>
    </div>
  )
}

import utc from 'dayjs/plugin/utc'
import relativeTime from 'dayjs/plugin/relativeTime'
import dayjs from 'dayjs'
dayjs.extend(utc)
dayjs.extend(relativeTime)

const Layout = React.lazy(() => import('./components/Layout/Layout'))
const LandingPage = React.lazy(() => import('./views/LandingPage/LandingPage'))
const UserPermissionsPage = React.lazy(() => import('src/views/Settings/UserPermissionsPage/UserPermissionsPage'))
const ProjectManagementPage = React.lazy(() => import('src/views/Admin/ProjectManagementPage/ProjectManagementPage'))

const colorSchemeManager = localStorageColorSchemeManager({ key: 'theme' })

const routes = Object.freeze({
  '/project': <Project />,
  '/': <h1>Home</h1>,
  '/settings/user-permissions-management': (
    <ProtectedRoute>
      <UserPermissionsPage />
    </ProtectedRoute>
  ),
  '/demo/generated-form': <DemoGeneratedForm />,
  '/demo/mantine-react-table': <DemoMantineReactTable />,
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
})

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
        defaultColorScheme="dark"
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

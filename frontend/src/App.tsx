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
                          <LandingPage />
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

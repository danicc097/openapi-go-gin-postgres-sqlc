import { BrowserRouter, Routes, Route } from 'react-router-dom'
import React, { useEffect, useState } from 'react'
import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
// import 'regenerator-runtime/runtime'
import { EuiProvider, useEuiTheme } from '@elastic/eui'
import { useUISlice } from 'src/slices/ui'
import { useNotificationAPI } from 'src/hooks/ui/useNotificationAPI'

const Layout = React.lazy(() => import('./components/Layout/Layout'))
const LandingPage = React.lazy(() => import('./views/LandingPage/LandingPage'))

export default function App() {
  const theme = useUISlice((state) => state?.theme)
  const { verifyNotificationPermission } = useNotificationAPI()
  const [notificationWarningSent, setNotificationWarningSent] = useState(false)

  useEffect(() => {
    if (!notificationWarningSent) {
      verifyNotificationPermission()
      setNotificationWarningSent(true)
    }
  }, [verifyNotificationPermission, notificationWarningSent])

  return (
    <EuiProvider colorMode={theme}>
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
            </Routes>
          </Layout>
        </React.Suspense>
      </BrowserRouter>
    </EuiProvider>
  )
}

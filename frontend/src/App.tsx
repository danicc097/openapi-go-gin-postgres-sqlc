import { BrowserRouter, Routes, Route } from 'react-router-dom'
import React from 'react'
import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
// import 'regenerator-runtime/runtime'
import { EuiProvider } from '@elastic/eui'
import { useUISlice } from 'src/slices/ui'

const Layout = React.lazy(() => import('./components/Layout/Layout'))
const LandingPage = React.lazy(() => import('./views/LandingPage/LandingPage'))

export default function App() {
  const theme = useUISlice((state) => state?.theme)

  return (
    <EuiProvider colorMode={theme}>
      <BrowserRouter basename="">
        <React.Suspense fallback={<div style={{ backgroundColor: 'azure', height: '100vh', width: '100vw' }} />}>
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

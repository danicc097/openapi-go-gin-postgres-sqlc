import React from 'react'
import App from './App'
import './index.css'
import TraceProvider from './TraceProvider'
import ReactDOM from 'react-dom/client'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from 'src/react-query'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <TraceProvider>
      <QueryClientProvider client={queryClient} /**persistOptions={{ persister }} */>
        <App />
      </QueryClientProvider>
    </TraceProvider>
  </React.StrictMode>,
)

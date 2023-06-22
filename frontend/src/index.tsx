import React from 'react'
import App from './App'
import './index.css'
import TraceProvider from './TraceProvider'
import ReactDOM from 'react-dom/client'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <TraceProvider>
      <App />
    </TraceProvider>
  </React.StrictMode>,
)

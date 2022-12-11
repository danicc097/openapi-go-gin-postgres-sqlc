import React from 'react'
import App from './App'
import './index.css'
import TraceProvider from './TraceProvider'
import ReactDOM from 'react-dom'
import './icons'

ReactDOM.render(
  <React.StrictMode>
    <TraceProvider>
      <App />
    </TraceProvider>
  </React.StrictMode>,
  document.getElementById('root'),
)

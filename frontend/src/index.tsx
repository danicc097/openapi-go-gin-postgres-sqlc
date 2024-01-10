import React from 'react'
import App from './App'
import './index.css'
import TraceProvider from './TraceProvider'
import ReactDOM from 'react-dom/client'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <TraceProvider>
      {' '}
      {/* FIXME: does not work if set after react router dom, for workaround see
        https://codesandbox.io/p/sandbox/reactour-tour-demo-using-react-router-dom-forked-yhw82c?file=%2Fcomponents%2FMain.js%3A12%2C20

        also see live demo csb: https://codesandbox.io/p/sandbox/reactour-demo-template-live-6z56m8x18k?file=%2FApp.js%3A154%2C27-154%2C45
        */}
      <App />
    </TraceProvider>
  </React.StrictMode>,
)

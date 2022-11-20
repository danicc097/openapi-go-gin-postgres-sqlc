import React from 'react'
import { Provider } from 'react-redux'
import App from './App'
import './index.css'
import configureReduxStore from './redux/store'
import TraceProvider from './TraceProvider'
import ReactDOM from 'react-dom'
import './icons'

export const store = configureReduxStore()

ReactDOM.render(
  <React.StrictMode>
    <TraceProvider>
      <Provider store={store}>
        <App />
      </Provider>
    </TraceProvider>
  </React.StrictMode>,
  document.getElementById('root'),
)

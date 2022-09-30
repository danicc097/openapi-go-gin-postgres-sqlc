import { MantineProvider } from '@mantine/core'
import React from 'react'
import ReactDOM from 'react-dom/client'
import { Provider } from 'react-redux'
import App from './App'
import './index.css'
import configureReduxStore from './redux/store'

export const store = configureReduxStore()

import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/sdk-trace-base'
import { WebTracerProvider } from '@opentelemetry/sdk-trace-web'
import { ZipkinExporter } from '@opentelemetry/exporter-zipkin'
import TraceProvider from 'src/TraceProvider'
// import { B3Propagator } from '@opentelemetry/propagator-b3'
// import { CompositePropagator, W3CTraceContextPropagator } from '@opentelemetry/core'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <TraceProvider>
      <Provider store={store}>
        <MantineProvider
          withGlobalStyles
          withNormalizeCSS
          theme={{
            colorScheme: 'dark',
            shadows: {
              md: '1px 1px 3px rgba(0, 0, 0, .25)',
              xl: '5px 5px 3px rgba(0, 0, 0, .25)',
            },
          }}
        >
          <App />
        </MantineProvider>
      </Provider>
    </TraceProvider>
  </React.StrictMode>,
)

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

const provider = new WebTracerProvider()
provider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()))
provider.addSpanProcessor(
  new SimpleSpanProcessor(
    new ZipkinExporter({
      // testing interceptor
      // getExportRequestHeaders: ()=> {
      //   return {
      //     foo: 'bar',
      //   }
      // }
    }),
  ),
)

provider.register()

export const tracer = provider.getTracer('example-tracer-web')

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
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
  </React.StrictMode>,
)

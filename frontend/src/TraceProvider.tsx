import { WebTracerProvider } from '@opentelemetry/sdk-trace-web'
import { ZoneContextManager } from '@opentelemetry/context-zone'
import { FetchInstrumentation } from '@opentelemetry/instrumentation-fetch'
import { registerInstrumentations } from '@opentelemetry/instrumentation'
import { Resource } from '@opentelemetry/resources'
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http'
import { ZipkinExporter } from '@opentelemetry/exporter-zipkin'
import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/sdk-trace-base'

const collectorOptions = {
  url: 'http://localhost:14268/api/traces',
  headers: {
    'Content-Type': 'application/json',
    'Access-Control-Allow-Headers': '*',
    // 'X-CSRF': '1',
  },
  concurrencyLimit: 10,
}

// Trace provider (Main aplication trace)
const provider = new WebTracerProvider({
  resource: new Resource({
    'service.name': 'Frontend',
  }),
})

// Exporter (opentelemetry collector hidden behind bff proxy)
const exporter = new OTLPTraceExporter(collectorOptions)

// Instrumentation configurations for frontend
const fetchInstrumentation = new FetchInstrumentation({
  ignoreUrls: ['https://some-ignored-url.com'],
})

fetchInstrumentation.setTracerProvider(provider)

provider.addSpanProcessor(new SimpleSpanProcessor(exporter))

provider.register({
  contextManager: new ZoneContextManager(),
})

// Registering instrumentations
registerInstrumentations({
  instrumentations: [new FetchInstrumentation()],
})

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

export const tracer = provider.getTracer('example-tracer-web')

export type TraceProviderProps = {
  children?: React.ReactNode
}

export default function TraceProvider({ children }: TraceProviderProps) {
  return <>{children}</>
}

import { WebTracerProvider } from '@opentelemetry/sdk-trace-web'
import { ZoneContextManager } from '@opentelemetry/context-zone'
import { FetchInstrumentation } from '@opentelemetry/instrumentation-fetch'
import { registerInstrumentations } from '@opentelemetry/instrumentation'
import { Resource } from '@opentelemetry/resources'
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http'
import { ZipkinExporter } from '@opentelemetry/exporter-zipkin'
import { BatchSpanProcessor, ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/sdk-trace-base'
import { DocumentLoadInstrumentation } from '@opentelemetry/instrumentation-document-load'
import { BatchSpanProcessorBase } from '@opentelemetry/sdk-trace-base/build/src/export/BatchSpanProcessorBase'

const collectorOptions = {
  // url: 'http://localhost:9411/api/v2/spans',
  // headers: {
  //   // 'Content-Type': 'application/json',
  //   // 'Access-Control-Allow-Headers': '*',
  //   // 'X-CSRF': '1',
  // },
  // concurrencyLimit: 10,
}

// Trace provider (Main aplication trace)
const provider = new WebTracerProvider({
  resource: new Resource({
    'service.name': 'Frontend',
  }),
})

// Exporter (opentelemetry collector hidden behind bff proxy)
// const exporter = new OTLPTraceExporter(collectorOptions)
// Instrumentation configurations for frontend
const fetchInstrumentation = new FetchInstrumentation({
  ignoreUrls: ['https://some-ignored-url.com'],
})

fetchInstrumentation.setTracerProvider(provider)

provider.addSpanProcessor(
  new BatchSpanProcessor(
    new ZipkinExporter({
      url: 'http://localhost:9411/api/v2/spans',
    }),
  ),
)
provider.addSpanProcessor(new BatchSpanProcessor(new ConsoleSpanExporter()))

provider.register()

// Registering instrumentations
registerInstrumentations({
  // , new DocumentLoadInstrumentation() not working as expected with SPA...
  instrumentations: [new FetchInstrumentation()],
})

// provider.addSpanProcessor(
//   new SimpleSpanProcessor(
//     new ZipkinExporter({
//       // testing interceptor
//       // getExportRequestHeaders: ()=> {
//       //   return {
//       //     foo: 'bar',
//       //   }
//       // }
//     }),
//   ),
// )

export const tracer = provider.getTracer('example-tracer-web')

export type TraceProviderProps = {
  children?: React.ReactNode
}

export default function TraceProvider({ children }: TraceProviderProps) {
  // tracer.startSpan('test-span').end()

  return <>{children}</>
}

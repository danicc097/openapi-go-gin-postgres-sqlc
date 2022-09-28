import { WebTracerProvider } from '@opentelemetry/sdk-trace-web'
import { ZoneContextManager } from '@opentelemetry/context-zone'
import { FetchInstrumentation } from '@opentelemetry/instrumentation-fetch'
import { XMLHttpRequestInstrumentation } from '@opentelemetry/instrumentation-xml-http-request'
import { registerInstrumentations } from '@opentelemetry/instrumentation'
import { Resource } from '@opentelemetry/resources'
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http'
import { ZipkinExporter } from '@opentelemetry/exporter-zipkin'
import { BatchSpanProcessor, ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/sdk-trace-base'
import { DocumentLoadInstrumentation } from '@opentelemetry/instrumentation-document-load'
import { BatchSpanProcessorBase } from '@opentelemetry/sdk-trace-base/build/src/export/BatchSpanProcessorBase'
import { B3Propagator } from '@opentelemetry/propagator-b3'
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions'
import { CompositePropagator, W3CBaggagePropagator, W3CTraceContextPropagator } from '@opentelemetry/core'
import { getWebAutoInstrumentations } from '@opentelemetry/auto-instrumentations-web'

const provider = new WebTracerProvider({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: 'frontend',
  }),
})

provider.addSpanProcessor(new BatchSpanProcessor(new ConsoleSpanExporter()))
provider.addSpanProcessor(
  new BatchSpanProcessor(
    new ZipkinExporter({
      url: 'http://localhost:9411/api/v2/spans',
      headers: {
        'Content-Type': 'application/json', // by default uses text/plain -> 400
      },
    }),
  ),
)
// FIXME otlp http://localhost:4318/v1/traces net::ERR_CONNECTION_REFUSED
provider.addSpanProcessor(
  new BatchSpanProcessor(
    new OTLPTraceExporter({
      url: 'http://localhost:4318/v1/traces',
      headers: {
        'Content-Type': 'application/json', // by default uses text/plain -> 400
      },
    }),
  ),
)

const contextManager = new ZoneContextManager()
provider.register({
  contextManager,
  propagator: new CompositePropagator({
    propagators: [new W3CBaggagePropagator(), new W3CTraceContextPropagator()],
  }),
})

registerInstrumentations({
  tracerProvider: provider,
  instrumentations: [
    getWebAutoInstrumentations({
      '@opentelemetry/instrumentation-fetch': {
        propagateTraceHeaderCorsUrls: /.*/,
        clearTimingResources: true,
      },
      '@opentelemetry/instrumentation-document-load': {
        enabled: false,
      },
    }),
  ],
})

export const tracer = provider.getTracer('frontend')

export type TraceProviderProps = {
  children?: React.ReactNode
}

export default function TraceProvider({ children }: TraceProviderProps) {
  // tracer.startSpan('test-span').end()

  return <>{children}</>
}

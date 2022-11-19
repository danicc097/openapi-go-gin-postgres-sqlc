import { WebTracerProvider } from '@opentelemetry/sdk-trace-web'
import { ZoneContextManager } from '@opentelemetry/context-zone'
import type { FetchCustomAttributeFunction } from '@opentelemetry/instrumentation-fetch'
import type { XHRCustomAttributeFunction } from '@opentelemetry/instrumentation-xml-http-request'
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
import opentelemetry from '@opentelemetry/api'

export const sessionID = crypto.randomUUID()

export enum AttributeKeys {
  SessionID = 'browser-session-id',
}

export function newFrontendSpan(name: string) {
  const span = tracer.startSpan(name)
  span.setAttribute(AttributeKeys.SessionID, sessionID)
  return span
}

type TraceProviderProps = {
  children?: React.ReactNode
}

// IMPORTANT: For host browser in localhost, ensure port 9411 forwarded.
export default function TraceProvider({ children }: TraceProviderProps) {
  const provider = new WebTracerProvider({
    resource: new Resource({
      [SemanticResourceAttributes.SERVICE_NAME]: 'frontend',
    }),
  })
  // OTLPTraceExporter not supported by jaeger https://github.com/jaegertracing/jaeger/issues/3479#issuecomment-1012199971
  // use zipkin
  const zipKinSpanProcessor = new BatchSpanProcessor(
    new ZipkinExporter({
      // TODO traefik label for prod since calls are made from outside
      // url: 'http://localhost:9411/api/v2/spans',
      headers: {
        'Content-Type': 'application/json',
      },
    }),
  )
  provider.addSpanProcessor(zipKinSpanProcessor)

  const contextManager = new ZoneContextManager()
  provider.register({
    contextManager,
    propagator: new CompositePropagator({
      propagators: [new W3CBaggagePropagator(), new W3CTraceContextPropagator()],
    }),
  })

  const applyCustomAttributesOnSpan: FetchCustomAttributeFunction = (span, request, result) => {
    span.setAttribute(AttributeKeys.SessionID, sessionID)
  }

  // cannot import ShouldPreventSpanCreation
  const shouldPreventSpanCreation = (eventType, element: HTMLElement, span) => {
    span.setAttribute(AttributeKeys.SessionID, sessionID)
  }

  const applyCustomAttributesOnSpanXHR: XHRCustomAttributeFunction = (span, xhr) => {
    span.setAttribute(AttributeKeys.SessionID, sessionID)
  }

  registerInstrumentations({
    tracerProvider: provider,
    instrumentations: [
      getWebAutoInstrumentations({
        '@opentelemetry/instrumentation-fetch': {
          propagateTraceHeaderCorsUrls: /.*/,
          clearTimingResources: true,
          applyCustomAttributesOnSpan,
        },
        '@opentelemetry/instrumentation-document-load': {
          enabled: false,
        },
        '@opentelemetry/instrumentation-user-interaction': {
          //  You can also use this handler to enhance created span with extra attributes.
          shouldPreventSpanCreation,
        },
        '@opentelemetry/instrumentation-xml-http-request': {
          propagateTraceHeaderCorsUrls: /.*/,
          clearTimingResources: true,
          applyCustomAttributesOnSpan: applyCustomAttributesOnSpanXHR,
        },
      }),
    ],
  })

  return <>{children}</>
}

export const tracer = opentelemetry.trace.getTracer('frontend')

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

// IMPORTANT: For host browser in localhost, ensure ports 9411 are forwarded.
// not everything will work if dynamic ports are used most likely, but behind traefik
// it should be fine
export default function TraceProvider({ children }: TraceProviderProps) {
  const provider = new WebTracerProvider({
    resource: new Resource({
      [SemanticResourceAttributes.SERVICE_NAME]: 'frontend',
    }),
  })
  const zipKinSpanProcessor = new BatchSpanProcessor(
    new ZipkinExporter({
      // url: 'http://localhost:9411/api/v2/spans',
      headers: {
        'Content-Type': 'application/json',
      },
    }),
  )
  provider.addSpanProcessor(zipKinSpanProcessor)
  // FIXME otlp http://localhost:4318/v1/traces CORS error
  // has been blocked by CORS policy: Response to preflight request doesn't pass access control check:
  // No 'Access-Control-Allow-Origin' header is present on the requested resource.
  // with zipkin we can have --collector.zipkin.allowed-origins="*" --collector.zipkin.allowed-headers="*"
  // so the ZipkinExporter works fine. Equivalent for otlp collector is...?
  provider.addSpanProcessor(
    new BatchSpanProcessor(
      new OTLPTraceExporter({
        // url: 'http://localhost:4318/v1/traces',
        // https://github.com/open-telemetry/opentelemetry-js/issues/3062
        // {} so it uses xhr instead of sendBeacon, but gives the same cors error
        headers: {},
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

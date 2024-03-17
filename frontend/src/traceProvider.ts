import { v4 as uuidv4 } from 'uuid'
import opentelemetry from '@opentelemetry/api'

// let ContextManager
// if (import.meta.env.TESTING) {
//   ContextManager = ZoneContextManagerPeerDep
// } else {
//   ContextManager = ZoneContextManager
// }

export const sessionID = uuidv4()

export enum AttributeKeys {
  SessionID = 'browser-session-id',
}

export function newFrontendSpan(name: string) {
  const span = tracer.startSpan(name)
  span.setAttribute(AttributeKeys.SessionID, sessionID)
  return span
}

export const tracer = opentelemetry.trace.getTracer('frontend')

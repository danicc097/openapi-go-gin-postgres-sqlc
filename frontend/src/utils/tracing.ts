import { tracer } from 'src/TraceProvider'

export const withSpan = <T extends Array<any>, U>(fn: (...args: T) => U, name: string) => {
  const span = tracer.startSpan(name)
  try {
    return (...args: T): U => fn(...args)
  } catch (error) {
    throw error
  } finally {
    span.end()
  }
}

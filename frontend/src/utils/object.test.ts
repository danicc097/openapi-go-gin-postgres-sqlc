import _ from 'lodash'
import { uiConfigCustomizer } from './object'
import { expect, test } from 'vitest'

const obj = {
  a: 1,
  b: 2,
  c: { b: 2 },
  d: 1,
  e: true,
  f: '123',
  g: { d: 1, e: true, f: '123' },
  h: { column: { width: 999 } },
}
const defaultObj = {
  a: [5],
  b: 3,
  c: { a: 1 },
  d: true,
  e: 'string',
  f: 123,
  g: { d: true, e: 'string', f: 123 },
  h: { column: { width: 10 } },
}

test('test merging objects', () => {
  expect(_.mergeWith(obj, defaultObj, uiConfigCustomizer)).toEqual({
    a: [5],
    b: 2,
    c: { a: 1, b: 2 },
    d: true,
    e: 'string',
    f: 123,
    g: { d: true, e: 'string', f: 123 },
    h: { column: { width: 999 } },
  })
})

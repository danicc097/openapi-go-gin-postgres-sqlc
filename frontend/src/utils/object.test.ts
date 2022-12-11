import _ from 'lodash'
import { deepMerge } from './object'
import { expect, test } from 'vitest'

const obj = {
  a: 1, // bad data
  c: { b: 2 },
  d: 1, // bad data
  e: true, // bad data
  f: '123',
  g: { d: 1, e: true, f: 123 }, // bad data except f since types don't match default's
  h: { column: { width: 999 } },
  i: ['item 1'],
}
const defaultObj = {
  a: [5],
  c: { a: 1 },
  d: true,
  e: 'string',
  f: 123,
  g: { d: true, e: 'string', f: 0 },
  h: { column: { width: 10 } },
  i: ['item 2'],
  j: ['j'],
}

test('test merging objects', () => {
  expect(_.mergeWith(obj, defaultObj, deepMerge)).toEqual({
    a: [5],
    c: { a: 1, b: 2 },
    d: true,
    e: 'string',
    f: 123,
    g: { d: true, e: 'string', f: 123 },
    h: { column: { width: 999 } },
    i: ['item 1'],
    j: ['j'],
  })
})

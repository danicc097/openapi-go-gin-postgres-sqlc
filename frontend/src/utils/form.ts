import type { Branded } from 'src/types/utils'
import { isObject } from 'src/utils/object'

export type SchemaKey = Branded<string, 'SchemaKey'>
export type FormField = Branded<string, 'FormField'>

export type ValidationErrors = Record<
  SchemaKey,
  {
    message: string
    index?: number
  }
>

/**
 * Converts react-hook-form errors to simpler internal gen formats.
 * For deeply nested errors, we just want to provide basic info in the callout for reference,
 * so we ignore intermediate indexes, if any (input component will have an error anyway).
 *
 * mode:
 *   - formField: ``item.0.nested.2``
 *   - schemaKey: ``item.nested``
 */
export function flattenRHFError({
  obj,
  prefix = '',
  ignoredKeys = [],
  mode = 'schemaKey',
  index = null,
}: {
  obj: Record<any, any>
  prefix?: string
  ignoredKeys?: string[]
  index?: number | null
  mode?: 'formField' | 'schemaKey'
}): ValidationErrors {
  return Object.keys(obj).reduce((acc: ValidationErrors, key) => {
    if (ignoredKeys.includes(key)) return acc

    let pre = prefix.length ? `${prefix}.` : ''
    if (mode == 'schemaKey') {
      pre = pre.replace(/\d+\.$/, '')
    }

    const val = obj[key]
    if (
      typeof val === 'object' &&
      !(val instanceof HTMLElement) && // inf recursion and useless
      val !== null
    ) {
      if (Array.isArray(val)) {
        for (const [idx, v] of val.entries()) {
          // nested array of objects
          if (isObject(v)) {
            Object.assign(acc, flattenRHFError({ obj: val, prefix: pre + key, ignoredKeys, mode, index: idx }))

            return acc
          }
          // rhf error array
          // extract any error (just one per schemakey in callout)
          if (v) {
            acc[pre + key] = { message: v.message, index: idx } // keep last index always
          }
        }

        return acc
      }

      // must be rhf error
      if (val.hasOwnProperty('type') && val.hasOwnProperty('ref') && val.hasOwnProperty('message')) {
        if (mode == 'schemaKey') {
          key = key.replace(/\.\d+$/, '')
        }
        acc[pre + key] = { message: val.message, index }

        return acc
      }
      Object.assign(acc, flattenRHFError({ obj: val, prefix: pre + key, ignoredKeys, mode }))
    } else {
      acc[pre + key] = val
    }

    return acc
  }, {})
}

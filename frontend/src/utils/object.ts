import _ from 'lodash'

/**
 * Fills in missing keys of any depth in `obj` with `defaultObj` keys.
 */
export const deepMerge = (obj, defaultObj: Record<string, any>) => {
  if (_.isObject(obj) && _.isObject(defaultObj)) {
    return _.mergeWith(obj, defaultObj, deepMerge)
  }

  // maybe will need to revisit to merge every item in case of nested objects or arrays in an array...
  if (_.isArray(obj) && _.isArray(defaultObj)) {
    return obj.every((item) => typeof item == typeof defaultObj[0]) ? obj : defaultObj
  }

  // assume bad data was used
  if (typeof obj != typeof defaultObj) return defaultObj

  return obj
}

export function isObject(input) {
  return input !== null && typeof input === 'object' && Object.getPrototypeOf(input).isPrototypeOf(Object)
}

/**
 * Returns an array of keys from the provided object, preserving the type information of the keys.
 */
export function keys<T>(obj: T): Array<keyof T> {
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  //@ts-ignore
  return Object.keys(obj)
}

/**
 * Returns an array of key-value pairs from the provided object, preserving the type information of the keys and   values.
 */
export function entries<T>(obj: T): Array<[keyof T, NonNullable<T[keyof T]>]> {
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  return Object.entries(obj)
}

export function hasNonEmptyValue(obj: any): boolean {
  if (typeof obj !== 'object') {
    return obj !== undefined && obj !== '' && obj !== null
  }

  for (const key in obj) {
    if (obj.hasOwnProperty(key) && hasNonEmptyValue(obj[key])) {
      return true
    }
  }

  return false
}

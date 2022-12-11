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

import _ from 'lodash'

export function uiConfigCustomizer(obj, defaultObj: Record<string, any>) {
  if (_.isObject(obj) && _.isObject(defaultObj)) {
    return _.mergeWith(obj, defaultObj, uiConfigCustomizer)
  }

  // maybe will need to revisit to merge every item in case of nested objects or arrays in an array...
  if (_.isArray(obj) && _.isArray(defaultObj)) {
    return obj.every((item) => typeof item == typeof defaultObj[0]) ? obj : defaultObj
  }

  if (typeof obj != typeof defaultObj) return defaultObj

  return obj
}

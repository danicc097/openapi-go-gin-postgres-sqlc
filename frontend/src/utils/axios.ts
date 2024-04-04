import { AxiosInstance } from 'axios'
import qs from 'qs'
/**
 * Axios utility class for adding token and Date handling to all request/responses.
 * Based on https://github.com/anymaniax/orval/issues/805
 */
export default class AxiosInterceptors {
  // hold instances of interceptor by their unique URL
  static requestInterceptors = new Map<string, number>()
  static responseInterceptors = new Map<string, number>()

  /**
   * Configures Axios request/reponse interceptors to add JWT token.
   *
   * @param {AxiosInstance} instance the Axios instance to remove interceptors
   * @param {string} token the JWT token to add to all Axios requests
   */
  static setupAxiosInstance = (instance: AxiosInstance, token: string) => {
    const appKey = instance.defaults.baseURL!

    instance.defaults.paramsSerializer = {
      serialize: (params) => {
        return qs.stringify(params, { arrayFormat: 'indices', encode: false })
      }, // kin-openapi and oapi-codegen runtime
    }

    const tokenRequestInterceptor = instance.interceptors.request.use(
      (config) => {
        if (token) {
          const headers = config.headers || {}
          headers.Authorization = `Bearer ${token}`
        }
        return config
      },
      (error) => {
        return Promise.reject(error)
      },
    )

    AxiosInterceptors.requestInterceptors.set(appKey, tokenRequestInterceptor)

    // if using useDates in orval TS types are Date
    const dateResponseInterceptor = instance.interceptors.response.use((originalResponse) => {
      handleDates(originalResponse.data)
      return originalResponse
    })

    AxiosInterceptors.responseInterceptors.set(appKey, dateResponseInterceptor)
  }

  /**
   * Cleanup Axios on sign out of application.
   *
   * @param {AxiosInstance} instance the Axios instance to remove interceptors
   */
  static teardownAxiosInstance = (instance: AxiosInstance) => {
    const appKey = instance.defaults.baseURL!
    instance.interceptors.request.clear()
    instance.interceptors.response.clear()
  }
}

const isoDateFormat = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d*)?(?:[-+]\d{2}:?\d{2}|Z)?$/

function isIsoDateString(value: any): boolean {
  return value && typeof value === 'string' && isoDateFormat.test(value)
}

function handleDates(body: any) {
  if (body === null || body === undefined || typeof body !== 'object') return body

  for (const key of Object.keys(body)) {
    const value = body[key]
    if (isIsoDateString(value)) {
      body[key] = new Date(value) // default JS conversion
      // body[key] = parseISO(value); // date-fns conversion
      // body[key] = luxon.DateTime.fromISO(value); // Luxon conversion
      // body[key] = moment(value).toDate(); // Moment.js conversion
    } else if (typeof value === 'object') {
      handleDates(value)
    }
  }
}

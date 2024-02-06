import CONFIG from 'config.json'
import Axios, { AxiosError, type AxiosRequestConfig } from 'axios'
import { apiPath } from 'utils/path'

export const AXIOS_INSTANCE = Axios.create({ baseURL: '<BACKEND URL>' })

export class ApiError extends Error {
  response?: AxiosError['response']
  constructor(message: string, response?: AxiosError['response']) {
    super(message)
    this.name = 'ApiError'
    this.response = response
  }
}

export const customInstance = <T>(config: AxiosRequestConfig, options?: AxiosRequestConfig): Promise<T> => {
  const source = Axios.CancelToken.source()
  const promise = AXIOS_INSTANCE({
    ...config,
    ...options,
    cancelToken: source.token,
    baseURL: apiPath(null),
  })
    .then(({ data }) => data)
    .catch((error: AxiosError) => {
      throw new ApiError(error.message, error.response as any)
    })

  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  promise.cancel = () => {
    source.cancel('Query was cancelled')
  }

  return promise
}

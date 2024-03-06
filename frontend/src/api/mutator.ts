import Axios, { AxiosError, type AxiosRequestConfig } from 'axios'
import { CONFIG } from 'src/config'
import { HTTPError } from 'src/gen/model'
import { apiPath } from 'src/services/apiPaths'

export const AXIOS_INSTANCE = Axios.create()

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
    baseURL: apiPath(),
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

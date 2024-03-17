import Axios, { AxiosError, type AxiosRequestConfig } from 'axios'
import { CONFIG } from 'src/config'
import { HTTPError } from 'src/gen/model'
import { apiPath } from 'src/services/apiPaths'

export const AXIOS_INSTANCE = Axios.create()

export class ApiError extends Error {
  response?: AxiosError<HTTPError>['response']
  constructor(message: string, response?: AxiosError<HTTPError>['response']) {
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

// must be called ErrorType for orval to replace. other options:
// https://github.com/anymaniax/orval/blob/b63ffe671e5eeb4e06730add9cb1b947b59798f5/docs/src/pages/guides/custom-axios.md?plain=1#L50
export type ErrorType<Error> = AxiosError<HTTPError>

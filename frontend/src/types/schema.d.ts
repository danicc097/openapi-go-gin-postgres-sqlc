/* eslint-disable @typescript-eslint/ban-ts-comment */
/* eslint-disable */
// @ts-nocheck
export type schemas = components['schemas']
/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  '/ping': {
    get: operations['Ping']
  }
  '/openapi.yaml': {
    get: operations['OpenapiYamlGet']
  }
  '/admin/ping': {
    get: operations['AdminPing']
  }
  '/user/me': {
    get: operations['GetCurrentUser']
  }
  [key: `/user/${string}`]: {
    put: operations['UpdateUser']
    delete: operations['DeleteUser']
  }
}

export interface components {
  schemas: {
    /** HTTPValidationError */
    HTTPValidationError: {
      /** Detail */
      detail?: components['schemas']['ValidationError'][]
    }
    /**
     * a User
     * @description represents User data to update
     * @example {
     *   "role": "manager",
     *   "first_name": "Jane",
     *   "last_name": "Doe"
     * }
     */
    UpdateUserRequest: {
      role?: components['schemas']['Role']
      first_name?: string
      last_name?: string
    }
    /**
     * Scope
     * @enum {string}
     */
    Scope: 'users:read' | 'users:write' | 'scopes:write' | 'settings:write' | 'work-item:review'
    /**
     * Role
     * @description User role.
     * @enum {string}
     */
    Role: 'guest' | 'user' | 'advanced user' | 'manager' | 'admin' | 'superadmin'
    /**
     * Organization
     * @description Organization a user belongs to.
     */
    Organization: string
    /**
     * a User
     * @description represents a user
     */
    GetCurrentUserRes: {
      /** Format: int64 */
      user_id?: number
      username?: string
      first_name?: string
      last_name?: string
      email?: string
      password?: string
      phone?: string
      role?: components['schemas']['Role']
      /** @description are organizations a user belongs to */
      orgs?: components['schemas']['Organization'][]
    }
    /** ValidationError */
    ValidationError: {
      /** Location */
      loc: string[]
      /** Message */
      msg: string
      /** Error Type */
      type: string
    }
  }
}

export interface operations {
  Ping: {
    responses: {
      /** OK */
      200: {
        content: {
          'text/plain': string
        }
      }
      /** Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError']
        }
      }
    }
  }
  OpenapiYamlGet: {
    responses: {
      /** OpenAPI YAML file. */
      200: {
        content: {
          'text/yaml': string
        }
      }
    }
  }
  AdminPing: {
    responses: {
      /** OK */
      200: {
        content: {
          'text/plain': string
        }
      }
      /** Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError']
        }
      }
    }
  }
  GetCurrentUser: {
    responses: {
      /** successful operation */
      200: {
        content: {
          'application/json': components['schemas']['GetCurrentUserRes']
        }
      }
    }
  }
  UpdateUser: {
    parameters: {
      path: {
        /** user_id that needs to be updated */
        id: string
      }
    }
    responses: {
      /** User not found */
      404: unknown
    }
    /** Updated user object */
    requestBody: {
      content: {
        'application/json': components['schemas']['UpdateUserRequest']
      }
    }
  }
  DeleteUser: {
    parameters: {
      path: {
        /** user_id that needs to be deleted */
        id: string
      }
    }
    responses: {
      /** User not found */
      404: unknown
    }
  }
}

export interface external {}

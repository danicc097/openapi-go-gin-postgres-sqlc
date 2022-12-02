/* eslint-disable @typescript-eslint/ban-ts-comment */
/* eslint-disable */
// @ts-nocheck
export type schemas = components['schemas']

/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


/** Type helpers */
type Without<T, U> = { [P in Exclude<keyof T, keyof U>]?: never };
type XOR<T, U> = (T | U) extends object ? (Without<T, U> & U) | (Without<U, T> & T) : T | U;
type OneOf<T extends any[]> = T extends [infer Only] ? Only : T extends [infer A, infer B, ...infer Rest] ? OneOf<[XOR<A, B>, ...Rest]> : never;

export interface paths {
  "/events": {
    get: operations["Events"];
  };
  "/ping": {
    /** Ping pongs */
    get: operations["Ping"];
  };
  "/openapi.yaml": {
    /** Returns this very OpenAPI spec. */
    get: operations["OpenapiYamlGet"];
  };
  "/admin/ping": {
    /** Ping pongs */
    get: operations["AdminPing"];
  };
  "/user/me": {
    /** returns the logged in user */
    get: operations["GetCurrentUser"];
  };
  "/user/{id}/authorization": {
    /** updates user role and scopes by id */
    patch: operations["UpdateUserAuthorization"];
  };
  "/user/{id}": {
    /** deletes the user by id */
    delete: operations["DeleteUser"];
    /** updates the user by id */
    patch: operations["UpdateUser"];
  };
}

export interface components {
  schemas: {
    /** HTTPValidationError */
    HTTPValidationError: {
      /** Detail */
      detail?: (components["schemas"]["ValidationError"])[];
    };
    /** @enum {string} */
    Scope: "test-scope" | "users:read" | "users:write" | "scopes:write" | "team-settings:write" | "project-settings:write" | "work-item:review";
    Scopes: (components["schemas"]["Scope"])[];
    /** @enum {string} */
    Role: "guest" | "user" | "advancedUser" | "manager" | "admin" | "superAdmin";
    /**
     * Notification type 
     * @description User notification type. 
     * @enum {string}
     */
    NotificationType: "personal" | "global";
    /**
     * WorkItem role 
     * @description Role in work item for a member. 
     * @enum {string}
     */
    WorkItemRole: "preparer" | "reviewer";
    /**
     * @description represents User data to update 
     * @example {
     *   "first_name": "Jane",
     *   "last_name": "Doe"
     * }
     */
    UpdateUserRequest: {
      /** @description originally from auth server but updatable */
      first_name?: string;
      /** @description originally from auth server but updatable */
      last_name?: string;
    };
    /**
     * @description represents User authorization data to update 
     * @example {
     *   "role": "manager",
     *   "scopes": [
     *     "test-scope"
     *   ]
     * }
     */
    UpdateUserAuthRequest: {
      role?: components["schemas"]["Role"];
      scopes?: components["schemas"]["Scopes"];
    };
    UserPublic: {
      apiKeyID?: number | null;
      /** Format: date-time */
      createdAt?: string;
      /** Format: date-time */
      deletedAt?: string | null;
      email?: string;
      firstName?: string | null;
      fullName?: string | null;
      lastName?: string | null;
      teams?: (components["schemas"]["TeamPublic"])[] | null;
      timeEntries?: (components["schemas"]["TimeEntryPublic"])[] | null;
      userID?: components["schemas"]["UuidUUID"];
      username?: string;
      workItems?: (components["schemas"]["WorkItemPublic"])[] | null;
    };
    /** ValidationError */
    ValidationError: {
      /** Location */
      loc: (string)[];
      /** Message */
      msg: string;
      /** Error Type */
      type: string;
    };
    PgtypeJSONB: Record<string, never>;
    UuidUUID: string;
    TaskPublic: {
      /** Format: date-time */
      createdAt?: string;
      /** Format: date-time */
      deletedAt?: string | null;
      finished?: boolean | null;
      metadata?: components["schemas"]["PgtypeJSONB"];
      taskID?: number;
      taskType?: components["schemas"]["TaskTypePublic"];
      taskTypeID?: number;
      title?: string;
      /** Format: date-time */
      updatedAt?: string;
      workItemID?: number;
    };
    TaskTypePublic: {
      color?: string;
      description?: string;
      name?: string;
      taskTypeID?: number;
      teamID?: number;
    } | null;
    TeamPublic: {
      /** Format: date-time */
      createdAt: string;
      description: string;
      metadata: components["schemas"]["PgtypeJSONB"];
      name: string;
      projectID: number;
      teamID: number;
      /** Format: date-time */
      updatedAt: string;
    };
    TimeEntryPublic: {
      activityID?: number;
      comment?: string;
      durationMinutes?: number | null;
      /** Format: date-time */
      start?: string;
      teamID?: number | null;
      timeEntryID?: number;
      userID?: components["schemas"]["UuidUUID"];
      workItemID?: number | null;
    };
    WorkItemCommentPublic: {
      /** Format: date-time */
      createdAt?: string;
      message?: string;
      /** Format: date-time */
      updatedAt?: string;
      userID?: components["schemas"]["UuidUUID"];
      workItemCommentID?: number;
      workItemID?: number;
    };
    WorkItemPublic: {
      closed?: boolean;
      /** Format: date-time */
      createdAt?: string;
      /** Format: date-time */
      deletedAt?: string | null;
      kanbanStepID?: number;
      metadata?: components["schemas"]["PgtypeJSONB"];
      tasks?: (components["schemas"]["TaskPublic"])[] | null;
      teamID?: number;
      timeEntries?: (components["schemas"]["TimeEntryPublic"])[] | null;
      title?: string;
      /** Format: date-time */
      updatedAt?: string;
      users?: (components["schemas"]["UserPublic"])[] | null;
      workItemComments?: (components["schemas"]["WorkItemCommentPublic"])[] | null;
      workItemID?: number;
      workItemTypeID?: number;
    };
    ModelsRole: string;
    UserResponse: {
      apiKey?: components["schemas"]["UserAPIKeyPublic"];
      /** Format: date-time */
      createdAt: string;
      /** Format: date-time */
      deletedAt: string | null;
      email: string;
      firstName: string | null;
      fullName: string | null;
      hasGlobalNotifications: boolean;
      hasPersonalNotifications: boolean;
      lastName: string | null;
      role: components["schemas"]["Role"];
      scopes: components["schemas"]["Scopes"];
      teams?: (components["schemas"]["TeamPublic"])[] | null;
      userID: components["schemas"]["UuidUUID"];
      username: string;
    };
    ModelsScope: string;
    UserAPIKeyPublic: {
      apiKey: string;
      /** Format: date-time */
      expiresOn: string;
      userID: components["schemas"]["UuidUUID"];
    } | null;
  };
  responses: never;
  parameters: {
    /**
     * @description user_id that needs to be updated 
     * @example 123e4567-e89b-12d3-a456-426614174000
     */
    UserID: string;
  };
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type external = Record<string, never>;

export interface operations {

  Events: {
    responses: {
      /** @description events */
      200: {
        content: {
          "text/event-stream": ({
              id?: string;
              data?: Record<string, never>;
            })[];
        };
      };
    };
  };
  Ping: {
    /** Ping pongs */
    responses: {
      /** @description OK */
      200: {
        content: {
          "text/plain": string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          "application/json": components["schemas"]["HTTPValidationError"];
        };
      };
    };
  };
  OpenapiYamlGet: {
    /** Returns this very OpenAPI spec. */
    responses: {
      /** @description OpenAPI YAML file. */
      200: {
        content: {
          "text/yaml": string;
        };
      };
    };
  };
  AdminPing: {
    /** Ping pongs */
    responses: {
      /** @description OK */
      200: {
        content: {
          "text/plain": string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          "application/json": components["schemas"]["HTTPValidationError"];
        };
      };
    };
  };
  GetCurrentUser: {
    /** returns the logged in user */
    responses: {
      /** @description ok */
      200: {
        content: {
          "application/json": components["schemas"]["UserResponse"];
        };
      };
    };
  };
  UpdateUserAuthorization: {
    /** updates user role and scopes by id */
    /** @description Updated user object */
    requestBody: {
      content: {
        "application/json": components["schemas"]["UpdateUserAuthRequest"];
      };
    };
    responses: {
      /** @description ok */
      200: {
        content: {
          "application/json": components["schemas"]["UserResponse"];
        };
      };
    };
  };
  DeleteUser: {
    /** deletes the user by id */
    responses: {
      /** @description User not found */
      404: never;
    };
  };
  UpdateUser: {
    /** updates the user by id */
    /** @description Updated user object */
    requestBody: {
      content: {
        "application/json": components["schemas"]["UpdateUserRequest"];
      };
    };
    responses: {
      /** @description ok */
      200: {
        content: {
          "application/json": components["schemas"]["UserPublic"];
        };
      };
    };
  };
}

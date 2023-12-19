import CONFIG_JSON from '../config.json'
import { JSONSchemaType } from 'ajv'
import jsonSchema from 'src/client-validator/gen/dereferenced-schema.json'
import JSON_SCHEMA_JSON from 'src/client-validator/gen/dereferenced-schema.json'
import OPERATION_AUTH_JSON from '../operationAuth.gen.json'
import type { Role, Scopes } from 'src/gen/model'
import { operations } from 'src/types/schema'
import ROLES_JSON from '../roles.json'
import SCOPES_JSON from '../scopes.json'
import type { Scope } from 'src/gen/model'

export const CONFIG = CONFIG_JSON

export const SCOPES = SCOPES_JSON as unknown as {
  [key in Scope]: typeof SCOPES_JSON[keyof typeof SCOPES_JSON]
}

export const ROLES = ROLES_JSON as unknown as {
  [key in Role]: typeof ROLES_JSON[keyof typeof ROLES_JSON]
}

export type OperationAuth = {
  scopes: Scopes
  role: Role
  requiresAuthentication: boolean
}

export const OPERATION_AUTH = OPERATION_AUTH_JSON as unknown as {
  [key in keyof operations]: OperationAuth
}

export type SchemaDefinitions = keyof typeof jsonSchema.definitions

export const JSON_SCHEMA = JSON_SCHEMA_JSON as unknown as {
  definitions: {
    [key in SchemaDefinitions]: JSONSchemaType<true>
  }
}

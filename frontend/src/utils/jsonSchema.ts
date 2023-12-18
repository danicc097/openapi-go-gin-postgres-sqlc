import { JSONSchemaType } from 'ajv'
import { JSONSchema } from 'json-schema-to-ts'
import type { Primitive, RecursiveKeyOf } from 'src/types/utils'

// use JSONSchemaType<true> from ajv instead
// export interface JsonSchemaField {
//   type?: Type | Type[]
//   format?: Format
//   items?: JsonSchemaField
//   properties?: { [key: string]: JsonSchemaField }
//   required?: string[]
// }

export type SchemaField = {
  type: Type | Format
  required: boolean
  isArray: boolean
}

export type Format = 'date-time' | 'date'
export type Type = 'boolean' | 'integer' | 'object' | 'string' | 'array' | 'null' | 'number'

export function parseSchemaFields(schema: JSONSchemaType<true>): Record<Primitive, SchemaField> {
  const schemaFields = {}

  function traverseSchema(obj: JSONSchemaType<true>, path: string[] = [], parent: JSONSchemaType<true> | null = null) {
    const isArrayOfObj = extractType(obj) === 'object' && !!obj.type?.includes('array') && obj.items?.properties
    const isObj = obj.properties && extractType(obj) === 'object'

    if (isArrayOfObj) {
      extract(obj?.items?.properties)
    } else if (isObj) {
      extract(obj.properties)
    }

    function extract(properties: JSONSchemaType<true>['properties']) {
      for (const key in properties) {
        const newPath = [...path, key]
        const property = properties[key]
        if (!property) {
          continue
        }

        schemaFields[newPath.join('.')] = {
          type: extractType(property),
          required: extractIsRequired(obj, parent, key),
          isArray: !!property.type?.includes('array'),
        }
        traverseSchema(property, newPath, property)
      }
    }
  }

  traverseSchema(schema)

  return schemaFields as Record<Primitive, SchemaField>
}

function extractIsRequired(obj: JSONSchemaType<true>, parent: JSONSchemaType<true> | null, key: string): boolean {
  if (!parent) {
    return (
      !!obj.required?.includes(key) &&
      !obj.properties?.[key]?.type?.includes('null') &&
      !obj?.items?.type?.includes('null')
    )
  }

  if (parent.items) {
    return extractIsRequired(obj, parent.items, key)
  }

  return (
    !!parent.required?.includes(key) &&
    !obj.properties?.[key]?.type?.includes('null') &&
    !obj.items?.type?.includes('null') &&
    (obj.type?.includes('array') && obj.items?.properties?.[key]?.type?.includes('array')
      ? !obj?.type?.includes('null')
      : true)
  )
}

function extractType(obj: JSONSchemaType<true>): Type | Format {
  const type = _type(obj.type)
  if (type === 'array') {
    if (obj?.items?.type === 'object') {
      return 'object'
    } else {
      return _type(obj?.items?.type)
    }
  }

  return obj.format ?? type

  function _type(x: JSONSchemaType<true>['type']) {
    return (Array.isArray(x) ? x.filter((t) => t !== 'null')[0] : x) as Type
  }
}

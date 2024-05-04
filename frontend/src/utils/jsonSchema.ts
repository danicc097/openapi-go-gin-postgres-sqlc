import { JSONSchemaType } from 'ajv'
import { JSONSchema4, JSONSchema4Array, JSONSchema4Object, JSONSchema4Type } from 'json-schema'
import { JSONSchema } from 'json-schema-to-ts'
import type { Primitive, RecursiveKeyOf } from 'src/types/utils'

export type SchemaField = {
  type: Type | Format
  required: boolean
  isArray: boolean
}

export type Format = 'date-time' | 'date'
export type Type = 'boolean' | 'integer' | 'object' | 'string' | 'array' | 'null' | 'number'

export function parseSchemaFields(schema: JSONSchema4): Record<Primitive, SchemaField> {
  if (schema.oneOf || schema.allOf || schema.anyOf) {
    throw Error('Can not parse oneOf, anyOf, allOf schemas. Please provide the discriminated schema instead')
  }
  const schemaFields = {}

  function traverseSchema(obj: JSONSchema4, path: string[] = [], parent: JSONSchema4 | null = null) {
    const isArrayOfObj =
      extractType(obj) === 'object' && !!obj.type?.includes('array') && (obj.items as JSONSchema4)?.properties
    const isObj = obj.properties && extractType(obj) === 'object'

    if (isArrayOfObj) {
      extract((obj.items as JSONSchema4)?.properties)
    } else if (isObj) {
      extract(obj.properties)
    }

    function extract(properties: JSONSchema4Object | JSONSchema4Array | undefined) {
      if (!properties) return

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

function extractIsRequired(obj: JSONSchema4, parent: JSONSchema4 | null, key: string): boolean {
  if (!parent) {
    return (
      (obj.required === true || !!(obj.required as string[])?.includes(key)) &&
      !obj.properties?.[key]?.type?.includes('null') &&
      !(obj?.items as JSONSchema4)?.type?.includes('null')
    )
  }

  if (parent.items) {
    return extractIsRequired(obj, parent.items, key)
  }

  return (
    (parent.required === true || !!(parent.required as string[])?.includes(key)) &&
    !obj.properties?.[key]?.type?.includes('null') &&
    !(obj.items as JSONSchema4)?.type?.includes('null') &&
    (obj.type?.includes('array') && (obj.items as JSONSchema4)?.properties?.[key]?.type?.includes('array')
      ? !obj?.type?.includes('null')
      : true)
  )
}

export type JsonSchemaType = Type | Type[]

function _type(x: JsonSchemaType) {
  return (Array.isArray(x) ? x.filter((t) => t !== 'null')[0] : x) as Type
}

function extractType(obj: JSONSchema4Object): Type | Format {
  const type = _type(obj.type as JsonSchemaType)
  if (type === 'array') {
    const items = obj.items as JSONSchema4Object
    if (items?.type === 'object') {
      return 'object'
    } else {
      return _type(items?.type as JsonSchemaType)
    }
  }

  return (obj.format ?? type) as Type | Format
}

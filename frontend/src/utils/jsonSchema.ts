import type { Primitive, RecursiveKeyOf } from 'src/types/utils'

export interface JsonSchemaField {
  type?: Type | Type[]
  format?: Format
  items?: JsonSchemaField
  properties?: { [key: string]: JsonSchemaField }
  required?: string[]
}

export type SchemaField = {
  type: Type | Format
  required: boolean
  isArray: boolean
}

export type Format = 'date-time' | 'date'
export type Type = 'boolean' | 'integer' | 'object' | 'string' | 'array' | 'null' | 'number'

export function parseSchemaFields(schema: JsonSchemaField): Record<Primitive, SchemaField> {
  const schemaFields = {}

  function traverseSchema(obj: JsonSchemaField, path: string[] = [], parent: JsonSchemaField | null = null) {
    const isArrayOfObj = extractType(obj) === 'object' && !!obj.type?.includes('array') && obj.items?.properties
    const isObj = obj.properties && extractType(obj) === 'object'

    if (isArrayOfObj) {
      extract(obj?.items?.properties)
    } else if (isObj) {
      extract(obj.properties)
    }

    function extract(properties: JsonSchemaField['properties']) {
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

function extractIsRequired(obj: JsonSchemaField, parent: JsonSchemaField | null, key: string): boolean {
  if (!parent) {
    return !!obj.required?.includes(key)
  }

  if (parent.items) {
    return extractIsRequired(obj, parent.items, key)
  }

  return !!parent.required?.includes(key)
}

function extractType(obj: JsonSchemaField): Type | Format {
  const type = _type(obj.type)
  if (type === 'array') {
    if (obj?.items?.type === 'object') {
      return 'object'
    } else {
      return _type(obj?.items?.type)
    }
  }

  return obj.format ?? type

  function _type(x: JsonSchemaField['type']) {
    return (Array.isArray(x) ? x.filter((t) => t !== 'null')[0] : x) as Type
  }
}

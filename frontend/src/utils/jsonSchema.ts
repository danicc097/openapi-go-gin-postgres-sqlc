interface JsonSchemaField {
  type?: string | string[]
  format?: string
  items?: JsonSchemaField
  properties?: { [key: string]: JsonSchemaField }
  required?: string[]
}

type FieldTypes = {
  [fieldName: string]: {
    type: string
    required: boolean
    isArray: boolean
  }
}

export function extractFieldTypes(schema: JsonSchemaField): FieldTypes {
  const fieldTypes: FieldTypes = {}

  function traverseSchema(obj: JsonSchemaField, path: string[] = [], parent: JsonSchemaField | null = null) {
    if (obj.properties) {
      for (const key in obj.properties) {
        const newPath = [...path, key]
        const property = obj.properties[key]
        fieldTypes[newPath.join('.')] = {
          type: extractType(property),
          required: extractIsRequired(obj, parent, key),
          isArray: !!property.type?.includes('array'),
        }
        traverseSchema(property, newPath, property)
      }
    } else if (extractType(obj) === 'arrayOfObject' && obj.items?.properties) {
      for (const key in obj.items.properties) {
        const newPath = [...path, key]
        const property = obj.items.properties[key]
        fieldTypes[newPath.join('.')] = {
          type: extractType(property),
          required: extractIsRequired(obj, parent.items, key),
          isArray: !!property.type?.includes('array'),
        }
        traverseSchema(property, newPath, property)
      }
    }
  }

  traverseSchema(schema)

  return fieldTypes
}

function extractIsRequired(obj: JsonSchemaField, parent: JsonSchemaField | null, key: string): boolean {
  if (!parent) {
    return obj.required?.includes(key)
  }

  return Array.isArray(parent?.required) ? parent.required.includes(key) : false
}

function extractType(obj: JsonSchemaField): string {
  const type = _type(obj.type)
  if (type === 'array') {
    if (obj?.items?.type === 'object') {
      return 'arrayOfObject'
    } else {
      return _type(obj?.items?.type)
    }
  }

  return obj.format ?? type

  function _type(x: string | string[]) {
    return Array.isArray(x) ? x.filter((t) => t !== 'null')[0] : x
  }
}
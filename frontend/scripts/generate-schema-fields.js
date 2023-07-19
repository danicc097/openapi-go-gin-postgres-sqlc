import JSON_SCHEMA from '../src/client-validator/gen/dereferenced-schema.json' assert { type: 'json' }

// TODO: return object indexed by projectName and save to workItemSchemaFields. will be used by front and back.
// instead of a global config mess we can have per-user config in local storage for visualization.
// simple if (!projects.includes(<project>)) guard sanity check will do it
const projects = JSON_SCHEMA.definitions.Project.enum
const schemaFields = parseSchemaFields(JSON_SCHEMA.definitions.RestDemoWorkItemsResponse)
console.log(schemaFields)

function parseSchemaFields(schema) {
  const schemaFields = {}

  function traverseSchema(obj, path = []) {
    const isArrayOfObj = extractType(obj) === 'object' && !!obj.type?.includes('array') && obj.items?.properties
    const isObj = obj.properties && extractType(obj) === 'object'

    if (isArrayOfObj) {
      extract(obj?.items?.properties)
    } else if (isObj) {
      extract(obj.properties)
    }

    function extract(properties) {
      for (const key in properties) {
        const newPath = [...path, key]
        const property = properties[key]
        if (!property) {
          continue
        }

        schemaFields[newPath.join('.')] = null
        traverseSchema(property, newPath)
      }
    }
  }

  traverseSchema(schema)

  return Object.keys(schemaFields)
}

function extractType(obj) {
  const type = _type(obj.type)
  if (type === 'array') {
    if (obj?.items?.type === 'object') {
      return 'object'
    } else {
      return _type(obj?.items?.type)
    }
  }

  return obj.format ?? type

  function _type(x) {
    return Array.isArray(x) ? x.filter((t) => t !== 'null')[0] : x
  }
}

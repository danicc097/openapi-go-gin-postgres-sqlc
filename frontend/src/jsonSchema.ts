import { JSONSchemaType } from 'ajv'
import jsonSchema from 'src/client-validator/gen/dereferenced-schema.json'
import JSON_SCHEMA from 'src/client-validator/gen/dereferenced-schema.json'

type SchemaDefinitions = keyof typeof jsonSchema.definitions

export default JSON_SCHEMA as unknown as {
  definitions: {
    [key in SchemaDefinitions]: JSONSchemaType<true>
  }
}

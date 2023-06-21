import type { RestDemoWorkItemCreateRequest } from 'src/gen/model'
import type { RecursiveKeyOf } from 'src/types/utils'
import { parseSchemaFields, type JsonSchemaField, type SchemaField } from 'src/utils/jsonSchema'
import { describe, expect, test } from 'vitest'

describe('parseSchemaFields', () => {
  test('should extract field types correctly from a JSON schema', () => {
    const schema = {
      properties: {
        base: {
          properties: {
            closed: {
              format: 'date-time',
              type: ['string', 'null'],
            },
            description: {
              type: 'string',
            },
            kanbanStepID: {
              type: 'integer',
            },
            metadata: {
              items: {
                minimum: 0,
                type: 'integer',
              },
              type: ['array', 'null'],
            },
            targetDate: {
              format: 'date-time',
              type: 'string',
            },
            teamID: {
              type: 'integer',
            },
            title: {
              type: 'string',
            },
            workItemTypeID: {
              type: 'integer',
            },
          },
          required: [
            'title',
            'description',
            'workItemTypeID',
            'metadata',
            'teamID',
            'kanbanStepID',
            'closed',
            'targetDate',
          ],
          type: 'object',
          $schema: 'http://json-schema.org/draft-04/schema#',
        },
        demoProject: {
          properties: {
            lastMessageAt: {
              format: 'date-time',
              type: 'string',
            },
            line: {
              type: 'string',
            },
            ref: {
              pattern: '^[0-9]{8}$',
              type: 'string',
            },
            reopened: {
              type: 'boolean',
            },
            workItemID: {
              type: 'integer',
            },
          },
          required: ['workItemID', 'ref', 'line', 'lastMessageAt', 'reopened'],
          type: 'object',
          $schema: 'http://json-schema.org/draft-04/schema#',
        },
        members: {
          items: {
            properties: {
              role: {
                title: 'WorkItem role',
                type: 'string',
                'x-generated': '-',
                enum: ['preparer', 'reviewer'],
                description: "represents a database 'work_item_role'",
                $schema: 'http://json-schema.org/draft-04/schema#',
              },
              userID: {
                type: 'string',
                $schema: 'http://json-schema.org/draft-04/schema#',
              },
            },
            required: ['userID', 'role'],
            type: 'object',
            $schema: 'http://json-schema.org/draft-04/schema#',
          },
          type: ['array', 'null'],
        },
        tagIDs: {
          items: {
            type: 'integer',
          },
          type: ['array', 'null'],
        },
      },
      required: ['demoProject', 'base', 'tagIDs', 'members'],
      type: 'object',
      'x-postgen-struct': 'RestDemoWorkItemCreateRequest',
      $schema: 'http://json-schema.org/draft-04/schema#',
    } as JsonSchemaField

    const schemaFields = parseSchemaFields(schema)

    const a: RecursiveKeyOf<RestDemoWorkItemCreateRequest> = 'members' // OK
    const b: RecursiveKeyOf<RestDemoWorkItemCreateRequest> = 'members.role' // OK
    const c1: RecursiveKeyOf<RestDemoWorkItemCreateRequest> = '.tagIDs' // FIXME: should giveError:
    const c: RecursiveKeyOf<RestDemoWorkItemCreateRequest> = 'members.role.role' // should giveError: Type '"members.role.role"' is not assignable to type '"members.role"'
    const d: RecursiveKeyOf<RestDemoWorkItemCreateRequest> = 'base.metadata' // ok
    const e: RecursiveKeyOf<RestDemoWorkItemCreateRequest> = 'demoProject.reopened' // ok

    const wantFields: Record<RecursiveKeyOf<RestDemoWorkItemCreateRequest>, SchemaField> = {
      base: { isArray: false, required: true, type: 'object' },
      'base.closed': { type: 'date-time', required: true, isArray: false },
      'base.description': { type: 'string', required: true, isArray: false },
      'base.kanbanStepID': { type: 'integer', required: true, isArray: false },
      'base.metadata': { type: 'integer', required: true, isArray: true },
      'base.targetDate': { type: 'date-time', required: true, isArray: false },
      'base.teamID': { type: 'integer', required: true, isArray: false },
      'base.title': { type: 'string', required: true, isArray: false },
      'base.workItemTypeID': { type: 'integer', required: true, isArray: false },
      demoProject: { isArray: false, required: true, type: 'object' },
      'demoProject.lastMessageAt': { type: 'date-time', required: true, isArray: false },
      'demoProject.line': { type: 'string', required: true, isArray: false },
      'demoProject.ref': { type: 'string', required: true, isArray: false },
      'demoProject.reopened': { type: 'boolean', required: true, isArray: false },
      'demoProject.workItemID': { type: 'integer', required: true, isArray: false },
      tagIDs: { type: 'integer', required: true, isArray: true },
      members: { type: 'object', required: true, isArray: true },
      'members.role': { type: 'string', required: true, isArray: false },
      'members.userID': { type: 'string', required: true, isArray: false },
    }
    /**

    form generator will use these keys. to generate multiple forms when is array we just check
    if parent (split by . and keep up to len-2) isArray (members) or the child itself isArray (tagIDs)

    it doesnt seem to be easy to get typed keys for these when arrays are involved.
    */

    expect(schemaFields).toEqual(wantFields)
  })
})

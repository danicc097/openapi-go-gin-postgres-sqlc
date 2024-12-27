import type { DeepPartial, GetKeys, RecursiveKeyOf, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import DynamicForm from 'src/utils/formGeneration'
import { parseSchemaFields, type SchemaField } from 'src/utils/jsonSchema'
import { describe, expect, test, vitest } from 'vitest'
import {
  getByTestId,
  render,
  screen,
  renderHook,
  fireEvent,
  act,
  getByText,
  waitFor,
  getQueriesForElement,
} from '@testing-library/react'
import '@testing-library/jest-dom'
import dayjs from 'dayjs'
import { entries, keys } from 'src/utils/object'
import { FormProvider, useForm } from 'react-hook-form'
import { ajvResolver } from '@hookform/resolvers/ajv'
import { fullFormats } from 'ajv-formats/dist/formats'
import { Group, Avatar, Space, Flex, MantineProvider } from '@mantine/core'
import { nameInitials } from 'src/utils/strings'
import { JSONSchemaType } from 'ajv'
import { selectOptionsBuilder } from 'src/utils/formGeneration.context'
import { JSONSchema } from 'json-schema-to-ts'
import userEvent from '@testing-library/user-event'
import { VirtuosoMockContext } from 'react-virtuoso'
import UserComboboxOption from 'src/components/Combobox/UserComboboxOption'
import { JSONSchema4 } from 'json-schema'

export const refPattern = '^[0-9]{8}$'

export const schema: JSONSchema4 = {
  properties: {
    base: {
      properties: {
        closed: {
          format: 'date-time',
          type: ['object', 'null'],
        },
        description: {
          type: 'string',
        },
        kanbanStepID: {
          type: 'integer',
        },
        metadata: {
          type: 'object',
        },
        targetDate: {
          format: 'date',
          type: 'object',
        },
        teamID: {
          type: 'integer',
        },
        // purposely name them nested items to ensure correct recursion
        items: {
          items: {
            properties: {
              items: {
                items: {
                  type: 'string',
                },
                type: ['array', 'null'],
              },
              userId: {
                items: {
                  type: 'string',
                },
                type: ['array', 'null'],
              },
              name: {
                type: 'string',
                $schema: 'http://json-schema.org/draft-04/schema#',
              },
            },
            required: ['userId', 'items', 'name'],
            type: 'object',
            $schema: 'http://json-schema.org/draft-04/schema#',
          },
          type: ['array', 'null'],
        },
        workItemTypeID: {
          type: 'integer',
        },
      },
      required: [
        'items',
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
          type: 'object',
        },
        line: {
          type: 'string',
        },
        ref: {
          pattern: refPattern,
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
    tagIDsMultiselect: {
      items: {
        type: 'integer',
      },
      type: ['array', 'null'],
    },
  },
  required: ['demoProject', 'base', 'tagIDsMultiselect', 'members'],
  type: 'object',
  'x-gen-struct': 'RestDemoWorkItemCreateRequest',
}

export const schemaFields: Record<GetKeys<TestTypes.DemoWorkItemCreateRequest>, SchemaField> = {
  base: { isArray: false, required: true, type: 'object' },
  'base.closed': { type: 'date-time', required: false, isArray: false },
  'base.description': { type: 'string', required: true, isArray: false },
  'base.kanbanStepID': { type: 'integer', required: true, isArray: false },
  'base.metadata': { type: 'object', required: true, isArray: false },
  'base.targetDate': { type: 'date', required: true, isArray: false },
  'base.teamID': { type: 'integer', required: true, isArray: false },
  'base.items': { type: 'object', required: false, isArray: true },
  'base.items.name': { type: 'string', required: true, isArray: false },
  'base.items.userId': { type: 'string', required: false, isArray: true },
  'base.items.items': { type: 'string', required: false, isArray: true },
  'base.workItemTypeID': { type: 'integer', required: true, isArray: false },
  demoProject: { isArray: false, required: true, type: 'object' },
  'demoProject.lastMessageAt': { type: 'date-time', required: true, isArray: false },
  'demoProject.line': { type: 'string', required: true, isArray: false },
  'demoProject.ref': { type: 'string', required: true, isArray: false },
  'demoProject.reopened': { type: 'boolean', required: true, isArray: false },
  'demoProject.workItemID': { type: 'integer', required: true, isArray: false },
  members: { type: 'object', required: false, isArray: true },
  'members.role': { type: 'string', required: true, isArray: false },
  'members.userID': { type: 'string', required: true, isArray: false },
  tagIDs: { type: 'integer', required: false, isArray: true },
  tagIDsMultiselect: { type: 'integer', required: false, isArray: true },
}

describe('schemas', () => {
  test('should extract field types correctly from a JSON schema', () => {
    expect(parseSchemaFields(schema)).toStrictEqual(schemaFields)
  })
})

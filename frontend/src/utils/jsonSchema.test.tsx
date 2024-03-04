import type { DeepPartial, GetKeys, RecursiveKeyOf, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import DynamicForm from 'src/utils/formGeneration'
import { JsonSchemaField, parseSchemaFields, type SchemaField } from 'src/utils/jsonSchema'
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

const tags = [...Array(10)].map((x, i) => {
  return {
    name: `tag #${i}`,
    color: `#${i % 10}34236`,
    workItemTagID: i,
    projectID: 1,
    description: 'description',
  }
})

const refPattern = '^[0-9]{8}$'

const schema = {
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
} as JsonSchemaField

const formInitialValues = {
  base: {
    items: [
      { items: ['0001', '0002'], userId: [], name: 'item-1' },
      { items: ['0011', '0012'], userId: [], name: 'item-2' },
    ],
    closed: dayjs('2023-03-24T20:42:00.000Z').toDate(),
    // targetDate: dayjs('2023-02-22').toDate(),
    description: 'some text',
    kanbanStepID: 1,
    teamID: 1,
    metadata: {},
    workItemTypeID: 1,
    targetDate: dayjs('2024-03-24T20:42:00.000Z').toDate(),
  },
  demoProject: {
    lastMessageAt: dayjs('2023-03-24T20:42:00.000Z').toDate(),
    line: '3e3e2',
    ref: '124321', // should fail pattern validation
    workItemID: 1,
    reopened: true,
  },
  tagIDs: ['aaa', 1, 2],
  tagIDsMultiselect: [0, 1, 2],
  members: [
    // with defaultValue of "member.role": {role: 'preparer'} it will fill null or undefined form values.
    // since userid exists and it's an initial value, it will show custom select card to work around https://github.com/mantinedev/mantine/issues/980
    // therefore its element input id does not exist
    { userID: 'a446259c-1083-4212-98fe-bd080c41e7d7' },
    // userid does not exist in selectOptions users -> will show empty input directly instead
    // TODO: should show warning on top of form
    { role: 'reviewer', userID: 'b446259c-1083-4212-98fe-bd080c41e7d7' },
  ],
} as TestTypes.DemoWorkItemCreateRequest

const schemaFields: Record<GetKeys<TestTypes.DemoWorkItemCreateRequest>, SchemaField> = {
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

// ok
// test('toHaveFormValues', () => {
//   render(
//     <form data-testid="login-form">
//       <input type="text" name="username" value="jane.doe" />
//       <input type="password" name="password" value="12345678" />
//       <input type="checkbox" name="rememberMe" checked />
//       <button type="submit">Sign in</button>
//     </form>,
//   )

//   const form = screen.getByTestId('login-form') as HTMLFormElement
//   expect(form).toHaveFormValues({
//     username: 'jane.doe',
//   })
// })

test('should extract field types correctly from a JSON schema', () => {
  expect(parseSchemaFields(schema)).toEqual(schemaFields)
})

describe('form generation', () => {
  test('should render form fields and buttons', async () => {
    const { result: form, rerender } = renderHook(() =>
      useForm<TestTypes.DemoWorkItemCreateRequest>({
        resolver: ajvResolver(schema as any, {
          strict: false,
          formats: fullFormats,
        }),
        mode: 'onChange',
        defaultValues: formInitialValues ?? {},
        // shouldUnregister: true, // defaultValues will not be merged against submission result.
      }),
    )

    const formName = 'demoWorkItemCreateForm'

    const mockSubmit = vitest.fn()
    const mockSubmitWithErrors = vitest.fn()
    const { container, baseElement } = render(
      <VirtuosoMockContext.Provider value={{ viewportHeight: 300, itemHeight: 100 }}>
        <MantineProvider>
          <FormProvider {...form.current}>
            <DynamicForm<TestTypes.DemoWorkItemCreateRequest, 'base.metadata'>
              onSubmit={(e) => {
                e.preventDefault()
                form.current.handleSubmit(
                  // needs to be called
                  (data) => {
                    console.log({ data })
                    mockSubmit(data)
                  },
                  (errors) => {
                    console.log({ errors })
                    mockSubmitWithErrors(errors)
                  },
                )(e)
              }}
              formName={formName}
              schemaFields={schemaFields}
              options={{
                renderOrderPriority: ['tagIDs', 'members'],
                labels: {
                  base: 'base', // just title via renderTitle
                  'base.closed': 'closed',
                  'base.description': 'description',
                  // 'base.metadata': 'metadata', // ignored -> not a key
                  'base.kanbanStepID': 'kanbanStepID',
                  'base.targetDate': 'targetDate',
                  'demoProject.reopened': 'reopened',
                  'base.teamID': 'teamID',
                  'base.items': 'items',
                  'base.items.name': 'name',
                  'base.items.items': 'items',
                  'base.items.userId': 'user',
                  'base.workItemTypeID': 'workItemTypeID',
                  demoProject: null, // won't render title
                  'demoProject.lastMessageAt': 'lastMessageAt',
                  'demoProject.line': 'line',
                  'demoProject.ref': 'ref',
                  'demoProject.workItemID': 'workItemID',
                  members: 'members',
                  'members.role': 'role',
                  'members.userID': 'User',
                  tagIDs: 'tagIDs',
                  tagIDsMultiselect: 'tagIDsMultiselect',
                },
                defaultValues: {
                  'demoProject.line': '43121234', // should be ignored since it's set
                  'members.role': 'preparer',
                },
                selectOptions: {
                  'members.userID': selectOptionsBuilder({
                    type: 'select',
                    values: [...Array(5)].map((x, i) => {
                      return {
                        username: `user${i}`,
                        email: `user${i}@mail.com`,
                        userID: `a446259c-1083-4212-98fe-bd080c41e7d${i}`,
                      }
                    }),
                    optionTransformer(el) {
                      return (
                        <Group align="center">
                          <div style={{ display: 'flex', alignItems: 'center' }}>
                            <Avatar size={35} radius="xl" data-test-id="header-profile-avatar" alt={el?.username}>
                              {nameInitials(el?.email || '')}
                            </Avatar>
                            <Space p={5} />
                          </div>

                          <div style={{ marginLeft: 'auto' }}>{el?.email}</div>
                        </Group>
                      )
                    },
                    ariaLabelTransformer(el) {
                      return el.email
                    },
                    formValueTransformer(el) {
                      return el.userID
                    },
                    pillTransformer(el) {
                      return <>el.email</>
                    },
                  }),
                  tagIDsMultiselect: selectOptionsBuilder({
                    type: 'multiselect',
                    values: tags,
                    optionTransformer(el) {
                      return (
                        <Group align="center">
                          <Flex align={'center'}></Flex>
                          <div style={{ marginLeft: 'auto' }}>{el?.name}</div>
                        </Group>
                      )
                    },
                    ariaLabelTransformer(el) {
                      return el.name
                    },
                    formValueTransformer(el) {
                      return el.workItemTagID
                    },
                    pillTransformer(el) {
                      return <>{el.name} label</>
                    },
                  }),
                },
              }}
            />
          </FormProvider>
        </MantineProvider>
      </VirtuosoMockContext.Provider>,
    )

    const dataTestIds = [
      'demoWorkItemCreateForm',
      'demoWorkItemCreateForm-tagIDs-title',
      'demoWorkItemCreateForm-tagIDs-add-button',
      'demoWorkItemCreateForm-tagIDs-0',
      'demoWorkItemCreateForm-tagIDs-remove-button-0',
      'demoWorkItemCreateForm-tagIDs-1',
      'demoWorkItemCreateForm-tagIDs-remove-button-1',
      'demoWorkItemCreateForm-tagIDs-2',
      'demoWorkItemCreateForm-tagIDs-remove-button-2',
      'demoWorkItemCreateForm-tagIDsMultiselect', // multiselects dont have titles - using vanilla
      'demoWorkItemCreateForm-members-title',
      'demoWorkItemCreateForm-members-add-button',
      'demoWorkItemCreateForm-members-remove-button-0',
      'demoWorkItemCreateForm-members.0.role',
      'demoWorkItemCreateForm-members.0.userID',
      'demoWorkItemCreateForm-members-remove-button-1',
      'demoWorkItemCreateForm-members.1.role',
      'demoWorkItemCreateForm-members.1.userID',
      'demoWorkItemCreateForm-base-title',
      'demoWorkItemCreateForm-base.closed',
      'demoWorkItemCreateForm-base.description',
      'demoWorkItemCreateForm-base.kanbanStepID',
      'demoWorkItemCreateForm-base.targetDate',
      'demoWorkItemCreateForm-base.teamID',
      'demoWorkItemCreateForm-base.items-title',
      'demoWorkItemCreateForm-base.items-add-button',
      'demoWorkItemCreateForm-base.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.0.name',
      'demoWorkItemCreateForm-base.items.0.userId-title',
      'demoWorkItemCreateForm-base.items.0.userId-add-button',
      'demoWorkItemCreateForm-base.items.0.items-title',
      'demoWorkItemCreateForm-base.items.0.items-add-button',
      'demoWorkItemCreateForm-base.items.0.items-0',
      'demoWorkItemCreateForm-base.items.0.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.0.items-1',
      'demoWorkItemCreateForm-base.items.0.items-remove-button-1',
      'demoWorkItemCreateForm-base.items-remove-button-1',
      'demoWorkItemCreateForm-base.items.1.name',
      'demoWorkItemCreateForm-base.items.1.userId-title',
      'demoWorkItemCreateForm-base.items.1.userId-add-button',
      'demoWorkItemCreateForm-base.items.1.items-title',
      'demoWorkItemCreateForm-base.items.1.items-add-button',
      'demoWorkItemCreateForm-base.items.1.items-0',
      'demoWorkItemCreateForm-base.items.1.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.1.items-1',
      'demoWorkItemCreateForm-base.items.1.items-remove-button-1',
      'demoWorkItemCreateForm-base.workItemTypeID',
      'demoWorkItemCreateForm-demoProject.lastMessageAt',
      'demoWorkItemCreateForm-demoProject.line',
      'demoWorkItemCreateForm-demoProject.ref',
      'demoWorkItemCreateForm-demoProject.reopened',
      'demoWorkItemCreateForm-demoProject.workItemID',
    ]

    const actualIds = [...document.querySelectorAll('[data-testid^="demoWorkItemCreateForm"]')].map((e) =>
      e.getAttribute('data-testid'),
    )
    // console.log({ actualIds })
    expect(actualIds.sort()).toEqual(dataTestIds.sort())

    // test should submit with default values if none changed

    // FIXME: dont check state, its not updated and theres no sensible workaround:
    // https://stackoverflow.com/questions/61813319/check-state-of-a-component-using-react-testing-library.
    // maybe call submit with mock onsubmit that sets global var and check
    // that return value is what we expect.
    // https://react-hook-form.com/advanced-usage#TestingForm
    // for better testing see : https://claritydev.net/blog/testing-react-hook-form-with-react-testing-library
    const checkbox = screen.getByTestId('demoWorkItemCreateForm-demoProject.reopened')
    expect(checkbox).toBeChecked()
    checkbox.click()
    expect(checkbox).not.toBeChecked()
    const formElement = screen.getByTestId(formName)
    const submitButton = screen.getByRole('button', { name: /Submit/ })
    await act(() => {
      submitButton.click()
    })

    expect(screen.getAllByRole('alert')).toHaveLength(3) // incl box
    expect(mockSubmitWithErrors).toBeCalled()
    expect(mockSubmitWithErrors.mock.calls[0][0]).toMatchObject({
      demoProject: { ref: { message: `must match pattern "${refPattern}"` } },
      tagIDs: [{ message: 'must be integer' }],
    })
    // TODO: fix errors in ref and tagids and then
    // compare mock data with expected (requires fixing combobox options in rtl first)

    // document.body includes options, etc. for debugging
    // expect(document.body).toMatchInlineSnapshot()

    // console.log({ options: screen.getAllByRole('option', { hidden: true }).map((opt) => opt.getAttribute('aria-label')) })

    await waitFor(async () => {
      const comboboxInput = screen.getByTestId('demoWorkItemCreateForm-members.0.userID') // FIXME: its just a text input! userID is the combobox
      expect(comboboxInput).toBeInTheDocument()
      await userEvent.click(comboboxInput, { pointerState: await userEvent.pointer({ target: comboboxInput }) }) // jsdom not displaying combobox

      console.log({ options: screen.getAllByRole('option').map((opt) => opt.getAttribute('aria-label')) })
      console.log({
        allOptions: screen.getAllByRole('option', { hidden: true }).map((opt) => opt.getAttribute('aria-label')),
      })

      await userEvent.click(screen.getByRole('option', { name: 'user5@mail.com', hidden: true })) // TODO: should include formKey as discriminator somewhere, else we have duplicates on multiple items.
    })
    // screen.debug(document)

    let firstMember = screen.getByTestId('demoWorkItemCreateForm-members.0.role') // FIXME: its just a text input! 0.member is the combobox
    expect(firstMember).toHaveDisplayValue('preparer')

    const option = await screen.findByRole('option', { name: /reviewer/i, hidden: true }) // [role="option"][value="preparer"]
    await act(() => {
      option.click()
    })
    firstMember = screen.getByTestId('demoWorkItemCreateForm-members.0.role')
    expect(firstMember).toHaveDisplayValue('reviewer')
    const secondMember = screen.getByTestId('demoWorkItemCreateForm-members.1.role')
    expect(secondMember).toHaveDisplayValue('reviewer')
  })
})

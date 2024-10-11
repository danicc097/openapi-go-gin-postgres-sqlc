import type { DeepPartial, GetKeys, RecursiveKeyOf, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import DynamicForm from 'src/utils/formGeneration'
import { parseSchemaFields, type SchemaField } from 'src/utils/jsonSchema'
import { describe, expect, test, vitest } from 'vitest'
import {
  getByTestId,
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
import { UserResponse } from 'src/gen/model'
import { schema, refPattern, schemaFields } from 'src/utils/jsonSchema.test'
import { setup } from 'src/test-utils/render'

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
  'demoWorkItemCreateForm-tagIDsMultiselect', // multiselects dont have titles - using vanilla input label
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
  'demoWorkItemCreateForm-search--members.0.userID',
  'demoWorkItemCreateForm-search--members.1.userID',
  'demoWorkItemCreateForm-search--tagIDsMultiselect',
  'demoWorkItemCreateForm-tagIDsMultiselect-remove--0',
  'demoWorkItemCreateForm-tagIDsMultiselect-remove--aaaa',
  'demoWorkItemCreateForm-tagIDsMultiselect-remove--2',
]

const tags = [...Array(10)].map((x, i) => {
  return {
    name: `tag #${i}`,
    color: `#${i}34236`,
    workItemTagID: i,
    projectID: 1,
    description: 'description',
  }
})

const badId = 'aaaa'

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
  tagIDs: [1, 3, 2],
  tagIDsMultiselect: [0, badId, 2],
  members: [
    // with defaultValue of "member.role": {role: 'preparer'} it will fill null or undefined form values.
    // since userid exists and it's an initial value, it will show custom select card to work around https://github.com/mantinedev/mantine/issues/980
    // therefore its element input id does not exist
    { userID: 'a446259c-1083-4212-98fe-bd080c41e7d7' },
    // userid does not exist in selectOptions users -> will show empty input directly instead
    // and a warning on top of form. TODO: mantine input element should have a warning mode, exactly as current errors but just in yellow.
    // it should just be a rightSection alert icon with popover saying the same as the
    // callout box. no need for css and setting classlists on helptext, borders, etc.

    { role: 'reviewer', userID: 'b446259c-1083-4212-98fe-bd080c41e7d7' },
  ],
} as TestTypes.DemoWorkItemCreateRequest

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

    setup(
      <FormProvider {...form.current}>
        <DynamicForm<TestTypes.DemoWorkItemCreateRequest, 'base.metadata' | 'demoProject'>
          onSubmit={(e) => {
            e.preventDefault()
            form.current.handleSubmit(
              // needs to be called
              (data) => {
                mockSubmit(data)
              },
              (errors) => {
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
              // demoProject: null, // won't render title
              'demoProject.lastMessageAt': 'lastMessageAt',
              'demoProject.line': 'line',
              'demoProject.ref': 'ref',
              'demoProject.workItemID': 'workItemID',
              members: 'members',
              'members.role': 'role',
              'members.userID': 'User',
              tagIDs: 'tagIDs',
              tagIDsMultiselect: 'Tag IDs',
            },
            defaultValues: {
              'demoProject.line': '43121234', // should be ignored since it's set
              'members.role': 'preparer',
            },
            selectOptions: {
              'members.userID': selectOptionsBuilder({
                type: 'select',
                values: [...Array(10)].map((x, i) => {
                  return {
                    username: `user${i}`,
                    email: `user${i}@mail.com`,
                    userID: `a446259c-1083-4212-98fe-bd080c41e7d${i}`,
                    role: 'user',
                  } as UserResponse
                }),
                optionTransformer(el) {
                  return <UserComboboxOption user={el} />
                },
                ariaLabelTransformer(el) {
                  return el.email
                },
                formValueTransformer(el) {
                  return el.userID
                },
                pillTransformer(el) {
                  return <>{el.email}</>
                },
                searchValueTransformer(el) {
                  return `${el.email} ${el.username}`
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
      </FormProvider>,
    )

    const submitButton = await screen.findByRole('button', { name: /Submit/i })
    expect(submitButton).toBeInTheDocument()

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
    await act(() => checkbox.click())
    expect(checkbox).not.toBeChecked()
    await act(() => checkbox.click())
    expect(checkbox).toBeChecked()

    expect(screen.getAllByRole('alert')).toHaveLength(1) // warning callout before submit
    await act(() => submitButton.click())

    expect(screen.getAllByRole('alert')).toHaveLength(1) // error callout only after submit
    expect(mockSubmitWithErrors).toBeCalledTimes(1)
    expect(mockSubmit).toBeCalledTimes(0)
    expect(mockSubmitWithErrors.mock.calls[0]).toMatchObject([
      {
        demoProject: { ref: { message: `must match pattern "${refPattern}"` } },
        tagIDsMultiselect: [undefined, { message: 'must be integer' }],
      },
    ])
    // TODO: fix errors in ref and tagids and then
    // compare mock data via mockSubmit with expected (requires fixing combobox options in rtl first)
    const refInput = screen.getByTestId('demoWorkItemCreateForm-demoProject.ref')
    await userEvent.clear(refInput)
    await userEvent.type(refInput, '99998888')

    await waitFor(async () => {
      const tagsSearchInput = screen.getByTestId('demoWorkItemCreateForm-search--tagIDsMultiselect')
      await userEvent.type(tagsSearchInput, 'tag #4')

      await userEvent.click(screen.getByRole('option', { name: 'tag #4', hidden: false })) // no need for discriminator if there's only a visible opt
      await userEvent.clear(tagsSearchInput)
    })
    const tagsSearchInput = screen.getByTestId('demoWorkItemCreateForm-search--tagIDsMultiselect')
    await userEvent.click(tagsSearchInput, { pointerState: await userEvent.pointer({ target: tagsSearchInput }) }) // no need for discriminator if there's only a visible opt

    expect(screen.getByRole('option', { name: 'tag #4', hidden: true }).getAttribute('aria-selected')).toBe('true') // here would need discriminator since its hidden after click

    const firstUserIDInput = screen.getByTestId('demoWorkItemCreateForm-members.0.userID') as HTMLInputElement
    expect(firstUserIDInput).toHaveAccessibleName('User')

    const email = 'user9@mail.com'
    const firstUserIDSearchInput = screen.getByTestId('demoWorkItemCreateForm-search--members.0.userID')
    await waitFor(async () => {
      await userEvent.click(firstUserIDInput, { pointerState: await userEvent.pointer({ target: firstUserIDInput }) })
      await userEvent.type(firstUserIDSearchInput, email)
      await userEvent.click(screen.getByRole('option', { name: email, hidden: false })) // no need for discriminator if there's only a visible opt
    })
    await userEvent.click(firstUserIDInput, { pointerState: await userEvent.pointer({ target: firstUserIDInput }) }) // show opt with search filter
    expect(screen.getByRole('option', { name: email, hidden: false }).getAttribute('aria-selected')).toBe('true')

    const badTagCloseButton = screen.getByTestId('demoWorkItemCreateForm-tagIDsMultiselect-remove--aaaa')
    await waitFor(async () => {
      await userEvent.click(badTagCloseButton, { pointerState: await userEvent.pointer({ target: badTagCloseButton }) })
    })

    // test remove entry in array of objects
    const closeButton = screen.getByTestId('demoWorkItemCreateForm-members-remove-button-1')
    await waitFor(async () => {
      await userEvent.click(closeButton, { pointerState: await userEvent.pointer({ target: closeButton }) })
    })

    /**
     * test final form values
     *  */
    await act(() => submitButton.click())

    expect(screen.queryAllByRole('alert')).toHaveLength(0)
    expect(mockSubmit).toBeCalledTimes(1)

    const newFormValues = formInitialValues
    newFormValues.tagIDsMultiselect = [0, 2, 4]
    newFormValues.members![0]!.role = 'preparer' // nested defaultValues if empty
    newFormValues.members![0]!.userID = 'a446259c-1083-4212-98fe-bd080c41e7d9' // user 9
    newFormValues.members?.splice(1, 1) // removed
    newFormValues.demoProject.ref = '99998888'

    expect(mockSubmit.mock.calls[0]).toStrictEqual([newFormValues])
  })
})

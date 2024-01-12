import _, { capitalize, concat, random } from 'lodash'
import React, { Fragment, forwardRef, memo, useEffect, useReducer, useState } from 'react'
import type { Scope, Scopes, UpdateUserAuthRequest, User } from 'src/gen/model'
import { getContrastYIQ, roleColor } from 'src/utils/colors'
import { joinWithAnd } from 'src/utils/format'

import type { Role } from 'src/client-validator/gen/models'
import PageTemplate from 'src/components/PageTemplate'
import type { ValidationErrors } from 'src/client-validator/validate'
import { updateUserAuthorization, useUpdateUserAuthorization } from 'src/gen/user/user'
import { validateField } from 'src/utils/validation'
import { UpdateUserAuthRequestDecoder } from 'src/client-validator/gen/decoders'
import { newFrontendSpan } from 'src/TraceProvider'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import { getGetCurrentUserMock } from 'src/gen/user/user.msw'
import type { PathType, RecursiveKeyOf, RequiredKeys } from 'src/types/utils'
import {
  Avatar,
  Badge,
  Button,
  Flex,
  Space,
  Text,
  Title,
  Select,
  Group,
  Modal,
  Checkbox,
  Code,
  Card,
  Box,
  type DefaultMantineColor,
  Grid,
  Tooltip,
  Divider,
  type ComboboxItem,
  Combobox,
  useCombobox,
  InputBase,
  Input,
} from '@mantine/core'
import { CodeHighlight } from '@mantine/code-highlight'
import { notifications } from '@mantine/notifications'
import { IconCheck, IconCircle } from '@tabler/icons'
import RoleBadge from 'src/components/Badges/RoleBadge'
import { entries, keys } from 'src/utils/object'
import { css } from '@emotion/css'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import ErrorCallout, { useCalloutErrors } from 'src/components/Callout/ErrorCallout'
import { ApiError } from 'src/api/mutator'
import { AxiosError } from 'axios'
import { isAuthorized } from 'src/services/authorization'
import { asConst } from 'json-schema-to-ts'
import type { components, schemas } from 'src/types/schema'
import { FormProvider, useForm, useFormContext, useWatch } from 'react-hook-form'
import { nameInitials } from 'src/utils/strings'
import type { AppError } from 'src/types/ui'
import classes from './UserPermissionsPage.module.css'
import UserComboboxOption from 'src/components/Combobox/UserComboboxOption'
import { useFormSlice } from 'src/slices/form'
import { JSON_SCHEMA, ROLES, SCOPES } from 'src/config'
import InfiniteLoader from 'src/components/Loading/InfiniteLoader'

type RequiredUserAuthUpdateKeys = RequiredKeys<UpdateUserAuthRequest>

const REQUIRED_USER_AUTH_UPDATE_KEYS: Record<RequiredUserAuthUpdateKeys, boolean> = {}

interface SelectUserItemProps extends React.ComponentPropsWithoutRef<'div'> {
  user: User
}

interface SelectRoleItemProps extends React.ComponentPropsWithoutRef<'div'> {
  label: string
  value: User['role']
}

function scopeColor(scopeName?: string): DefaultMantineColor {
  switch (scopeName) {
    case 'read':
      return 'green'
    case 'write':
    case 'edit':
      return 'orange'
    case 'delete':
      return 'red'
    default:
      return 'blue'
  }
}

const SelectRoleItem = ({ value }: SelectRoleItemProps) => {
  return (
    <Combobox.Option value={value}>
      <RoleBadge role={value} />
    </Combobox.Option>
  )
}

export default function UserPermissionsPage() {
  const [selectedUser, setSelectedUser] = useState<User | null>(null)
  const [userOptions, setUserOptions] = useState<Array<SelectUserItemProps> | null>(null)
  const { user } = useAuthenticatedUser()

  const [allUsers] = useState(
    [...Array(20)].map((x, i) => {
      return getGetCurrentUserMock()
    }),
  )

  const roleOptions = entries(ROLES)
    .filter(([role, v]) => isAuthorized({ user, requiredRole: role }))
    .map(([role, v]) => ({
      label: role,
      value: role,
    }))

  const scopeEditPanels: Record<string, Partial<typeof SCOPES>> = entries(SCOPES).reduce((acc, [key, value]) => {
    const [group, scope] = key.split(':', 2) as [string, string]
    if (!acc[group]) {
      acc[group] = {}
    }
    acc[group][key] = value
    return acc
  }, {})

  useEffect(() => {
    if (userOptions === null) {
      setUserOptions(
        allUsers
          ? allUsers.map((user) => ({
              label: user.email,
              value: user.email,
              user,
            }))
          : null,
      )
    } else {
      setUserOptions(userOptions)
    }
  }, [allUsers, userOptions])

  const formName = 'user-permissions-form'

  const { extractCalloutErrors, setCalloutErrors, calloutErrors, extractCalloutTitle } = useCalloutErrors(formName)

  // const { mutateAsync: updateUserAuthorization } = useUpdateUserAuthorization()

  const form = useForm<UpdateUserAuthRequest>({
    defaultValues: {},
  })

  const submitRoleUpdate = async () => {
    const span = newFrontendSpan('submitRoleUpdate')
    try {
      if (!selectedUser) return
      const updateUserAuthRequest = UpdateUserAuthRequestDecoder.decode(form.getValues())
      const payload = await updateUserAuthorization(selectedUser.userID, updateUserAuthRequest)
      console.log('fulfilled', payload)
      notifications.show({
        id: ToastId.FormSubmit,
        title: 'Submitted',
        color: 'primary',
        icon: <IconCheck size="1.2rem" />,
        autoClose: 15000,
        message: 'Submitted',
      })
      setCalloutErrors([])
    } catch (error) {
      console.error(error)
      if (error.validationErrors) {
        setCalloutErrors(error.validationErrors)
        console.log('error')
        return
      }
      setCalloutErrors([error])
    }
    span?.end()
  }

  const handleError = (errors: typeof form.formState.errors) => {
    if (errors) {
      console.log('some errors found')
      console.log(errors)

      // TODO validate everything and show ALL validation errors
      // (we dont want to show very long error messages in each form
      // field, just that the field has an error,
      // so all validation errors are aggregated with full description in a callout)
      try {
        UpdateUserAuthRequestDecoder.decode(form.getValues())
        setCalloutErrors([])
      } catch (error) {
        if (error.validationErrors) {
          setCalloutErrors(error.validationErrors)
          console.error(error)
          return
        }
        setCalloutErrors([error])
      }
    }
  }

  const onRoleSelectableChange = (role) => {
    console.log(role)
    form.setValue('role', role)
  }

  const onEmailSelectableChange = (email) => {
    const user = allUsers.find((user) => user.email === email)
    if (!user) return
    console.log(user)
    setSelectedUser(user)
    form.setValue('role', user.role)
    form.setValue('scopes', user.scopes)
  }

  const [isModalVisible, setIsModalVisible] = useState(false)
  const closeModal = () => setIsModalVisible(false)
  const showModal = () => setIsModalVisible(true)

  const onRoleUpdateSubmit = async () => {
    showModal()
  }

  const registerProps = form.register('role')

  useWatch({ name: 'role', control: form.control })

  const selectedOption = userOptions?.find((option) => {
    return option.user.email === selectedUser?.email
  })

  const [search, setSearch] = useState('')

  const combobox = useCombobox({
    onDropdownClose: () => {
      combobox.resetSelectedOption()
      combobox.focusTarget()
      setSearch('')
    },

    onDropdownOpen: () => {
      combobox.focusSearchInput()
    },
  })
  const comboboxOptions =
    userOptions
      ?.filter((item: any) => JSON.stringify(item.value).toLowerCase().includes(search.toLowerCase().trim()))
      .map((option) => {
        const value = String(option.user.email)

        return (
          <Combobox.Option value={value} key={value}>
            <UserComboboxOption user={option.user} key={JSON.stringify(option.user)} />
          </Combobox.Option>
        )
      }) || []

  const element = (
    <FormProvider {...form}>
      <ErrorCallout title={extractCalloutTitle()} errors={concat(extractCalloutErrors())} />
      <Space pt={12} />
      <Title size={12}>
        <Text>Form</Text>
      </Title>
      <FormData />
      <Space pt={12} />
      <form onSubmit={form.handleSubmit(onRoleUpdateSubmit, handleError)}>
        <Flex direction="column">
          {/* TODO: in v7: https://mantine.dev/combobox/?e=SelectOptionComponent */}
          <Combobox
            store={combobox}
            withinPortal={true}
            position="bottom-start"
            withArrow
            onOptionSubmit={async (value) => {
              const option = userOptions?.find((option) => String(option.user.email) === value)
              console.log({ onChangeOption: option })
              if (!option) return
              onEmailSelectableChange(value)
              combobox.closeDropdown()
            }}
          >
            <Combobox.Target withAriaAttributes={false}>
              <InputBase
                className={classes.select}
                component="button"
                type="button"
                pointer
                rightSection={<Combobox.Chevron />}
                onClick={() => combobox.toggleDropdown()}
                rightSectionPointerEvents="none"
                multiline
              >
                {selectedUser ? (
                  <UserComboboxOption user={selectedUser} key={JSON.stringify(selectedUser.email)} />
                ) : (
                  <Input.Placeholder>{`Pick user`}</Input.Placeholder>
                )}
              </InputBase>
            </Combobox.Target>

            <Combobox.Dropdown>
              <Combobox.Search
                miw={'100%'}
                value={search}
                onChange={(event) => setSearch(event.currentTarget.value)}
                placeholder={`Search user`}
              />
              <Combobox.Options
                mah={200} // scrollable
                style={{ overflowY: 'auto' }}
              >
                {comboboxOptions.length > 0 ? comboboxOptions : <Combobox.Empty>Nothing found</Combobox.Empty>}
              </Combobox.Options>
            </Combobox.Dropdown>
          </Combobox>
          {/* <Select
            label="Select user to update"
            value={UserComboboxOption}
            data-test-subj="updateUserAuthForm__selectable"
            searchable
            filter={({ options, search }) => {
              const splittedSearch = search.toLowerCase().trim().split(' ')
              return (options as ComboboxItem[]).filter((option) => {
                const words = option.label.toLowerCase().trim().split(' ')
                return splittedSearch.every((searchWord) => words.some((word) => word.includes(searchWord)))
              })
            }}
            data={userOptions ?? []}
            onChange={onEmailSelectableChange}
          /> */}
          {/* TODO: {renderSuperSelect<UpdateUserAuthRequest, 'role'>({
            formKey: 'role',
            form,
            options: roleOptions,
            requiredFormKeys: REQUIRED_USER_AUTH_UPDATE_KEYS,
            onSuperSelectChange: onRoleSuperSelectChange,
          })} */}
        </Flex>
        <Space pt={12} />
        {selectedUser?.email && (
          <>
            <Divider m={8} />
            <Select
              label={
                <Title size={15} mt={4} mb={4}>
                  Update role
                </Title>
              }
              disabled={!isAuthorized({ user, requiredRole: selectedUser.role })}
              // itemComponent={SelectRoleItem} // TODO: COMBOBOX
              data-test-subj="updateUserAuthForm__selectable_Role"
              defaultValue={selectedUser.role}
              data={roleOptions ?? []}
              {...registerProps}
              onChange={(value) => registerProps.onChange({ target: { name: 'role', value } })}
            />
            <Space pt={12} />
            <Title size={15} mt={4} mb={4}>
              Update scopes
            </Title>
            <Card shadow="md" padding="lg" radius="md" withBorder>
              {entries(scopeEditPanels).map(([group, scopes]) => (
                <CheckboxPanel
                  user={user}
                  userSelection={selectedUser}
                  key={group}
                  title={group.replace(/-/g, ' ').replace(/^\w{1}/g, (c) => c.toUpperCase())}
                  scopes={scopes}
                />
              ))}
            </Card>
            <Space pt={24} />
            <Button disabled={selectedUser === null} data-test-subj="updateUserAuthForm__submit" onClick={showModal}>
              Update authorization settings
            </Button>
          </>
        )}
      </form>
      <Modal
        opened={isModalVisible}
        title={
          <Text fw={'bold'} size={'md'}>
            Update auth information
          </Text>
        }
        onClose={closeModal}
        data-test-subj="updateUserAuthForm__confirmModal"
      >
        <>
          {`You're about to update auth information for `}
          <strong>{selectedUser?.email}</strong>.<p>Are you sure you want to do this?</p>
          <Group style={{ justifyContent: 'flex-end' }}>
            <Button variant="subtle" color="orange" onClick={closeModal}>
              Cancel
            </Button>
            <Button
              onClick={async () => {
                await submitRoleUpdate()
                closeModal()
              }}
            >
              Update
            </Button>
          </Group>
        </>
      </Modal>
    </FormProvider>
  )

  return (
    <PageTemplate minWidth={800}>
      <>
        <Title>User permissions</Title>
        <Space />
        {element}
      </>
    </PageTemplate>
  )
}

function FormData() {
  const form = useFormContext()

  form.watch()

  return <CodeHighlight language="json" code={JSON.stringify(form.getValues(), null, 4)}></CodeHighlight>
}

interface CheckboxPanelProps {
  title: string
  scopes: Partial<typeof SCOPES>
  user: User
  userSelection: User
}

const CheckboxPanel = ({ user, userSelection, title, scopes }: CheckboxPanelProps) => {
  const form = useFormContext()

  const handleCheckboxChange = (key: Scope, checked: boolean) => {
    if (checked) {
      form.setValue('scopes', form.getValues('scopes')?.concat([key]))
    } else {
      form.setValue(
        'scopes',
        form.getValues('scopes')?.filter((scope) => scope !== key),
      )
    }
  }

  const scopeChangeAllowed = (scope: Scope) => {
    if (isAuthorized({ user, requiredRole: 'admin' })) {
      return { allowed: true }
    }
    if (!isAuthorized({ user, requiredRole: userSelection?.role })) {
      return { allowed: false, message: 'You are not allowed to change scopes for this user' }
    }
    if (!isAuthorized({ user, requiredScopes: [scope] })) {
      return { allowed: false, message: 'You do not have this scope' }
    }

    return { allowed: true }
  }

  useWatch({ name: 'scopes', control: form.control })

  return (
    <Box className={classes.box}>
      <Title size={15} mt={4} mb={8}>
        {title}
      </Title>
      {entries(scopes).map(([key, scope]) => {
        const scopeName = key.split(':')[1]
        const { allowed, message } = scopeChangeAllowed(key)
        const isChecked = form.getValues('scopes')?.includes(key)

        return (
          <div key={key}>
            <Tooltip
              label={<Text size={'sm'}>{`${message}`}</Text>}
              position="left"
              withArrow
              disabled={allowed}
              withinPortal
            >
              <Grid
                style={{
                  display: 'flex',
                  alignItems: 'center',
                  marginBottom: '2px',
                  filter: !allowed ? 'grayscale(1)' : '',
                }}
              >
                <Grid.Col span={2}>
                  <Flex direction="row">
                    <Checkbox
                      checked={isChecked}
                      size="xs"
                      id={key}
                      color="blue"
                      disabled={!allowed}
                      onChange={(e) => handleCheckboxChange(key, e.target.checked)}
                    />
                    <Space pl={10} />
                    <Badge radius={4} size="xs" color={scopeColor(scopeName)}>
                      {scopeName}
                    </Badge>
                  </Flex>
                </Grid.Col>
                <Grid.Col span="auto">
                  <Text size={'md'}>{scope.description}</Text>
                </Grid.Col>
              </Grid>
            </Tooltip>
          </div>
        )
      })}
    </Box>
  )
}

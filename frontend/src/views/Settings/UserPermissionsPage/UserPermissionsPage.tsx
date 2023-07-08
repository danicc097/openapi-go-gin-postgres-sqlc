import _, { capitalize, random } from 'lodash'
import React, { Fragment, forwardRef, memo, useEffect, useReducer, useState } from 'react'
import type { RestDemoWorkItemCreateRequest, Scope, Scopes, UpdateUserAuthRequest, User } from 'src/gen/model'
import { getContrastYIQ, roleColor } from 'src/utils/colors'
import { joinWithAnd } from 'src/utils/format'
import SCOPES from 'src/scopes'

import type { Role } from 'src/client-validator/gen/models'
import PageTemplate from 'src/components/PageTemplate'
import type { ValidationErrors } from 'src/client-validator/validate'
import { updateUserAuthorization, useUpdateUserAuthorization } from 'src/gen/user/user'
import { validateField } from 'src/utils/validation'
import { RestDemoWorkItemCreateRequestDecoder, UpdateUserAuthRequestDecoder } from 'src/client-validator/gen/decoders'
import { newFrontendSpan } from 'src/TraceProvider'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import { getGetCurrentUserMock } from 'src/gen/user/user.msw'
import type { PathType, RecursiveKeyOf, RequiredKeys } from 'src/types/utils'
import jsonSchema from 'src/client-validator/gen/dereferenced-schema.json'
import {
  Avatar,
  Badge,
  Button,
  Flex,
  Space,
  Text,
  Title,
  Select,
  type SelectItem,
  Group,
  Modal,
  Checkbox,
  Code,
  Card,
  Box,
  type DefaultMantineColor,
  Grid,
  type MultiSelectValueProps,
  Tooltip,
  Divider,
} from '@mantine/core'
import { Prism } from '@mantine/prism'
import { notifications } from '@mantine/notifications'
import { IconCheck } from '@tabler/icons'
import RoleBadge from 'src/components/RoleBadge'
import { entries, keys } from 'src/utils/object'
import { css } from '@emotion/css'
import ROLES from 'src/roles'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import ErrorCallout from 'src/components/ErrorCallout/ErrorCallout'
import { ApiError } from 'src/api/mutator'
import { AxiosError } from 'axios'
import { isAuthorized } from 'src/services/authorization'
import { asConst } from 'json-schema-to-ts'
import type { components, schemas } from 'src/types/schema'
import { useForm } from 'react-hook-form'

type RequiredUserAuthUpdateKeys = RequiredKeys<UpdateUserAuthRequest>

const REQUIRED_USER_AUTH_UPDATE_KEYS: Record<RequiredUserAuthUpdateKeys, boolean> = {}

interface SelectUserItemProps extends React.ComponentPropsWithoutRef<'div'> {
  label: string
  value: User['email']
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

const SelectRoleItem = forwardRef<HTMLDivElement, SelectRoleItemProps>(
  ({ value, ...others }: SelectRoleItemProps, ref) => {
    return (
      <div ref={ref} {...others}>
        <RoleBadge role={value} />
      </div>
    )
  },
)

const SelectUserItem = forwardRef<HTMLDivElement, SelectUserItemProps>(
  ({ value, user, ...others }: SelectUserItemProps, ref) => {
    return (
      <div ref={ref} {...others}>
        <Group noWrap spacing="lg" align="center">
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <Avatar size={35} radius="xl" data-test-id="header-profile-avatar" alt={user?.username}>
              {user.fullName
                ?.split(' ')
                .map((n) => n[0].toUpperCase())
                .join('')}
            </Avatar>
            <Space p={5} />
            <RoleBadge role={user.role} />
          </div>

          <div style={{ marginLeft: 'auto' }}>{user?.email}</div>
        </Group>
      </div>
    )
  },
)

export default function UserPermissionsPage() {
  const [userSelection, setUserSelection] = useState<User>(null)
  const [roleSelection, setRoleSelection] = useState<Role>(null)
  const [userOptions, setUserOptions] = useState<Array<SelectUserItemProps>>(undefined)
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
    if (userOptions === undefined) {
      setUserOptions(
        allUsers
          ? allUsers.map((user) => ({
              label: user.email,
              value: user.email,
              user,
            }))
          : undefined,
      )
    } else {
      setUserOptions(userOptions)
    }
  }, [allUsers, userOptions])

  const [calloutError, setCalloutError] = useState<AppError>(null)

  // const { mutateAsync: updateUserAuthorization } = useUpdateUserAuthorization()

  const form = useForm<UpdateUserAuthRequest>({
    defaultValues: {},
  })

  const submitRoleUpdate = async () => {
    const span = newFrontendSpan('submitRoleUpdate')
    try {
      const updateUserAuthRequest = UpdateUserAuthRequestDecoder.decode(form.values)
      const payload = await updateUserAuthorization(userSelection.userID, updateUserAuthRequest)
      console.log('fulfilled', payload)
      notifications.show({
        id: ToastId.FormSubmit,
        title: 'Submitted',
        color: 'primary',
        icon: <IconCheck size="1.2rem" />,
        autoClose: 15000,
        message: 'Submitted',
      })
      setCalloutError(null)
    } catch (error) {
      console.error(error)
      if (error.validationErrors) {
        setCalloutError(error.validationErrors)
        console.log('error')
        return
      }
      setCalloutError(error)
    }
    span?.end()
  }

  const handleError = (errors: typeof form.formState.errors) => {
    if (Object.values(errors).some((v) => v)) {
      console.log('some errors found')
      console.log(errors)

      // TODO validate everything and show ALL validation errors
      // (we dont want to show very long error messages in each form
      // field, just that the field has an error,
      // so all validation errors are aggregated with full description in a callout)
      try {
        UpdateUserAuthRequestDecoder.decode(form.values)
        setCalloutError(null)
      } catch (error) {
        if (error.validationErrors) {
          setCalloutError(error.validationErrors)
          console.error(error)
          return
        }
        setCalloutError(error)
      }
    }
  }

  const onRoleSelectableChange = (role) => {
    console.log(role)
    form.setValue('role', role)
  }

  const onEmailSelectableChange = (email) => {
    const user = allUsers.find((user) => user.email === email)
    console.log(user)
    setUserSelection(user)
    form.setValue('role', user.role)
    form.setValue('scopes', user.scopes)
  }

  const [isModalVisible, setIsModalVisible] = useState(false)
  const closeModal = () => setIsModalVisible(false)
  const showModal = () => setIsModalVisible(true)

  const onRoleUpdateSubmit = async () => {
    showModal()
  }

  interface CheckboxPanelProps {
    title: string
    scopes: Partial<typeof SCOPES>
  }

  const CheckboxPanel = ({ title, scopes }: CheckboxPanelProps) => {
    const [disabledScopes, setDisabledScopes] = useState<string[]>([])

    const handleCheckboxChange = (key: Scope, checked: boolean) => {
      if (checked) {
        form.setValue('scopes', [...form.getValues('scopes'), key])
      } else {
        form.setValue(
          'scopes',
          form.getValues('scopes')?.filter((scope) => scope !== key),
        )
      }
    }

    const scopeChangeAllowed = (scope: Scope) => {
      return isAuthorized({ user, requiredRole: 'admin' }) || isAuthorized({ user, requiredScopes: [scope] })
    }

    return (
      <Box
        mb={12}
        sx={(theme) => ({
          backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.colors.gray[0],
          borderRadius: theme.radius.md,
          padding: '4px 16px',
        })}
      >
        <Title size={15} mt={4} mb={8}>
          {title}
        </Title>
        {entries(scopes).map(([key, scope]) => {
          const scopeName = key.split(':')[1]
          const isDisabled = !scopeChangeAllowed(key)
          const isChecked = form.getValues('scopes')?.includes(key)

          return (
            <div key={key}>
              <Tooltip
                label={<Text size={10}>You do not have this scope</Text>}
                position="left"
                withArrow
                disabled={!isDisabled}
                withinPortal
              >
                <Grid
                  style={{
                    display: 'flex',
                    alignItems: 'center',
                    marginBottom: '2px',
                    filter: isDisabled ? 'grayscale(1)' : '',
                  }}
                >
                  <Grid.Col span={2}>
                    <Flex direction="row">
                      <Checkbox
                        checked={isChecked}
                        size="xs"
                        id={key}
                        color="blue"
                        disabled={isDisabled}
                        onChange={(e) => handleCheckboxChange(key, e.target.checked)}
                      />
                      <Space pl={10} />
                      <Badge radius={4} size="xs" color={scopeColor(scopeName)}>
                        {scopeName}
                      </Badge>
                    </Flex>
                  </Grid.Col>
                  <Grid.Col span="auto">
                    <Text size={14}>{scope.description}</Text>
                  </Grid.Col>
                </Grid>
              </Tooltip>
            </div>
          )
        })}
      </Box>
    )
  }

  // TODO: export to generic helper. should have custom error just for api calls, see mutator.ts
  const getErrors = () => {
    if (!calloutError) return []

    // TODO: instead construct based on spec HTTPError which internally could have validationError array with loc, etc, see FastAPI template
    // or a regular error with message, title, detail, status...
    // and construct appropriately
    if (calloutError instanceof ApiError) return [calloutError.message]

    // external call error
    if (calloutError instanceof AxiosError) return [calloutError.message]

    // client side validation
    return calloutError?.errors?.map((v, i) => `${v.invalidParams.name}: ${v.invalidParams.reason}`)
  }

  const demoWorkItemCreateSchema = asConst(jsonSchema.definitions.RestDemoWorkItemCreateRequest)
  console.log(demoWorkItemCreateSchema)

  const element = (
    <>
      {JSON.stringify(calloutError)}
      <ErrorCallout title="Error updating user" errors={getErrors()} />
      <Space pt={12} />
      <Title size={12}>
        <Text>Form</Text>
      </Title>
      <Prism language="json">{JSON.stringify(form, null, 4)}</Prism>
      <Space pt={12} />
      <form
        onSubmit={form.handleSubmit(onRoleUpdateSubmit, handleError)}
        // error={getErrors()}
      >
        <Flex direction="column">
          <Select
            label="Select user to update"
            itemComponent={SelectUserItem}
            data-test-subj="updateUserAuthForm__selectable"
            searchable
            filter={(value, item) =>
              item.label?.toLowerCase().includes(value.toLowerCase().trim()) ||
              item.description?.toLowerCase().includes(value.toLowerCase().trim())
            }
            data={userOptions ?? []}
            onChange={onEmailSelectableChange}
          />
          {/* TODO: {renderSuperSelect<UpdateUserAuthRequest, 'role'>({
            formKey: 'role',
            form,
            options: roleOptions,
            requiredFormKeys: REQUIRED_USER_AUTH_UPDATE_KEYS,
            onSuperSelectChange: onRoleSuperSelectChange,
          })} */}
        </Flex>
        <Space pt={12} />
        {userSelection?.email && (
          <>
            <Divider m={8} />

            {isAuthorized({ user, requiredRole: userSelection.role }) && (
              <>
                <Select
                  label={
                    <Title size={15} mt={4} mb={4}>
                      Update role
                    </Title>
                  }
                  itemComponent={SelectRoleItem}
                  data-test-subj="updateUserAuthForm__selectable_Role"
                  defaultValue={userSelection.role}
                  data={roleOptions ?? []}
                  value={form.getValues('role')}
                  onChange={onRoleSelectableChange}
                />
                <Space pt={12} />
              </>
            )}
            <Title size={15} mt={4} mb={4}>
              Update scopes
            </Title>
            <Card shadow="md" padding="lg" radius="md" withBorder>
              {entries(scopeEditPanels).map(([group, scopes]) => (
                <CheckboxPanel
                  key={group}
                  title={group.replace(/-/g, ' ').replace(/^\w{1}/g, (c) => c.toUpperCase())}
                  scopes={scopes}
                />
              ))}
            </Card>
            <Space pt={24} />
            <Button disabled={userSelection === null} data-test-subj="updateUserAuthForm__submit" onClick={showModal}>
              Update authorization settings
            </Button>
          </>
        )}
      </form>
      <Modal
        opened={isModalVisible}
        title={
          <Text weight={'bold'} size={18}>
            Update auth information
          </Text>
        }
        onClose={closeModal}
        data-test-subj="updateUserAuthForm__confirmModal"
      >
        <>
          {`You're about to update auth information for `}
          <strong>{userSelection?.email}</strong>.<p>Are you sure you want to do this?</p>
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
    </>
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

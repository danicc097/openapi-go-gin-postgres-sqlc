import _, { capitalize, random } from 'lodash'
import React, { Fragment, forwardRef, useEffect, useReducer, useState } from 'react'
import type { Scope, Scopes, UpdateUserAuthRequest, UserResponse } from 'src/gen/model'
import { getContrastYIQ, roleColor } from 'src/utils/colors'
import { joinWithAnd } from 'src/utils/format'
import scopes from '@scopes'
import roles from '@roles'
import type { Role } from 'src/client-validator/gen/models'
import PageTemplate from 'src/components/PageTemplate'
import type { ValidationErrors } from 'src/client-validator/validate'
import { useUpdateUserAuthorization } from 'src/gen/user/user'
import { Form, useForm, type UseFormReturnType } from '@mantine/form'
import { validateField } from 'src/utils/validation'
import { UpdateUserAuthRequestDecoder } from 'src/client-validator/gen/decoders'
import { newFrontendSpan } from 'src/TraceProvider'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import { getGetCurrentUserMock } from 'src/gen/user/user.msw'
import type { RequiredKeys } from 'src/types/utils'
import { Avatar, Badge, Button, Flex, Space, Text, Title, Select, type SelectItem, Group } from '@mantine/core'
import { Prism } from '@mantine/prism'
import { Modal } from 'mantine-design-system'
import { notifications } from '@mantine/notifications'
import { IconCheck } from '@tabler/icons'

type RequiredUserAuthUpdateKeys = RequiredKeys<UpdateUserAuthRequest>

const REQUIRED_USER_AUTH_UPDATE_KEYS: Record<RequiredUserAuthUpdateKeys, boolean> = {}

interface ItemProps extends React.ComponentPropsWithoutRef<'div'> {
  label: string
  value: UserResponse['email']
  user: UserResponse
}

const Item = forwardRef<HTMLDivElement, ItemProps>(({ value, user, ...others }: ItemProps, ref) => {
  const color = roleColor(user.role)

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
          <Badge
            size="sm"
            radius="md"
            style={{ backgroundColor: color, color: getContrastYIQ(color) === 'black' ? 'whitesmoke' : 'black' }}
          >
            {user.role}
          </Badge>
        </div>

        <div style={{ marginLeft: 'auto' }}>{user?.email}</div>
      </Group>
    </div>
  )
})

export default function UserPermissionsPage() {
  const [userSelection, setUserSelection] = useState<UserResponse>(null)
  const [userOptions, setUserOptions] = useState<Array<ItemProps>>(undefined)

  const [allUsers] = useState(
    [...Array(20)].map((x, i) => {
      return getGetCurrentUserMock()
    }),
  )

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

  const [calloutErrors, setCalloutError] = useState<ValidationErrors>(null)

  // const { mutateAsync: updateUserAuthorization } = useUpdateUserAuthorization()

  const form = useForm<UpdateUserAuthRequest>({
    initialValues: {},
    validateInputOnChange: true,
    validate: {
      role: (v, vv, path) => validateField(UpdateUserAuthRequestDecoder, path, vv),
      scopes: (v, vv, path) => validateField(UpdateUserAuthRequestDecoder, path, vv),
    },
  })

  const fetchData = async () => {
    try {
      const updateUserAuthRequest = UpdateUserAuthRequestDecoder.decode(form.values)
      // const payload = await updateUserAuthorization({ data: updateUserAuthRequest, id: '' })
      // console.log('fulfilled', payload)
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
        // TODO setFormErrors instead
        console.log('error')
        return
      }
      setCalloutError(error)
    }
  }

  const handleError = (errors: typeof form.errors) => {
    if (Object.values(errors).some((v) => v)) {
      console.log('some errors found')

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

  const onEmailSelectableChange = (email) => {
    console.log(email)
    const user = allUsers.find((user) => user.email === email)
    setUserSelection(user)
    form.values.role = user.role
  }

  const [isModalVisible, setIsModalVisible] = useState(false)
  const closeModal = () => setIsModalVisible(false)
  const showModal = () => setIsModalVisible(true)

  const submitRoleUpdate = async () => {
    const span = newFrontendSpan('fetchData')
    await fetchData()
    span.end()

    closeModal()
  }

  const onRoleUpdateSubmit = async () => {
    showModal()
  }

  // TODO:
  // const roleOptions: EuiSuperSelectProps<Role>['options'] = Object.keys(roles).map((key: Role) => {
  //   const name = capitalize(key.replace(/([A-Z])/g, ' $1').trim())
  //   return {
  //     value: key,
  //     inputDisplay: (
  //       <EuiHealth color={roleColor(key)} style={{ lineHeight: 'inherit' }}>
  //         {name}
  //       </EuiHealth>
  //     ),
  //     dropdownDisplay: (
  //       <Fragment>
  //         <EuiHealth color={roleColor(key)} style={{ lineHeight: 'inherit' }}>
  //           {name}
  //         </EuiHealth>
  //       </Fragment>
  //     ),
  //   }
  // })

  const onRoleSuperSelectChange = (value) => {
    form.values.role = value
  }

  const getErrors = () =>
    calloutErrors ? calloutErrors?.errors?.map((v, i) => `${v.invalidParams.name}: ${v.invalidParams.reason}`) : null

  const element = (
    <>
      {getErrors()}
      <Space pt={12} />
      <Title size={1}>
        <Text>Form</Text>
      </Title>
      <Prism language="json">{JSON.stringify(form, null, 4)}</Prism>
      <Space pt={12} />
      <form
        onSubmit={form.onSubmit(onRoleUpdateSubmit, handleError)}
        // error={getErrors()}
      >
        <Flex direction="column">
          <Select
            itemComponent={Item}
            aria-label="Searchable example"
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
          <Button
            disabled={userSelection === null}
            data-test-subj="updateUserAuthForm__submit"
          >{`Update role for ${userSelection.email}`}</Button>
        )}
      </form>
      <Modal
        opened={isModalVisible}
        title={`Update auth information`}
        onClose={closeModal}
        data-test-subj="updateUserAuthForm__confirmModal"
      >
        <>
          {_.unescape(`You're about to update auth information for `)}
          <strong>{userSelection?.email}</strong>.<p>Are you sure you want to do this?</p>
          <Button onClick={closeModal}>Cancel</Button>
          <Button onClick={submitRoleUpdate}>Update</Button>
        </>
      </Modal>
    </>
  )

  return (
    <PageTemplate>
      <>
        <Title>{_.unescape(`Manually update a user's role and scopes.`)}</Title>
        <Space />
        {element}
      </>
    </PageTemplate>
  )
}

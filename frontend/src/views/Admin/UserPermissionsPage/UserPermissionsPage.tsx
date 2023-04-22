import {
  EuiBadge,
  EuiButton,
  EuiCallOut,
  EuiCodeBlock,
  EuiConfirmModal,
  EuiFlexGroup,
  EuiFlexItem,
  EuiForm,
  EuiFormRow,
  EuiHeader,
  EuiHealth,
  EuiIcon,
  EuiSelectable,
  EuiSelectableOption,
  EuiSpacer,
  EuiSuperSelect,
  EuiSuperSelectOption,
  EuiSuperSelectProps,
  EuiText,
  EuiTextColor,
  EuiTitle,
} from '@elastic/eui'
import _, { capitalize, random } from 'lodash'
import React, { Fragment, useEffect, useReducer, useState } from 'react'
import type { Scope, Scopes, UpdateUserAuthRequest, UserResponse } from 'src/gen/model'
import { roleColor } from 'src/utils/colors'
import { joinWithAnd } from 'src/utils/format'
import scopes from '@scopes'
import roles from '@roles'
import type { Role } from 'src/client-validator/gen/models'
import PageTemplate from 'src/components/PageTemplate/PageTemplate'
import type { ValidationErrors } from 'src/client-validator/validate'
import { useUpdateUserAuthorization } from 'src/gen/user/user'
import { useForm, UseFormReturnType } from '@mantine/form'
import { validateField } from 'src/utils/validation'
import { UpdateUserAuthRequestDecoder } from 'src/client-validator/gen/decoders'
import { newFrontendSpan } from 'src/TraceProvider'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import { createLabel, renderSuperSelect } from 'src/utils/forms'
// import { getGetCurrentUserMock } from 'src/gen/user/user.msw'
import UserAvatar from 'src/components/UserAvatar/UserAvatar'
import type { RequiredKeys } from 'src/types/utils'

type RequiredUserAuthUpdateKeys = RequiredKeys<UpdateUserAuthRequest>

const REQUIRED_USER_AUTH_UPDATE_KEYS: Record<RequiredUserAuthUpdateKeys, boolean> = {}

export default function UserPermissionsPage() {
  const [userSelection, setUserSelection] = useState<UserResponse>(null)
  const [userOptions, setUserOptions] = useState<Array<EuiSelectableOption<any>>>(undefined)
  const { addToast, dismissToast, theme } = useUISlice()

  const [allUsers] = useState<UserResponse[]>(
    [...Array(1)].map((x, i) => {
      return {
        role: 'user',
        scopes: ['users:read'],
        apiKey: null,
        teams: null,
        projects: null,
        user: {
          userID: 'c7fd2433-dbb7-4612-ab13-ddb0d3404728',
          username: 'user_2',
          email: 'user_2@email.com',
          firstName: 'Name 2',
          lastName: 'Surname 2',
          fullName: 'Name 2 Surname 2',
          hasPersonalNotifications: false,
          hasGlobalNotifications: true,
          createdAt: new Date('2023-04-01T06:24:22.390699Z'),
          deletedAt: null,
          timeEntries: null,
          userAPIKey: null,
          teams: null,
          workItems: null,
        },
      }
    }),
  )

  useEffect(() => {
    if (userOptions === undefined) {
      setUserOptions(
        allUsers
          ? allUsers.map((user) => ({
              label: (
                <>
                  <UserAvatar user={user} size={'s'}></UserAvatar> <>{user?.user.email}</>
                </>
              ),
              append: <EuiBadge color={roleColor(user.role)}>{user.role}</EuiBadge>,
              role: user.role,
              showIcons: false,
              value: user,
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
      addToast({
        id: ToastId.FormSubmit,
        title: 'Submitted',
        color: 'primary',
        iconType: 'check',
        toastLifeTimeMs: 15000,
        text: 'Submitted',
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

  const onEmailSelectableChange = (newOptions) => {
    setUserOptions(newOptions)
    const user = newOptions.filter((option) => !!option?.checked)[0]?.value
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

  const title = (
    <div>
      <EuiFlexGroup gutterSize="s" alignItems="center" responsive={false}>
        <EuiFlexItem grow={false}>
          <EuiIcon type="eraser" size="m" />
        </EuiFlexItem>

        <EuiFlexItem>
          <EuiTitle size="xs">
            <h3 style={{ color: 'dodgerblue' }}>Update user authorization</h3>
          </EuiTitle>
        </EuiFlexItem>
      </EuiFlexGroup>

      <EuiText size="s">
        <p>
          <EuiTextColor color="subdued">{_.unescape(`Manually update a user's role and scopes.`)}</EuiTextColor>
        </p>
      </EuiText>
    </div>
  )

  const roleOptions: EuiSuperSelectProps<Role>['options'] = Object.keys(roles).map((key: Role) => {
    const name = capitalize(key.replace(/([A-Z])/g, ' $1').trim())
    return {
      value: key,
      inputDisplay: (
        <EuiHealth color={roleColor(key)} style={{ lineHeight: 'inherit' }}>
          {name}
        </EuiHealth>
      ),
      dropdownDisplay: (
        <Fragment>
          <EuiHealth color={roleColor(key)} style={{ lineHeight: 'inherit' }}>
            {name}
          </EuiHealth>
        </Fragment>
      ),
    }
  })

  const onRoleSuperSelectChange = (value) => {
    form.values.role = value
  }

  const getErrors = () =>
    calloutErrors ? calloutErrors?.errors?.map((v, i) => `${v.invalidParams.name}: ${v.invalidParams.reason}`) : null

  const renderModal = (): any => {
    return isModalVisible ? (
      <EuiConfirmModal
        title={`Update auth information`}
        onCancel={closeModal}
        onConfirm={submitRoleUpdate}
        cancelButtonText="Cancel"
        confirmButtonText="Update"
        defaultFocusedButton="confirm"
        buttonColor="warning"
        data-test-subj="updateUserAuthForm__confirmModal"
      >
        <>
          {_.unescape(`You're about to update auth information for `)}
          <strong>{userSelection.user.email}</strong>.<p>Are you sure you want to do this?</p>
        </>
      </EuiConfirmModal>
    ) : null
  }

  const element = (
    <>
      {getErrors()}
      <EuiSpacer></EuiSpacer>
      <EuiTitle size="xs">
        <EuiText>Form</EuiText>
      </EuiTitle>
      <EuiCodeBlock language="json">{JSON.stringify(form, null, 4)}</EuiCodeBlock>
      <EuiSpacer></EuiSpacer>
      <EuiForm
        component="form"
        onSubmit={form.onSubmit(onRoleUpdateSubmit, handleError)}
        isInvalid={Boolean(form.errors.length)}
        error={getErrors()}
      >
        <EuiFlexGroup direction="column">
          <EuiFormRow fullWidth label="Select the user's email">
            <EuiSelectable
              aria-label="Searchable example"
              data-test-subj="updateUserAuthForm__selectable"
              searchable
              searchProps={{
                onChange: (searchValue, matchingOptions) => {
                  null
                },
              }}
              options={userOptions ?? []}
              singleSelection="always"
              onChange={onEmailSelectableChange}
            >
              {(list, search) => (
                <Fragment>
                  {search}
                  {list}
                </Fragment>
              )}
            </EuiSelectable>
          </EuiFormRow>
          {renderSuperSelect<UpdateUserAuthRequest, 'role'>({
            formKey: 'role',
            form,
            options: roleOptions,
            requiredFormKeys: REQUIRED_USER_AUTH_UPDATE_KEYS,
            onSuperSelectChange: onRoleSuperSelectChange,
          })}
        </EuiFlexGroup>
        <EuiSpacer />
        <EuiButton
          fill
          type="submit"
          isDisabled={userSelection === null}
          color="primary"
          data-test-subj="updateUserAuthForm__submit"
        >{`Update role for ${userSelection?.user.email ?? '...'}`}</EuiButton>
      </EuiForm>
      {renderModal()}
    </>
  )

  return (
    <PageTemplate header={{ children: title }} content={element} restrictWidth={'40vw'} buttons={[]} offset={100} />
  )
}

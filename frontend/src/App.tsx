import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import { CreateUserRequestDecoder } from './client-validator/gen/decoders'
import { useCreateUserMutation } from './redux/slices/gen/internalApi'
import { useUI } from 'src/hooks/ui'
import { Alert, Button, Group, PasswordInput, Text, TextInput } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { Prism } from '@mantine/prism'
import type { Decoder } from 'src/client-validator/gen/helpers'
import type { schemas } from 'src/types/schema'
import type { ValidationErrors } from 'src/client-validator/validate'
import { useForm } from '@mantine/form'
import { validateField } from 'src/utils/validation'
import { AttributeKeys, newFrontendSpan, tracer } from 'src/TraceProvider'
import roles from '@roles'
import scopes from '@scopes'

// TODO role changing see:
// https://codesandbox.io/s/wonderful-danilo-u3m1jz?file=/src/TransactionsTable.js
// data driven components:
// best: https://icflorescu.github.io/mantine-datatable/examples/row-context-menu
// https://github.com/Kuechlin/mantine-data-grid
// https://codesandbox.io/s/react-table-datagrid-forked-r19mf7

/*
TODO landing page is a demonstration of openapi+codegen workflow in both back and front:
1. show how project works itself
backend:
openapi -> gen -> postgen (pending sqlc merging, if necessary) -> implement logic
frontend:
openapi -> gen (rtk + client side validation) -> automatic form validation and queries
2. dummy form with complex schema: datetime, patterns, enums

TODO opentelemetry:

@opentelemetry/exporter-jaeger currently is only supported on Node.js.
For web environments, as Jaeger supports zipkin format reporting
(https://www.jaegertracing.io/docs/1.35/apis/#zipkin-formats-stable),
you can use https://github.com/open-telemetry/opentelemetry-js/tree/main/packages/opentelemetry-exporter-zipkin
to report to Jaeger backends.

*/

type RequiredUserCreateKeys = RequiredKeys<schemas['CreateUserRequest']>

const REQUIRED_USER_CREATE_KEYS: Record<RequiredUserCreateKeys, boolean> = {
  username: true,
  email: true,
  password: true,
}

function App() {
  // TODO object with validation errors and api response errors
  // and extracted accordingly
  const [calloutErrors, setCalloutError] = useState<ValidationErrors>(null)

  const { addToast } = useUI()
  const [createUser, createUserResult] = useCreateUserMutation()

  type CreateUserRequestForm = schemas['CreateUserRequest'] & {
    passwordConfirm: string
  }

  const form = useForm<CreateUserRequestForm>({
    // TODO blank
    initialValues: { username: 'user', email: 'user@mail', password: '12341234', passwordConfirm: '12341234' },
    validateInputOnChange: true,
    validate: {
      username: (v, vv, path) => validateField(CreateUserRequestDecoder, path, vv),
      email: (v, vv, path) => validateField(CreateUserRequestDecoder, path, vv),
      password: (v, vv, path) => validateField(CreateUserRequestDecoder, path, vv),
      passwordConfirm: (v, vv, path) => (v !== vv.password ? 'Passwords do not match' : null),
    },
  })

  const fetchData = async () => {
    try {
      const createUserRequest = CreateUserRequestDecoder.decode(form.values)

      const payload = await createUser(createUserRequest).unwrap()
      console.log('fulfilled', payload)
      addToast('done')
      setCalloutError(null)
    } catch (error) {
      if (error.validationErrors) {
        setCalloutError(error.validationErrors)
        // TODO setFormErrors instead
        console.error(error)
        addToast('error')
        return
      }
      setCalloutError(error)
    }
  }

  const renderResult = () =>
    createUserResult ? (
      <Prism style={{ textAlign: 'left' }} language="json">
        {JSON.stringify(createUserResult, null, '\t')}
      </Prism>
    ) : null

  // TODO handle ValidationErrors(🆗) and api response errors
  // "error": {
  // 	"status": 409,  -->statusCodeToReasonPhrase[statusCode]
  // 	"data": {
  // 		"error": "error creating user",
  // 		"message": "username --- already exists"
  // 	}
  // },
  const renderErrors = () =>
    calloutErrors ? (
      <Alert
        style={{ textAlign: 'start' }}
        icon={<IconAlertCircle size={16} />}
        title={`${calloutErrors.message}`}
        color="red"
      >
        {calloutErrors?.errors?.map((v, i) => (
          <p key={i} style={{ margin: '4px' }}>
            • <strong>{v.invalidParams.name}</strong>: {v.invalidParams.reason}
          </p>
        ))}
      </Alert>
    ) : null

  const hasValidationErrors = (field: string): boolean => {
    let hasError = false
    calloutErrors?.errors?.forEach((v) => {
      if (v.invalidParams.name === field) {
        hasError = true
      }
    })

    return hasError
  }

  const handleError = (errors: typeof form.errors) => {
    if (Object.values(errors).some((v) => v)) {
      console.log('some errors found')

      // TODO validate everything and show ALL validation errors
      // (we dont want to show very long error messages in each form
      // field, just that the field has an error,
      // so all validation errors are aggregated with full description in a callout)
      try {
        CreateUserRequestDecoder.decode(form.values)
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

  const handleSubmit = async (values: typeof form.values, e) => {
    e.preventDefault()
    const span = newFrontendSpan('fetchData')
    fetchData()
    span.end()
  }

  return (
    <div className="App" style={{ maxWidth: '500px', minWidth: '400px', textAlign: 'left' }}>
      <div>
        <div>{renderResult()}</div>
        <div>{renderErrors()}</div>
      </div>
      <div className="card" style={{ display: 'flex', flexDirection: 'column' }}>
        <form onSubmit={form.onSubmit(handleSubmit, handleError)}>
          <TextInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['email']}
            label="Email"
            placeholder="mail@example.com"
            {...form.getInputProps('email')}
          />
          <TextInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['username']}
            label="Username"
            placeholder="username"
            {...form.getInputProps('username')}
          />
          <PasswordInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['password']}
            label="Password"
            placeholder="password"
            {...form.getInputProps('password')}
          />
          <PasswordInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['passwordConfirm']}
            label="Confirm password"
            placeholder="password"
            {...form.getInputProps('passwordConfirm')}
          />
          <Group position="right" mt="md">
            <Button type="submit">Submit</Button>
          </Group>
        </form>
      </div>
    </div>
  )
}

export default App

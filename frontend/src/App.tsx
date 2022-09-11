import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import { CreateUserRequestDecoder } from './client-validator/gen/decoders'
import { useCreateUserMutation } from './redux/slices/gen/internalApi'
import { useUI } from 'src/hooks/ui'
import { Alert, Button, Group, PasswordInput, Text, TextInput } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { Prism } from '@mantine/prism'
import { Decoder } from 'src/client-validator/gen/helpers'
import type { schemas } from 'src/types/schema'
import type { ValidationErrors } from 'src/client-validator/validate'

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



code highl. - https://mantine.dev/others/prism/

*/

type RequiredUserCreateKeys = RequiredKeys<schemas['CreateUserRequest']>

const REQUIRED_USER_CREATE_KEYS: Record<RequiredUserCreateKeys, boolean> = {
  username: true,
  email: true,
  password: true,
}

function App() {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [errors, setError] = useState<ValidationErrors>(null)

  const { addToast } = useUI()
  const [createUser, createUserResult] = useCreateUserMutation()

  // TODO
  // onChange: validate the whole thing on each field change,
  // and for each field that has validation error
  // AND value != "" --> set formError[field] = true
  // onSubmit: renderError with everything, and set formError[field] = true for all
  // of them
  const fetchData = async () => {
    try {
      const createUserRequest = CreateUserRequestDecoder.decode({
        email: email,
        password: 'password',
        username: username,
      })

      const payload = await createUser(createUserRequest).unwrap()
      console.log('fulfilled', payload)
      addToast('done')
      setError(null)
    } catch (error) {
      if (error.validationErrors) {
        setError(error.validationErrors)
        // TODO setFormErrors instead
        console.error(error)
        addToast('error')
        return
      }
      setError(error)
    }
  }

  const renderResult = () =>
    createUserResult ? (
      <Prism style={{ textAlign: 'left' }} language="json">
        {JSON.stringify(createUserResult, null, '\t')}
      </Prism>
    ) : null

  // TODO handle ValidationErrors(🆗) and api response errors
  const renderErrors = () =>
    errors ? (
      <Alert
        style={{ textAlign: 'start' }}
        icon={<IconAlertCircle size={16} />}
        title={`${errors.message}`}
        color="red"
      >
        {errors?.errors?.map((v, i) => (
          <p key={i} style={{ margin: '4px' }}>
            • <strong>{v.invalidParams.name}</strong>: {v.invalidParams.reason}
          </p>
        ))}
      </Alert>
    ) : null

  const hasErrors = (field: string): boolean => {
    let hasError = false
    errors?.errors?.forEach((v) => {
      if (v.invalidParams.name === field) {
        hasError = true
      }
    })

    return hasError
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    fetchData()
  }

  return (
    <div className="App" style={{ maxWidth: '500px', minWidth: '400px', textAlign: 'left' }}>
      <div>
        <div>{renderResult()}</div>
        <div>{renderErrors()}</div>
      </div>
      <form onSubmit={handleSubmit}>
        <div className="card" style={{ display: 'flex', flexDirection: 'column' }}>
          <TextInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['email']}
            label="Email"
            name="email"
            // TODO formErrors[field] instead of true (e.g. passwords not matching is
            // outside openapi spec)
            error={hasErrors('email') ? true : null}
            // TODO abstract generic onChange(name, value, decoder)
            onChange={(e) => {
              setEmail(e.target.value)
            }}
            placeholder="mail@example.com"
          />
          <TextInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['username']}
            label="Username"
            name="username"
            error={hasErrors('username') ? true : null}
            onChange={(e) => {
              setUsername(e.target.value)
            }}
            placeholder="username"
          />
          <PasswordInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['password']}
            label="Password"
            name="password"
            error={hasErrors('password') ? true : null}
            onChange={(e) => {
              setUsername(e.target.value)
            }}
            placeholder="password"
          />
          <PasswordInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['password']}
            label="Confirm password"
            name="password"
            error={hasErrors('password') ? true : null}
            onChange={(e) => {
              setUsername(e.target.value)
            }}
            placeholder="password"
          />
          <Group position="right" mt="md">
            <Button type="submit">Submit</Button>
          </Group>
        </div>
      </form>
    </div>
  )
}

export default App

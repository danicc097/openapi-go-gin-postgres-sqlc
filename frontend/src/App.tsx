import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import { CreateUserRequestDecoder } from './client-validator/gen/decoders'
import { useCreateUserMutation } from './redux/slices/gen/internalApi'
import { useUI } from 'src/hooks/ui'
import { Alert, Button, Group, Text, TextInput } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { ValidationErrors } from 'src/client-validator/validate'
import { Prism } from '@mantine/prism'
import { components } from 'src/types/schema'

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

type RequiredUserCreateKeys = RequiredKeys<components['schemas']['CreateUserRequest']>

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

  return (
    <div className="App" style={{ maxWidth: '500px', minWidth: '400px', textAlign: 'left' }}>
      <div>
        <div>{renderResult()}</div>
        <div>{renderErrors()}</div>
      </div>
      <form
        onSubmit={(e) => {
          e.preventDefault()
          fetchData()
        }}
      >
        <div className="card" style={{ display: 'flex', flexDirection: 'column' }}>
          <TextInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['email']}
            label="Email"
            name="email"
            error={errors?.errors?.map((v) => {
              if (v.invalidParams.name === 'email') return true
            })}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="mail@example.com"
          />
          <TextInput
            withAsterisk={REQUIRED_USER_CREATE_KEYS['username']}
            label="Username"
            name="username"
            error={errors?.errors?.map((v) => {
              if (v.invalidParams.name === 'username') return true
            })}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="username"
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

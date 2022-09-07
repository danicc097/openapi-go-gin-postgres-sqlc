import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import { CreateUserRequestDecoder } from './client-validator/gen/decoders'
import { useCreateUserMutation } from './redux/slices/gen/internalApi'
import { useUI } from 'src/hooks/ui'

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

function App() {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [error, setError] = useState('')

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

  return (
    <div className="App">
      <div>
        <div>
          <pre>{JSON.stringify(createUserResult)}</pre>
        </div>
        <div>
          <pre>{JSON.stringify(error)}</pre>
        </div>
      </div>
      <form
        onSubmit={(e) => {
          e.preventDefault()
          fetchData()
        }}
      >
        <div className="card" style={{ display: 'flex', flexDirection: 'column' }}>
          <label htmlFor="email">Email:</label>
          <input type="text" id="email" onChange={(e) => setEmail(e.target.value)} name="Email"></input>
          <label htmlFor="username">Username:</label>
          <input type="text" id="username" onChange={(e) => setUsername(e.target.value)} name="Username"></input>
          Create user
          <input type="submit" value="Submit" />
        </div>
      </form>
    </div>
  )
}

export default App

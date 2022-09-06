import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import { useCreateUserMutation } from './store/internalApi'

// TODO role changing see:
// https://codesandbox.io/s/wonderful-danilo-u3m1jz?file=/src/TransactionsTable.js

function App() {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')

  const [createUser, createUserResult] = useCreateUserMutation()

  const fetchData = async () => {
    try {
      const payload = await createUser({
        email: email,
        password: 'fgsgefse',
        username: username,
      }).unwrap()
      console.log('fulfilled', payload)
    } catch (error) {
      console.error('rejected', error)
    }
  }

  return (
    <div className="App">
      <div>
        <div>{JSON.stringify(createUserResult)}</div>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <form>
        <div className="card" style={{ display: 'flex', flexDirection: 'column' }}>
          <label htmlFor="email">Email:</label>
          <input type="text" id="email" onChange={(e) => setEmail(e.target.value)} name="Email"></input>
          <label htmlFor="username">Username:</label>
          <input type="text" id="username" onChange={(e) => setUsername(e.target.value)} name="Username"></input>
          Create user
          <input
            onClick={(e) => {
              e.preventDefault()
              fetchData().catch((error) => console.error('rejected', error))
            }}
            onSubmit={(e) => {
              fetchData().catch((error) => console.error('rejected', error))
            }}
            type="submit"
            value="Submit"
          />
        </div>
      </form>
    </div>
  )
}

export default App

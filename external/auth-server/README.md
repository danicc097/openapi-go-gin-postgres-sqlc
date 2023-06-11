# users

`1.dev.json`: users to test out login in development.
Users loaded in ascending filename order, therefore ensure e2e tests always
loaded last.
TODO:
in dev mode would need to create from JSON with random external id (external
auth login won't be used for those specific users).
Frontend auth logic will have a guard if env = dev, and if we set a `DEV_USER`
is set
```ts
import devUsers from ".../1.dev.json"

const DEV_USER: <keyof devUsers | null> = "admin"
```

then we call dedicated dev api login route which only works if app_env is dev
and creates access token for user via authn service

+ initial-data knows its "dev" env so it just registers 1.dev.json users with a
  random external_id, which wont be used.
+ having x-api-key for messing with backend (already done).
On the other hand,  E2E gets different env to test login, redirects, token creation, etc.
E2E doesn't really need initial-data since we will create everything from scratch, ie we will test
register and update cases when logging in via auth server first and consequent
times


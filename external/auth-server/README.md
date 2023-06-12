# users

`*.e2e.json`: users used in E2E tests. NOTE: used in tests.
`*.admin.json`: shared admin users. NOTE: used in tests.
`*.dev.json`: predefined users to check out login behavior in development. NOTE: used for
dev initial-data (TODO:).
`*.local.json`: you may additional users to check out login behavior in
development. Not tracked, add as desired.

Users are loaded in ascending filename order, therefore ensure e2e tests always
loaded last.
TODO:
in dev mode would need to create from JSON with random external id (external
auth login won't be used for those specific users).

IMPORTANT: even if we run local oidc server, for dev initial-data we need existing users.
instead of this mess of changing current user in dev and adding adhoc backend behavior,
we can have simple logic in internal/services/authentication.go:61
if env=dev to update external_id just as with superadmin.
(we could have initial-data register all users in users/base.json actually when app
env is dev or ci... the externalID is already known. we dont care about
local.json since we can manually log in to test things). however, base.json
users are used to create dummy entities.
IMPORTANT: do not change how superAdmin initial-data works, since its shared for
all envs
this way initial-data creates users found in *.dev.json, *.admin.json with nil
external id (instead of empty string).
Would need change:
```sql
create unique index on users (external_id)
where
  external_id is not null;
```
and can test out what happens when users change by changing local.json users.
oidc-server:
-  should watch /data/users for changes and update users accordingly. add mutex.
-  should not silently override users, instead shutdown completely
  so that user overrides from local.json (or duplicated between e2e and base in base.json) are not allowed. much easier to work it
  and reason about.


below is not needed anymore. will login with auth server in development
and backend authetnication checks if AppEnv == dev or ci as explained above
```ts
// ... Frontend auth hook ....
// if (process.env.NODE_ENV === "dev") {
//   const localUsers = import(".../local.json") // with symlink
//   const baseUsers = import(".../base.json") // with symlink

//   authServerUsers = { ...localUsers, ...baseUsers } //

//   const DEV_USER: <keyof authServerUsers | null> = "admin" // auto login
// }
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


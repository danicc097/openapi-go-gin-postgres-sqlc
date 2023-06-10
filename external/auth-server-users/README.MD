`1.dev.json`: users to test out login in development.
TODO would need to find best strategy to load them in initial-data cmd. Or maybe have dev mode without login and tokens in frontend.
could just have env var to switch auth server login required for dev on and off.
+ initial-data knows its "dev" env so it just registers certain users.
+ having x-api-key for messing with backend (already done).
On the other hand,  E2E gets different env to test login, redirects, token creation, etc.
E2E doesn't really need initial-data since we will create everything from scratch, ie we will test
register and update cases when logging in via auth server first and consequent
times

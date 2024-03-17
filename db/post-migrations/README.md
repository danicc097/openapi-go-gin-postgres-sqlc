All `*.up.sql` files are always executed like regular migrations after main `migrations`'s `up` regardless of `post-migrations` current
version.

`post-migrations` always applies its down migrations before execution.

Post-migration commands should be idempotent. Do not use `*.down.sql` files.

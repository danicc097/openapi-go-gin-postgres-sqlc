# oidc-server

OpenID Connect development server based on
https://github.com/zitadel/oidc/tree/main/example/server. Features

# Setup

Authorization server listens on port `10001`. Expose accordingly if using
containers.


# Runtime environment variables

- `ISSUER`: fully qualified domain name.
- `DATA_DIR`: absolute path to stored mock data. e.g. `/data`

# Required files

- `${DATA_DIR}/users/*.json`: JSON files with key-value pairs of users for easier
  testing. Keys are ignored. Server will raise errors at login page if duplicated IDs are
  found for easier debugging. The `${DATA_DIR}/users` folder is continuously watched for changes. See
  `storage/user.go`'s `User` for available fields.

- `${DATA_DIR}/redirect_uris.txt`: valid redirect URIs to load at startup.

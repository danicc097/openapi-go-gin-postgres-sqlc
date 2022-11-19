# openapi-go-gin-postgres-sqlc

[![Go Report Card](https://goreportcard.com/badge/github.com/danicc097/openapi-go-gin-postgres-sqlc)](https://goreportcard.com/report/github.com/danicc097/openapi-go-gin-postgres-sqlc)
[![GoDoc](https://pkg.go.dev/badge/github.com/danicc097/openapi-go-gin-postgres-sqlc)](https://pkg.go.dev/github.com/danicc097/openapi-go-gin-postgres-sqlc)
[![tests](https://github.com/danicc097/openapi-go-gin-postgres-sqlc/actions/workflows/tests.yaml/badge.svg)](https://github.com/danicc097/openapi-go-gin-postgres-sqlc/actions/workflows/tests.yaml)

API-first and Database-first approach with OpenAPI v3, sqlc+xo codegen,
generated backend, frontend client and validators and an unimaginative title.

## What's this for?

Your OpenAPI v3 spec becomes a real single source of truth for the full stack. Any
change to it is validated and cascades down to:

- **frontend**: generated API queries (`rtk-query`). User-friendly [generated] client-side validation
  (fork of `openapi-typescript-validator`).
- **backend**: generated Gin server (custom `oapi-codegen` and post-generation).
  Request and response validation (`kin-openapi`). Generated CRUD and index queries via `xo`
  and custom queries via `sqlc` by leveraging custom `xo` template generation
  that ensures compatibility.

Additionally, it features OpenTelemetry in both browser (automatic and
manual instrumentation) and backend services (manual instrumentation) via
Jaeger, TimescaleDB and Promscale (certified storage backend).

## Makefile alternative

You get dynamic `x` function and `x` options parameters documentation _and_
autocompletion (`complete -C project project`) for
free (from your own source itself and comments)
so they're always up to date without any repetitive work: add/remove functions
and flags at will.

![](.github/autodocs.png)

All calls to `x` functions are logged (distinguishing stdout and stderr) for easier parallel execution and nested
calls tracking:

![](.github/logging.png)

And help for any `x` function is easily searchable when the app inevitably grows
with `--x-help`:

![](.github/help-x-function.png)

## Code generation

Docs WIP

<!-- xo custom templates with cardinality comments for join generation, schema from structs, spec sync -->

## Architecture

Simplified:

![](.github/system-diagram.png)

## Changelog

- `v0.2`: discard `OpenAPITools/openapi-generator` with custom postgeneration for `deepmap/oapi-codegen`. Any
  change to fix broken generator functionality requires opening a PR or a disturbing
  amount of postgeneration code. Templates getting out of hand and also require
  a PR for custom functions. `oapi-codegen` much more extensible and idiomatic
  being already Go and properly maintained.

## Known issues

- Nested functions in `project`'s `x` functions will break automatic
  documentation for that particular function due to a bug in `declare` where the last nested function line
  number is returned instead of the parent.

## TODOs

- Pgx v5 + [logging](https://github.com/jackc/pgx/issues/1381) (dependent on
  sqlc support) to support custom struct tag scanning and allow switch to json
  camel

- Meaningful project name.

- System design docs/diagrams.

- [Oauth2 as openapi
  spec](https://github.com/ybelenko/oauth2_as_oas3_components/tree/master/dist/components)
  with endpoints clearly documented based on RFCs
  Can generate a mock with e.g.
  [openapi-mock](https://github.com/muonsoft/openapi-mock).

- For parsing kinopenapi validation errors to our own more user
  friendly ValidationError check out
  https://github.com/getkin/kin-openapi/pull/197
  as well as
  unpack_errors_test.go + adapt the generic `convertError` to our needs
  until data validation is updated:
  https://github.com/getkin/kin-openapi/pull/412

- frontend miscellanea:
  1. codegen from oas:
  - ~~ts client (openapitools)~~ keep away from this project
  - react-query components (fabien0102/openapi-codegen)
  - React Query hooks, Axios requests and Typescript types (rametta/rapini) generation
  - Redux toolkit has its [own
    generator](https://github.com/reduxjs/redux-toolkit/tree/master/packages/rtk-query-codegen-openapi)
    and can generate hooks. Uses rtk-query, in essenceequivalent to react-query.
    Creators don't use openapi so that's a red flag for the generator itself.

  2. state management:
    - zustand + react-query should by far cover all needs. Data will be heavily
      dependent on backend.

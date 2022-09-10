
# openapi-go-gin-postgres-sqlc

[![Go Report Card](https://goreportcard.com/badge/github.com/danicc097/openapi-go-gin-postgres-sqlc)](https://goreportcard.com/report/github.com/danicc097/openapi-go-gin-postgres-sqlc)
[![GoDoc](https://pkg.go.dev/badge/github.com/danicc097/openapi-go-gin-postgres-sqlc)](https://pkg.go.dev/github.com/danicc097/openapi-go-gin-postgres-sqlc)

API-first and Database-first approach with OpenAPI v3, sqlc codegen and an
unimaginative title.
Featuring an overhaul of the [Go Gin
server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go-gin-server.md)
generator templates and a sensible post-generation tool that allows you to use cleanly
structured, easily extendable code by smartly merging nodes
from your modified and generated files' abstract syntax trees.

You also get dynamic ``x`` function and ``x`` flags arguments documentation _and_ autocompletion for
free (from your own source and comments)
so they're always up to date without any repetitive work: add/remove functions
and flags at will.

![](.github/autodocs.png)

## What's this for?

Your OpenAPI v3 spec becomes a real single source of truth for the full stack. Any
change to it is validated and cascades down to:
 - **frontend**: generated API queries (`rtk-query`) and user-friendly client-side validation
   (fork of `openapi-typescript-validator`).
 - **backend**: generated Gin server (`openapi-generator` and custom
   post-generation), request and response validation (`kin-openapi`).

## TODOs

  - GH actions

  - [Oauth2 as openapi
    spec](https://github.com/ybelenko/oauth2_as_oas3_components/tree/master/dist/components)
    with endpoints clearly documented based on RFCs
    Can generate a mock with e.g.
    [openapi-mock](https://github.com/muonsoft/openapi-mock).

  - equivalent of Python exception handler context manager but with global
  middleware on the api version route group:
      https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware
      in combination with internal/rest/rest.go renderErrorResponse
    For parsing kinopenapi validation errors to our own more user
    friendly ValidationError check out
    https://github.com/getkin/kin-openapi/pull/197
    as well as
    unpack_errors_test.go + adapt the generic ``convertError`` to our needs
    until data validation is updated:
    https://github.com/getkin/kin-openapi/pull/412


  - frontend miscellanea:
    1. codegen from oas:
      - ts client (openapitools)
      - react-query components (fabien0102/openapi-codegen)
      - React Query hooks, Axios requests and Typescript types (rametta/rapini)
      - Redux toolkit has its [own
        generator](https://github.com/reduxjs/redux-toolkit/tree/master/packages/rtk-query-codegen-openapi)
        and can generate hooks. Uses rtk-query.


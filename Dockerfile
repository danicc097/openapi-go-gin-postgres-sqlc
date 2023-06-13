# TODO :1.19.3-buster, ensure vendor up to date
FROM golang:1.19-alpine AS build

ARG DOCKER_UID
ARG DOCKER_GID

WORKDIR /go/src
COPY go.* ./
# RUN go mod download # it will access network/cache, which is not necessary with -mod=vendor
COPY . .
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/root/.cache/go-build \
  GOWORK=off go build -mod=vendor -o rest-server ./cmd/rest-server

# RUN groupadd -g $DOCKER_GID rootless
# RUN useradd -m rootless -u $DOCKER_UID -g $DOCKER_GID
# RUN chown -R $DOCKER_UID:$DOCKER_GID /go/src
# RUN chmod -R 775 /go/src

FROM alpine:3.15 AS runtime

# requires mounting /etc/ssl/certs/ca-certificates.crt:/etc/ssl/certs/ca-certificates.crt:ro
RUN apk --no-cache add ca-certificates

ENV GIN_MODE=release
COPY --from=build /go/src/rest-server \
  # bound to current app state
  /go/src/openapi.yaml \
  /go/src/scopes.json \
  /go/src/roles.json \
  /go/src/operationAuth.gen.json \
  ./

ENTRYPOINT [ "./rest-server" ]

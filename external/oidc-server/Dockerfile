FROM golang:1.19.3-buster as builder

WORKDIR /go/src

COPY ./go.* ./
RUN go mod download
COPY . ./
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/root/.cache/go-build \
  go build -o /go/src/server ./

FROM alpine:3.15 AS runtime

COPY --from=builder /go/src/server ./


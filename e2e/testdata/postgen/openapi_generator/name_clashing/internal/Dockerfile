# TODO adapt

FROM golang:1.18 AS build
WORKDIR /go/src
COPY gen ./gen
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]

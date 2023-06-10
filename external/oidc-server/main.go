package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/external/oidc-server/exampleop"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/external/oidc-server/storage"
)

func main() {
	ctx := context.Background()

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	issuer := os.Getenv("OIDC_ISSUER")
	storage := storage.NewStorage(storage.NewUserStore(issuer))

	port := "10001" // exposed on OIDC_SERVER_PORT

	router := exampleop.SetupServer(issuer, storage)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Default().Printf("listening at: %s", server.Addr)
	err := server.ListenAndServe()
	// if running directly localhost manually add certs
	// err := server.ListenAndServeTLS("certificates/localhost.pem", "certificates/localhost-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}

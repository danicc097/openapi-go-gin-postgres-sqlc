package postgen_test

import (
	"fmt"
	"os"
	"path"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/postgen"
)

func ExampleGenerateHandlers() {
	handlers := []postgen.Handler{
		{
			OperationId: "MyGeneratedOperationId1",
			Comment:     "MyGeneratedOperationId1 has this cool comment.",
			Origin:      "api_1.go",
		},
		{
			OperationId: "MyGeneratedOperationId2",
			Comment:     "MyGeneratedOperationId2 has this cool comment.",
			Origin:      "api_2.go",
		},
	}
	postgen.GenerateHandlers(handlers, os.Stdout)
	//Output:
	//package handlers
	//
	//import (
	//	"net/http"
	//
	//	"github.com/gin-gonic/gin"
	//)
	//
	//// MyGeneratedOperationId1 has this cool comment.
	//// Origin: api_1.go
	//func MyGeneratedOperationId1(c *gin.Context) {
	//	c.String(http.StatusNotImplemented, "501 not implemented")
	//}
	//
	//// MyGeneratedOperationId2 has this cool comment.
	//// Origin: api_2.go
	//func MyGeneratedOperationId2(c *gin.Context) {
	//	c.String(http.StatusNotImplemented, "501 not implemented")
	//}
}

func ExampleParseHandlers() {
	cwd, _ := os.Getwd()
	handlers := postgen.ParseHandlers(path.Join(cwd, "testdata/handlers.go"))

	hs := []postgen.Handler{}
	for _, v := range handlers {
		hs = append(hs, v)
	}

	fmt.Printf("Operation IDs: %s", hs)
	//Output:
	//Operation IDs: [CreateUsersWithArrayInput DeleteUser GetUserByName]
}

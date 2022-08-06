package postgen

import (
	"os"
)

func ExampleGenerateHandlers() {

	handlers := []Handler{
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
	GenerateHandlers(handlers, os.Stdout)
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

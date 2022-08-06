package postgen

import (
	"os"
)

func ExampleGenerateHandlers() {

	handlers := []Handler{
		{
			OperationId: "MyGeneratedOperationId1",
			Comment:     "MyGeneratedOperationId1 has this cool comment.",
		},
		{
			OperationId: "MyGeneratedOperationId2",
			Comment:     "MyGeneratedOperationId2 has this cool comment.",
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
	//func MyGeneratedOperationId1(c *gin.Context) {
	//	c.String(http.StatusNotImplemented, "501 not implemented")
	//}
	//
	//// MyGeneratedOperationId2 has this cool comment.
	//func MyGeneratedOperationId2(c *gin.Context) {
	//	c.String(http.StatusNotImplemented, "501 not implemented")
	//}
}

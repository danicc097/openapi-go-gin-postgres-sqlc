package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/a/:a", func(c *gin.Context) {
		a1 := c.Param("a")
		log.Printf("/a/:a route invoked. a=%q\n", a1)
	})
	router.GET("/a/:a/b/:b", func(c *gin.Context) {
		a2 := c.Param("a")
		b2 := c.Param("b")
		log.Printf("/a/:a/b/:b route invoked. a=%q, b=%q\n", a2, b2)
	})
	router.Run(":11080")
}

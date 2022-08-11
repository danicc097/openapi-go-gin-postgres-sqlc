package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FakeDataFile - test data_file to ensure it's escaped correctly
func FakeDataFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUsersWithArrayInput creates list of users with given input array.
func CreateUsersWithArrayInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// DeleteUser delete user.
func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

func GetUserByName(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUsersWithArrayInput creates list of users with given input array.
// Origin: api_user.go
func CreateUsersWithArrayInput(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// DeleteUser delete user.
// Origin: api_user.go
func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetUserByName get user by user name.
// Origin: api_user.go
func GetUserByName(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// LoginUser logs user into the system.
// Origin: api_user.go
func LoginUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// LogoutUser logs out current logged in user session.
// Origin: api_user.go
func LogoutUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// PlaceOrder place an order for a pet.
// Origin: api_store.go
func PlaceOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdatePet update an existing pet.
// Origin: api_pet.go
func UpdatePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdatePetWithForm updates a pet in the store with form data.
// Origin: api_pet.go
func UpdatePetWithForm(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UpdateUser updated user.
// Origin: api_user.go
func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// UploadFile uploads an image.
// Origin: api_pet.go
func UploadFile(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

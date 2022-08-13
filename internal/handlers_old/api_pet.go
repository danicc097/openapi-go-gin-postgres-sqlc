package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddPet - Add a new pet to the store
func AddPet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// DeletePet - Deletes a pet
func DeletePet(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// FindPetsByStatus - Finds Pets by status
func FindPetsByStatus(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// FindPetsByTags - Finds Pets by tags
// Deprecated
func FindPetsByTags(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetPetById - Find pet by ID
func GetPetById(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

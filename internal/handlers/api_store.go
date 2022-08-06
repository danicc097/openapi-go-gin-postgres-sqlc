package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteOrder - Delete purchase order by ID
func DeleteOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetInventory - Returns pet inventories by status
func GetInventory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetOrderById - Find purchase order by ID
func GetOrderById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

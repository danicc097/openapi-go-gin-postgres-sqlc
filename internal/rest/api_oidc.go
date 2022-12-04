package rest

import (
	"github.com/gin-gonic/gin"
)

func (h *Handlers) MyProviderLogin(c *gin.Context) {
	c.Set(skipRequestValidation, true)
}

func (h *Handlers) MyProviderCallback(c *gin.Context) {
	c.Set(skipRequestValidation, true)
}

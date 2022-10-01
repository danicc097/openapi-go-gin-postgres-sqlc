package rest

import (
	"net/http"

	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Store handles routes with the 'store' tag.
type Store struct {
	svc services.Store
	// add or remove services, etc. as required
}

// NewStore returns a new handler for the 'store' route group.
func NewStore(svc services.Store) *Store {
	return &Store{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Store) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []route{
		{
			Name:        string(deleteOrder),
			Method:      http.MethodDelete,
			Pattern:     "/store/order/:orderId",
			HandlerFunc: h.deleteOrder,
			Middlewares: h.middlewares(deleteOrder),
		},
		{
			Name:        string(getInventory),
			Method:      http.MethodGet,
			Pattern:     "/store/inventory",
			HandlerFunc: h.getInventory,
			Middlewares: h.middlewares(getInventory),
		},
		{
			Name:        string(getOrderById),
			Method:      http.MethodGet,
			Pattern:     "/store/order/:orderId",
			HandlerFunc: h.getOrderById,
			Middlewares: h.middlewares(getOrderById),
		},
		{
			Name:        string(placeOrder),
			Method:      http.MethodPost,
			Pattern:     "/store/order",
			HandlerFunc: h.placeOrder,
			Middlewares: h.middlewares(placeOrder),
		},
	}

	registerRoutes(r, routes, "/store", mws)
}

// middlewares returns individual route middleware per operation id.
func (h *Store) middlewares(opID storeOpID) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// deleteOrder delete purchase order by id.
func (h *Store) deleteOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// getInventory returns pet inventories by status.
func (h *Store) getInventory(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// getOrderById find purchase order by id.
func (h *Store) getOrderById(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// placeOrder place an order for a pet.
func (h *Store) placeOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

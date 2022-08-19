package handlers

import (
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	services "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/gin-gonic/gin"
)

// Store handles routes with the 'store' tag.
type Store struct {
	svc services.Store
	// add or remove services, etc. as required
}

// NewStore returns a new handler for the 'store' route group.
// Edit as required.
func NewStore(svc services.Store) *Store {
	return &Store{
		svc: svc,
	}
}

// Register connects handlers to an existing router group with the given middlewares.
// Generated method. DO NOT EDIT.
func (h *Store) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "DeleteOrder",
			Method:      http.MethodDelete,
			Pattern:     "/store/order/:orderId",
			HandlerFunc: h.DeleteOrder,
			Middlewares: h.middlewares("DeleteOrder"),
		},
		{
			Name:        "GetInventory",
			Method:      http.MethodGet,
			Pattern:     "/store/inventory",
			HandlerFunc: h.GetInventory,
			Middlewares: h.middlewares("GetInventory"),
		},
		{
			Name:        "GetOrderById",
			Method:      http.MethodGet,
			Pattern:     "/store/order/:orderId",
			HandlerFunc: h.GetOrderById,
			Middlewares: h.middlewares("GetOrderById"),
		},
		{
			Name:        "PlaceOrder",
			Method:      http.MethodPost,
			Pattern:     "/store/order",
			HandlerFunc: h.PlaceOrder,
			Middlewares: h.middlewares("PlaceOrder"),
		},
	}
	rest.RegisterRoutes(r, routes, "/store", mws)
}

// middlewares returns individual route middleware per operation id.
// Edit as required.
func (h *Store) middlewares(opID string) []gin.HandlerFunc {
	switch opID {
	default:
		return []gin.HandlerFunc{}
	}
}

// DeleteOrder delete purchase order by id.
func (h *Store) DeleteOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetInventory returns pet inventories by status.
func (h *Store) GetInventory(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetOrderById find purchase order by id.
func (h *Store) GetOrderById(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// PlaceOrder place an order for a pet.
func (h *Store) PlaceOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

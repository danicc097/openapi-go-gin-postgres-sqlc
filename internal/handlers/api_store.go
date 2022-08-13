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
	// add necessary services, etc. as required
}

// NewStore returns a new handler for the 'store' route group.
// Edit as required.
func NewStore(svc services.Store) *Store {
	return &Store{
		svc: svc,
	}
}

// Register connects the handlers to a router with the given middleware.
// GENERATED METHOD. Only Middlewares will be saved between runs.
func (t *Store) Register(r *gin.RouterGroup, mws []gin.HandlerFunc) {
	routes := []rest.Route{
		{
			Name:        "DeleteOrder",
			Method:      http.MethodDelete,
			Pattern:     "/store/order/:orderId",
			HandlerFunc: t.DeleteOrder,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "GetInventory",
			Method:      http.MethodGet,
			Pattern:     "/store/inventory",
			HandlerFunc: t.GetInventory,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "GetOrderById",
			Method:      http.MethodGet,
			Pattern:     "/store/order/:orderId",
			HandlerFunc: t.GetOrderById,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Name:        "PlaceOrder",
			Method:      http.MethodPost,
			Pattern:     "/store/order",
			HandlerFunc: t.PlaceOrder,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	rest.RegisterRoutes(r, routes, "/store", mws)
}

// DeleteOrder delete purchase order by id.
func (t *Store) DeleteOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetInventory returns pet inventories by status.
func (t *Store) GetInventory(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// GetOrderById find purchase order by id.
func (t *Store) GetOrderById(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

// PlaceOrder place an order for a pet.
func (t *Store) PlaceOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

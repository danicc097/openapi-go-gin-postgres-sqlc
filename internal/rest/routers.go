package rest

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route contains information for a single endpoint.
type Route struct {
	// Name is the route name.
	Name string
	// Method is the string for the HTTP method.
	Method string
	// Pattern is the URI pattern.
	Pattern string
	// Group is the router group.
	Group string
	// HandlerFunc is the handler function for this route.
	HandlerFunc gin.HandlerFunc
	// Middlewares represent this route's middleware chain.
	Middlewares []gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// RegisterRoute creates a router group, adds group middleware and registers routes with their own middleware.
func RegisterRoutes(router *gin.RouterGroup, rr []Route, group string, mws []gin.HandlerFunc) {
	if reg, _ := regexp.Compile("/([a-z]+)"); !reg.MatchString(group) {
		panic("Invalid router group: " + group)
	}

	gn := group
	if group == "/default" {
		gn = ""
	}

	rg := router.Group(gn)

	for _, mw := range mws {
		rg.Use(mw)
	}

	for _, r := range rr {
		p := strings.Replace(r.Pattern, gn, "", 1)
		handlers := append(r.Middlewares, r.HandlerFunc)

		switch r.Method {
		case http.MethodGet:
			rg.GET(p, handlers...)
		case http.MethodPost:
			rg.POST(p, handlers...)
		case http.MethodPut:
			rg.PUT(p, handlers...)
		case http.MethodPatch:
			rg.PATCH(p, handlers...)
		case http.MethodDelete:
			rg.DELETE(p, handlers...)
		}
	}
}

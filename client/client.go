package client

import (
	"github.com/erbesharat/goverapi/router"
	"github.com/gorilla/mux"
)

// Client contains the client configuration like http and database instances
type Client struct {
	Mux *mux.Router
}

// New create a new overapi client instance
func New() *Client {
	c := &Client{
		Mux: mux.NewRouter(),
	}

	// Add the routes defiend in routes.go to the http client
	c.SetupRoutes()

	return c
}

// SetupRoutes Adds routes to the http mux instance
func (c *Client) SetupRoutes() {
	for request, handler := range router.Routes {
		c.Mux.HandleFunc(request.Path, handler).Methods(request.Method)
	}
}

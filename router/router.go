package router

import (
	"net/http"

	"github.com/erbesharat/goverapi/actions"
)

type request struct {
	Path   string
	Method string
}

// Routes attaches each request to a http handler
var Routes = map[request]func(http.ResponseWriter, *http.Request){
	{Path: "/", Method: "GET"}: actions.Root,
}

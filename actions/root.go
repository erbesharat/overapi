package actions

import (
	"fmt"
	"net/http"
)

// Root handles the service's root path aka "/"
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn!")
}

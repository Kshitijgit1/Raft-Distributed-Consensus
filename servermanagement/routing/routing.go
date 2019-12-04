package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouting manages all the routes
func SetupRouting(r *mux.Router) *mux.Router {
	r.HandleFunc("/startRaft", StartRaft).Methods(http.MethodGet, http.MethodOptions)
	return r
}

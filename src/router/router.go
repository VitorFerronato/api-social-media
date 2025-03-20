package router

import (
	routes "api/src/router/rotas"

	"github.com/gorilla/mux"
)

// GenerateRouter will return the configurated routes
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigurateRouter(r)
}

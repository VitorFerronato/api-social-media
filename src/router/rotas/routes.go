package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func ConfigurateRouter(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postRoutes...)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}

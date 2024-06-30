package routes

import (
	"api/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Representa as rotas que serão utilizadas na API
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Funçao que insere as rotas no router para serem entao consumidas
func ConfigRoutes(rt *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {

		if route.AuthRequired == true {
			rt.HandleFunc(route.URI,
				middleware.Logger(middleware.Authentication(route.Function)),
			).Methods(route.Method)
		} else {
			rt.HandleFunc(route.URI,
				middleware.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return rt
}

package routes

import (
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

	for _, newRoutes := range routes {
		rt.HandleFunc(newRoutes.URI, newRoutes.Function).Methods(newRoutes.Method)
	}

	return rt
}

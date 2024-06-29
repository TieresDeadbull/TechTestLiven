package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Gera um router com as rotas configuradas
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}

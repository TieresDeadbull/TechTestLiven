package middleware

import (
	"api/src/auth"
	"api/src/response"
	"fmt"
	"net/http"
)

// Função que apresenta as informações do request
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Making request...\n Método %s \n URI %s \n Host %s \n ", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Verifica se usuario está autenticado para acessar a rota que exige esse requisito
func Authentication(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			response.Err(w, http.StatusUnauthorized, err)
			return
		}
		fmt.Println("Validating...")
		nextFunc(w, r)
	}
}

package middleware

import (
	"fmt"
	"net/http"
)

// Função que apresenta as informações do request
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Making request...\n Método %s \n URI %s \n Host %s \n ", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Verifica se usuario está autenticado para acessar a rota que exige esse requisito
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validating...")
		next(w, r)
	}
}

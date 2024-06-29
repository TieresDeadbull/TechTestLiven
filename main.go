package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Função principal da API
func main() {
	config.Load()

	fmt.Println("Running API")
	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}

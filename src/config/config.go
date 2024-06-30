package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexao com o banco MySql
	ConectionString = ""

	//porta onde API roda
	Port = 0

	//Chave usada para assinatura do token
	SecretKey []byte
)

// Inicializa variaveis de ambiente
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv(("API_PORT")))
	if err != nil {
		Port = 9000
	}

	ConectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}

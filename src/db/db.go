package db

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConectionString)
	if err != nil {
		fmt.Println("Erro ao tentar abrir o Banco de Dados")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

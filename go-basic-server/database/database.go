package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "example"
	user     = "postgres"
	password = "saferman14"
)

func InitDb() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		host,
		port,
		dbname,
		user,
		password,
	)

	db, err := sql.Open("pgx", dbInfo)
	if err != nil {
		panic(err)
	}

	log.Print("db connected")
	return db
}

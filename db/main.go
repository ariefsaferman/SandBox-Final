package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "db_demo"
	user     = "postgres"
	password = "postgres"
)

func main() {

	dbInfo := fmt.Sprintf("host=%s port=%d dbname%s user=%s password=%s sslmode=disable",
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
	defer db.Close()
	fmt.Println("connected database")
}

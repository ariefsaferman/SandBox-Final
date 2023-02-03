package main

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/db"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/server"
)

func main() {
	// connect db
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}

	// automigrate
	db.Get().AutoMigrate(entity.User{}, entity.Wallet{}, entity.Transaction{})

	// define server
	server.Init()
}

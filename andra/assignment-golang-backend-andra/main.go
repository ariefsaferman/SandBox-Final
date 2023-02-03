package main

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/db"
	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}
	server.Init()
}

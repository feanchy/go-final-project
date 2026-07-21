package main

import (
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	log.Println("PORT:", port)

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	err = db.Init(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = server.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}

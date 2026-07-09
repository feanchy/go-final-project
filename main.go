package main

import (
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"
	"os"
)

func main() {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = server.Run()
	if err != nil {
		log.Fatal()
	}
}

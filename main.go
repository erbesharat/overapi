package main

import (
	"log"
	"net/http"
	"os"

	"github.com/erbesharat/goverapi/db"
	"github.com/erbesharat/goverapi/fetch"
	"github.com/erbesharat/goverapi/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("[overapi]: Error loading .env file: %s", err.Error())
	}

	if len(os.Args) < 2 {
		log.Fatalf("[overapi]: Please pass subcommand. Supported subcommands: ['serve', 'fetch']")
	}

	database, err := db.New(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatalf("[overapi]: Couldn't create the DB: %s", err.Error())
	}
	defer database.Close()

	switch os.Args[1] {
	case "serve":
		handler := server.New(database)
		log.Printf("[overapi]: Server is running on port %s", os.Getenv("PORT"))

		log.Fatal(http.ListenAndServe(":8080", handler))
	case "fetch":
		if err := database.Clear(); err != nil {
			log.Fatalf("Coudln't refresh the database: %s", err.Error())
		}
		err = fetch.FetchHeros(database)
		if err != nil {
			log.Fatalf("[overapi]: Couldn't fetch heroes: %s", err.Error())
		}
		err = fetch.FetchAbilities(database)
		if err != nil {
			log.Fatalf("[overapi]: Couldn't fetch abilities: %s", err.Error())
		}
	}
}

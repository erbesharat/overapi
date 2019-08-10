package main

import (
	"log"
	"net/http"
	"os"

	"github.com/erbesharat/goverapi/client"
	"github.com/erbesharat/goverapi/db"
	"github.com/erbesharat/goverapi/fetch"
	"github.com/joho/godotenv"

	// Import postgres driver for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
		c := client.New()

		log.Printf("[overapi]: Server is running on port %s", os.Getenv("PORT"))

		log.Fatal(http.ListenAndServe(":8080", c.Mux))
	case "fetch":
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

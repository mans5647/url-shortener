package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"url-shortener/handlers"
	"url-shortener/database"
)


func main() {

	if len(os.Args) == 1 {
		log.Fatalln("Not enough arguments")
	}

	servicePort, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatalln("Invalid value for port number")
	}

	if !database.OpenPostgresConnection(database.DefaultDsn) {
		log.Fatal("Failed to open db connection!")
	}

	if !database.AutoMigrateTables() {
		log.Fatal("Failed to migrate tables!")
	}

	// handlers
	http.HandleFunc("/shorten", handlers.ShortenUrlHandler)
	http.HandleFunc("/{code}", handlers.RedirectByLinkCodeHandler)
	http.HandleFunc("/clear", handlers.DeleteAllUrlsHandler)

	log.Printf("Starting url shortening service at http://localhost:%d", servicePort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil))
}
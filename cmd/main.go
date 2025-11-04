package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"url-shortener/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
	}

	port := os.Getenv("port")

	mongoHost := os.Getenv("mongo_host")
	mongoPort, _ := strconv.Atoi(os.Getenv("mongo_port"))
	mongoDb := os.Getenv("mongo_db")
	mongoColl := os.Getenv("mongo_collection")

	handler, err := handlers.NewDatabaseSequenceShortener(mongoDb, mongoColl, mongoHost, mongoPort)
	
	if err != nil {
		log.Fatal("Failed to create handler: %w", err)
	}
	
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.POST("/api/v1/shorten", handler.HandleShorten)
	e.GET("/api/v1/redirect/:code", handler.HandleRedirect)

	e.Start(fmt.Sprintf(":%s", port))
}
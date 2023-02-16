package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/martin1917/EWallet/internal/app"
	"github.com/martin1917/EWallet/internal/handler"
	"github.com/martin1917/EWallet/internal/infrastructure/repository"
	"github.com/martin1917/EWallet/internal/infrastructure/repository/pg"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := pg.NewPostgresDB(pg.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repositories := repository.NewPgRepositories(db)
	services := app.NewServices(repositories)
	handler := handler.NewHandler(services)
	router := handler.InitRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), router))
}

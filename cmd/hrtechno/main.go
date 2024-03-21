package main

import (
	"fmt"
	"os"

	"github.com/Georgi-Progger/survey-api/internal/app"
	"github.com/joho/godotenv"
)

// @title Filmlib API
// @version 1.0
// @description API Server for films and actors

// @host localhost:8091
// @BasePath /
// @schemes https http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := app.Config{}
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	port := os.Getenv("PORT")
	cfg.EchoPort = string(fmt.Sprintf(":%s", port))
	app, err := app.NewApplication(&cfg)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}
}

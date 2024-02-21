package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Georgi-Progger/survey-api/internal/handler"
	"github.com/Georgi-Progger/survey-api/internal/repository"
	"github.com/Georgi-Progger/survey-api/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type (
	App struct {
		Service *service.Service
		config  *Config
		Echo    *echo.Echo
	}

	Config struct {
		EchoPort string
		Db       Db
	}

	Db struct {
		DbConnect string
	}
)

func (a *App) Start() error {
	return a.Echo.Start(a.config.EchoPort)
}

func NewApplication(cfg *Config) (*App, error) {
	db, err := ConnectDatabase()
	if err != nil {
		log.Panic("connectionString error..")
	}
	repo := repository.NewRepository(db)
	surveyService := service.NewService(repo)
	handler := handler.NewHandler(surveyService)
	return &App{
		Service: surveyService,
		config:  cfg,
		Echo:    handler.InitRoutes(),
	}, nil
}

func ConnectDatabase() (*sql.DB, error) {
	cfg := Config{}
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	cfg.Db.DbConnect = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", cfg.Db.DbConnect)
	if err != nil {
		log.Panic("no connect to database...")
	}

	//defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
	return db, nil
}

package app

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
	"os"

	"github.com/Georgi-Progger/survey-api/internal/handler"
	"github.com/Georgi-Progger/survey-api/internal/repository"
	"github.com/Georgi-Progger/survey-api/internal/service"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	doMigration(db)
	repo := repository.NewRepository(db)
	surveyService := service.NewService(repo)
	handler := handler.NewHandler(surveyService)
	return &App{
		Service: surveyService,
		config:  cfg,
		Echo:    handler.InitRoutes(),
	}, nil
}

func doMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		os.Getenv("DB_NAME"), driver)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.Up()
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
	//dbSSL := os.Getenv("DB_SSLMODE")
	cfg.Db.DbConnect = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	//db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, dbSSL))
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

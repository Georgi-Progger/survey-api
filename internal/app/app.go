package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main.go/internal/service"

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
	surveyService := service.NewService(db)
	return &App{
		Service: surveyService,
		config:  cfg,
		Echo:    NewRouter(surveyService),
	}, nil
}

func ConnectDatabase() (*sql.DB, error) {
	cfg := Config{}
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	cfg.Db.DbConnect = os.Getenv("DB_CONNECT")
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

func NewRouter(s *service.Service) *echo.Echo {
	router := echo.New()

	router.Use(middleware.Logger())

	candidateGroup := router.Group("/candidate")

	candidateGroup.POST("/create", s.InsertCandidate)
	candidateGroup.GET("/interview", s.SelectInterview)

	return router
}

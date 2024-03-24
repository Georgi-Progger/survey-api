package handler

import (
	"net/http"

	_ "github.com/Georgi-Progger/survey-api/docs"
	"github.com/Georgi-Progger/survey-api/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()

	router.Use(middleware.Logger())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	candidateGroup := router.Group("/candidate")

	candidateGroup.POST("/create", h.InsertCandidate)
	candidateGroup.GET("/questions", h.SelectInterview)
	candidateGroup.POST("/registration", h.RegistrCandidate)
	candidateGroup.POST("/auth", h.AuthUser)

	interviewGroup := router.Group("/interview")
	interviewGroup.GET("/question", h.GetAllVQuestions)
	interviewGroup.POST("/video", h.UploadFile)

	return router
}

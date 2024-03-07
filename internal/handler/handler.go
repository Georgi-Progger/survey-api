package handler

import (
	"github.com/Georgi-Progger/survey-api/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	candidateGroup := router.Group("/candidate")

	candidateGroup.POST("/create", h.InsertCandidate)
	candidateGroup.GET("/questions", h.SelectInterview)
	candidateGroup.POST("/save/video", h.UploadFile)
	candidateGroup.POST("/registration", h.RegistrCandidate)

	return router
}

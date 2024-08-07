package handler

import (
	"net/http"

	_ "github.com/Georgi-Progger/survey-api/docs"
	"github.com/Georgi-Progger/survey-api/internal/service"
	"github.com/Georgi-Progger/survey-api/pkg/jwt"
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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	loginFreeCandidateGroup := router.Group("/candidate")

	loginFreeCandidateGroup.POST("/registration", h.RegistrCandidate)
	loginFreeCandidateGroup.POST("/auth", h.AuthUser)
	loginFreeCandidateGroup.POST("/newpass", h.ChangeUserPassword)

	candidateGroup := loginFreeCandidateGroup.Group("")

	candidateGroup.Use(userAuthMiddleware)
	candidateGroup.POST("/create", h.InsertCandidate)
	candidateGroup.GET("/questions", h.SelectInterview)

	userGroup := router.Group("/user")
	userGroup.Use(userAuthMiddleware)
	userGroup.GET("/info", h.getCurrentUserInfo)

	interviewGroup := router.Group("/interview")
	interviewGroup.Use(userAuthMiddleware)
	interviewGroup.GET("/question", h.GetAllVQuestions)
	interviewGroup.GET("/test", h.getAllTQuestions)
	interviewGroup.GET("/result/:id", h.getPonomarUserResult)
	interviewGroup.POST("/test", h.insertTQuestionAnswers)
	interviewGroup.POST("/video", h.UploadFile)

	adminGroup := router.Group("/admin")
	adminGroup.Use(adminAuthMiddleware)
	adminGroup.GET("/user/role", h.GetAllUsersWithForm)
	adminGroup.GET("/user/role/:id", h.GetAllUsersWithRole)
	adminGroup.GET("/user/interview/answers/:id", h.GetAllVAnswersByUserId)
	adminGroup.POST("/user/role/set", h.SetUserRole)
	return router
}

func candidateAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := jwt.ValidateRole(c, 1)
		if err != nil {
			return c.JSON(401, map[string]string{"error": err.Error()})
		}
		return next(c)
	}
}

func userAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := jwt.ValidateJWT(c)
		if err != nil {
			return c.JSON(401, map[string]string{"error": err.Error()})
		}
		return next(c)
	}
}

func adminAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := jwt.ValidateRole(c, 2)
		if err != nil {
			return c.JSON(401, map[string]string{"error": err.Error()})
		}
		return next(c)
	}
}

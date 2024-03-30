package handler

import (
	"log"
	"net/http"

	"github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func (h *Handler) getAllTQuestions(c echo.Context) error {
	res, err := h.services.TQuestion.GetAll()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get all video test questions"})
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) insertTQuestionAnswers(c echo.Context) error {
	var ans []model.UserTestAnswer
	if err := c.Bind(&ans); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON provided"}) 
	}
	err := h.services.TQuestion.InsertAnswers(jwt.GetUserIdFromContext(c), ans)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to save answers in db"}) 
	}
	return c.NoContent(http.StatusCreated)
}

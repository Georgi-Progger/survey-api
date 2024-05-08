package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllVQuestions(c echo.Context) error {
	resp, err := h.services.VQuestion.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get all video interviw questions: " + err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}
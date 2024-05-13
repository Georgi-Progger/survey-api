package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllVQuestions(c echo.Context) error {
	resp, err := h.services.VQuestion.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get all video interviw questions: " + err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllVAnswersByUserId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user id: " + err.Error()})
	}
	result, err := h.services.VQuestion.GetAllByUserIdWithQuestions(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid role id: " + err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
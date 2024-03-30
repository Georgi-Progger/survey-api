package handler

import (
	"net/http"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) InsertCandidate(c echo.Context) error {
	var candidate model.Candidate
	if err := c.Bind(&candidate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	if err := h.services.Candidate.Create(c.Request().Context(), candidate); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert candidate"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Candidate successfully created"})
}

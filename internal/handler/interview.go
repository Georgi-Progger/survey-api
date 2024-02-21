package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SelectInterview(c echo.Context) error {
	name := c.FormValue("nameInterview")

	res, err := h.services.Interview.GetInterviewQuestion(c.Request().Context(), name)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

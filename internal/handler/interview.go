package handler

import (
	"net/http"

	"github.com/Georgi-Progger/survey-api/internal/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SelectInterview(c echo.Context) error {
	if err := util.ValidateAdminRoleJWT(c); err != nil {
		return c.JSON(500, map[string]string{"err": "Auth error"})
	}
	name := c.FormValue("nameInterview")

	res, err := h.services.Interview.GetInterviewQuestion(c.Request().Context(), name)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

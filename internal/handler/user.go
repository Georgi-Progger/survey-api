package handler

import (
	"net/http"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) RegistrCandidate(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	res, err := password.Generate(10, 10, 10, false, false)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(res), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(passwordHash)
	h.services.User.Save(c.Request().Context(), user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Candidate successfully created"})
}

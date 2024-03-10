package handler

import (
	"fmt"
	"net/http"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) RegistrCandidate(c echo.Context) error {
	user := model.User{RoleId: 1}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	res, err := password.Generate(6, 6, 0, false, true)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(res)
	err = h.services.Sender.Send(user.Phonenumber, res)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(res), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(passwordHash)
	id, err := h.services.User.Save(c.Request().Context(), user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]int{"id": id})
}

func (h *Handler) AuthUser(c echo.Context) error {
	requestUser := model.User{}
	if err := c.Bind(&requestUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	dbUser, err := h.services.GetUserByPhonenumber(requestUser.Phonenumber)
	if err != nil || checkPasswords(requestUser.Password, dbUser.Password) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid credentials"})
	}
	jwtStr, err := util.GenerateJWT(dbUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "InternalServerError"})
	}

	return c.JSON(200, map[string]string{"jwt": jwtStr})
}

func checkPasswords(requestPassword, dbPassword string) bool {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(requestPassword), bcrypt.DefaultCost)
	if err != nil || string(passwordHash) != dbPassword {
		return false
	}
	return true
}

package handler

import (
	"fmt"
	"net/http"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

type userAuthDTO struct {
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"password"`
}
type userRegDTO struct {
	Phonenumber string `json:"phonenumber"`
}

// @Summary Registration
// @Tags Auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body userRegDTO true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /candidate/registration [post]
func (h *Handler) RegistrCandidate(c echo.Context) error {
	userInfo := userRegDTO{}
	if err := c.Bind(&userInfo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	res, err := password.Generate(6, 6, 0, false, true)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(res)
	err = h.services.Sender.Send(userInfo.Phonenumber, res)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(res), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user := model.User{
		Phonenumber: userInfo.Phonenumber,
		Password:    string(passwordHash),
		RoleId:      1,
	}

	id, err := h.services.User.Save(c.Request().Context(), user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]int{"id": id})
}

// @Summary SignIn
// @Tags Auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body userAuthDTO true "credentials"
// @Success 200 {string} map[string]string "jwt"
// @Failure 400 {object} error "Failed to decode request body. Invalid JSON"
// @Failure 500 {object} error "Failed to generate JWT"
// @Router /candidate/auth [post]
func (h *Handler) AuthUser(c echo.Context) error {
	requestUser := userAuthDTO{}
	if err := c.Bind(&requestUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	dbUser, err := h.services.GetUserByPhonenumber(requestUser.Phonenumber)
	if err != nil || checkPasswords(requestUser.Password, dbUser.Password) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid credentials"})
	}
	jwtStr, err := jwt.GenerateJWT(dbUser)
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

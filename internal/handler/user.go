package handler

import (
	"net/http"
	"strconv"

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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON: " + err.Error()})
	}
	res, err := password.Generate(6, 6, 0, false, true)
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
	err = h.services.Sender.Send(userInfo.Phonenumber, res)
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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON: " + err.Error()})
	}
	dbUser, err := h.services.GetUserByPhonenumber(requestUser.Phonenumber)
	if err != nil || !isEqualPassword(requestUser.Password, dbUser.Password) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid credentials: " + err.Error()})
	}
	jwtStr, err := jwt.GenerateJWT(dbUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "InternalServerError: " + err.Error()})
	}

	return c.JSON(200, map[string]string{"jwt": jwtStr})
}

func isEqualPassword(requestPassword, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(requestPassword))
	return err == nil
}

func (h *Handler) GetAllUsersWithRole(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid role id: " + err.Error()})
	}
	users, err := h.services.User.GetAllWithRole(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

type roleChangeDTO struct {
	RoleId int `json:"role_id"`
	UserId int `json:"user_id"`
}

func (h *Handler) SetUserRole(c echo.Context) error {
	var dto roleChangeDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON: " + err.Error()})
	}
	err := h.services.Role.SetRole(dto.UserId, dto.RoleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to change role: " + err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) ChangeUserPassword(c echo.Context) error {
	userInfo := userRegDTO{}
	if err := c.Bind(&userInfo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON: " + err.Error()})
	}
	res, err := password.Generate(6, 6, 0, false, true)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(res), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user, err := h.services.User.GetUserByPhonenumber(userInfo.Phonenumber)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	user.Password = string(passwordHash)
	err = h.services.User.Update(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = h.services.Sender.Send(userInfo.Phonenumber, res)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *Handler) getCurrentUserInfo(c echo.Context) error {
	id := jwt.GetUserIdFromContext(c)
	user, err := h.services.User.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Get user error: " + err.Error()})
	}
	candidate, err := h.services.Candidate.GetByUserId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Get candidate info error: " + err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"role_id":     user.RoleId,
		"first_name":  candidate.FirstName,
		"last_name":   candidate.LastName,
		"middle_name": candidate.MiddleName})
}

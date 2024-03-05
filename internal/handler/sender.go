package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Send(c echo.Context) error {
	test := h.services.Sender.Send("79289284851", "123")

	return c.JSON(http.StatusCreated, map[string]bool{"message": test})
}

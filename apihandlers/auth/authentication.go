package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	c.String(http.StatusOK, "hello world")
}
package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	user := c.Get("user")

	log.Println(user)

	return c.JSON(http.StatusOK, user)
}

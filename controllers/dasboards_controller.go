package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	user := c.Get("user")

	log.Println(user)
	type response struct {
		User    interface{} `json:"user"`
		Message string      `json:"message"`
	}

	return c.JSON(http.StatusOK, response{
		User:    user,
		Message: "hello TEAM ! 2024/01/14 00-15",
	})
}

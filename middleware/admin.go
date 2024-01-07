package middleware

import (
	"duongdx/example/models"
	"log"

	"github.com/labstack/echo"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(models.User)
		log.Println(user)

		if user.IsAdmin {
			return next(c)
		}

		return echo.ErrForbidden
	}
}

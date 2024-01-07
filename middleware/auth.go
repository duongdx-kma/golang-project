package middleware

import (
	"duongdx/example/models"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			log.Printf("Token does not exist")

			return echo.ErrUnauthorized
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return echo.ErrUnauthorized
		}

		if token.Valid {
			c.Set("user", models.User{
				ID:      int64(claims["id"].(float64)),
				Name:    claims["name"].(string),
				Age:     uint8(claims["age"].(float64)),
				Address: claims["address"].(string),
				IsAdmin: claims["is_admin"].(bool),
			})

			return next(c)
		}

		return echo.ErrUnauthorized
	}
}

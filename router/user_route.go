package router

import (
	"duongdx/example/controllers"
	database "duongdx/example/initializers"
	customMiddleware "duongdx/example/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func UserInit(e *echo.Echo, sql *database.SQL) {
	isLogedIn := customMiddleware.AuthenticationMiddleware
	IsAdmin := customMiddleware.IsAdmin
	userController := controllers.NewUserController(sql)

	// <--------login--------->
	e.POST("/login", func(c echo.Context) error {
		return userController.Authentication(c)
	})

	// <--------register--------->
	e.POST("/register", func(c echo.Context) error {
		user, err := userController.Create(c)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, user)
	})

	// <--------normal user--------->
	e.GET("/", controllers.Hello, isLogedIn)

	users := e.Group("/users", isLogedIn)
	users.GET("", func(c echo.Context) error {
		users, err := userController.FindAll(c)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, users)
	})

	users.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		user, err := userController.FindById(c, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, user)
	})

	// <------admin----->
	users.PUT("/:id", func(c echo.Context) error {
		id := c.Param("id")
		user, err := userController.Update(c, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, user)
	}, IsAdmin)
}

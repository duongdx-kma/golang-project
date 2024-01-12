package router

import (
	database "duongdx/example/initializers"
	"duongdx/example/websocket"
	"net/http"

	"github.com/labstack/echo"
)

func InitSocket(e *echo.Echo, sql *database.SQL) {
	projectHandler := websocket.NewProjectHandler(sql)
	taskHandler := websocket.NewTaskHandler(sql)

	e.GET("/projects", func(c echo.Context) error {
		projects, err := projectHandler.FindAll(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, projects)
	})
	e.GET("/projects/get-users/:project_id", func(c echo.Context) error {
		projects, err := projectHandler.GetUserId(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, projects)
	})

	e.GET("/tasks", func(c echo.Context) error {
		projects, err := taskHandler.FindAll(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, projects)
	})
}

package websocket

import (
	"duongdx/example/initializers"
	"duongdx/example/models"
	"duongdx/example/repositories"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	TaskRepository repositories.TaskInterface
}

func NewTaskHandler(sql *initializers.SQL) *Handler {
	return &Handler{
		TaskRepository: &repositories.TaskRepository{
			SQL: sql,
		},
	}
}

func (h *Handler) CreateTask(form models.CreateTaskSchema) (models.TaskSelected, error) {
	task, err := h.TaskRepository.CreateTask(form)

	if err != nil {
		return task, echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("Create task - Create task failed %s", err))
	}

	return task, nil
}

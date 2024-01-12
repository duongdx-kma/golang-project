package websocket

import (
	"duongdx/example/initializers"
	"duongdx/example/models"
	"duongdx/example/repositories"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func (h *Handler) FindAll(e echo.Context) ([]models.TaskSelected, error) {
	projectId, err := strconv.ParseInt(e.QueryParam("project_id"), 10, 64)

	if err != nil {
		log.Println("convert data project_id error")

		return []models.TaskSelected{}, err
	}

	projects, err := h.TaskRepository.FindAll(e.Request().Context(), projectId)

	if err != nil {
		log.Println("Find projects - find project by user_id failed", projects)
	}

	return projects, nil
}

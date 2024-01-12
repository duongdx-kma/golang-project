package websocket

import (
	"duongdx/example/initializers"
	"duongdx/example/models"
	"duongdx/example/repositories"
	"log"
	"strconv"

	"github.com/labstack/echo"
)

type ProjectHandler struct {
	ProjectRepository repositories.ProjectInterface
}

func NewProjectHandler(sql *initializers.SQL) *ProjectHandler {
	return &ProjectHandler{
		ProjectRepository: &repositories.ProjectRepository{
			SQL: sql,
		},
	}
}

func (h *ProjectHandler) CreateProject(form models.CreateProjectSchema) (models.ProjectSelected, error) {
	project, err := h.ProjectRepository.CreateProject(form)
	if err != nil {
		log.Printf("Create project - Create project failed %s", err)
	}

	return project, err
}

func (h *ProjectHandler) FindAll(e echo.Context) ([]models.ProjectSelected, error) {
	userId, err := strconv.ParseInt(e.QueryParam("user_id"), 10, 64)

	if err != nil {
		log.Println("convert data user_id error")

		return []models.ProjectSelected{}, err
	}

	projects, err := h.ProjectRepository.FindAll(e.Request().Context(), userId)

	if err != nil {
		log.Println("Find projects - find project by user_id failed", projects)
	}

	return projects, nil
}

func (h *ProjectHandler) GetUserId(e echo.Context) ([]int64, error) {
	users := make([]int64, 0)

	projectId, err := strconv.ParseInt(e.Param("project_id"), 10, 64)

	if err != nil {
		log.Println("convert data project_id error")

		return users, err
	}

	users, err = h.ProjectRepository.GetUserId(e.Request().Context(), projectId)

	if err != nil {
		log.Println("Find userssss error", err)
	}

	return users, nil
}

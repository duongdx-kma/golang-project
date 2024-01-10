package websocket

import (
	"duongdx/example/initializers"
	"duongdx/example/models"
	"duongdx/example/repositories"
	"log"
)

type ProjectHandler struct {
	// hub               *Hub
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
		log.Fatalf("Create project - Create project failed %s", err)
	}

	return project, err
}

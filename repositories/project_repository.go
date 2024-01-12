package repositories

import (
	"context"
	database "duongdx/example/initializers"
	"duongdx/example/models"
	"fmt"
	"log"
	"time"
)

type ProjectInterface interface {
	CreateProject(createProjectSchema models.CreateProjectSchema) (models.ProjectSelected, error)
	FindAll(ctx context.Context, userId int64) ([]models.ProjectSelected, error)
	GetUserId(ctx context.Context, ProjectId int64) ([]int64, error)
}

type ProjectRepository struct {
	SQL *database.SQL
}

func (db *ProjectRepository) CreateProject(createProjectSchema models.CreateProjectSchema) (models.ProjectSelected, error) {
	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	var project models.ProjectSelected
	now := time.Now()

	statement := `INSERT INTO projects(name, created_at, updated_at) VALUES (:name, :created_at, :updated_at)`
	project.Name = createProjectSchema.Name
	project.CreatedAt = &now
	project.UpdatedAt = &now
	result, err := db.SQL.DB.NamedExec(statement, project)

	if err != nil {
		log.Println("Insert data project failed", err)

		return models.ProjectSelected{}, err
	}

	// get last inserted project id
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("get data just have been created is fail", err)
	}
	project.ProjectId = lastId
	project.Users = createProjectSchema.Users
	statementRelation := `INSERT INTO project_user(project_id, user_id, created_at, updated_at) VALUES`

	for _, userId := range createProjectSchema.Users {
		statementRelation += fmt.Sprintf(
			"(%d, %d, '%s', '%s'),",
			lastId,
			userId,
			project.CreatedAt.Format("2006-01-02 15:04:05"),
			project.UpdatedAt.Format("2006-01-02 15:04:05"),
		)
	}

	statementRelation = statementRelation[:len(statementRelation)-1]
	_, err = db.SQL.DB.Exec(statementRelation)

	if err != nil {
		log.Printf(statementRelation, createProjectSchema.Users)
		log.Println("Insert relation", err)
	}

	return project, nil
}

func (db *ProjectRepository) FindAll(ctx context.Context, userId int64) ([]models.ProjectSelected, error) {
	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	var projects []models.Project
	query := `
		SELECT 
			projects.project_id,
			name,
			temp.user_id
		FROM projects 
		JOIN project_user ON projects.project_id = project_user.project_id
			AND project_user.user_id = ?
		JOIN (SELECT 
				project_user.project_id,
				project_user.user_id
			FROM project_user
		) AS temp ON temp.project_id = projects.project_id
		WHERE projects.deleted_at IS NULL
		ORDER BY projects.created_at DESC
	`

	err := db.SQL.DB.SelectContext(ctx, &projects, query, userId)

	if err != nil {
		log.Println("Get projects errorrrr: ", err)

		return []models.ProjectSelected{}, err
	}

	projectSelected := make(map[int64]models.ProjectSelected)
	temUser := make(map[int64][]int64)

	for _, project := range projects {
		temUser[project.ProjectId] = append(temUser[project.ProjectId], int64(project.UserId))

		projectSelected[project.ProjectId] = models.ProjectSelected{
			EventName: "create/project",
			ProjectId: project.ProjectId,
			Name:      project.Name,
			Users:     temUser[project.ProjectId],
		}
	}

	v := make([]models.ProjectSelected, 0)
	for _, value := range projectSelected {
		v = append(v, value)
	}

	return v, nil
}

func (db *ProjectRepository) GetUserId(ctx context.Context, projectId int64) ([]int64, error) {
	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	var projects []struct {
		ID int64 `json:"user_id" db:"user_id,omitempty"`
	}

	query := `
		SELECT 
			project_user.user_id
		FROM project_user
		WHERE project_user.project_id = ?
	`

	err := db.SQL.DB.SelectContext(ctx, &projects, query, projectId)

	if err != nil {
		log.Println("Get projects errorrrr: ", err)

		return make([]int64, 0), err
	}

	users := make([]int64, 0)

	for _, project := range projects {
		users = append(users, project.ID)
	}

	return users, nil
}

package repositories

import (
	database "duongdx/example/initializers"
	"duongdx/example/models"
	"fmt"
	"log"
	"time"
)

type ProjectInterface interface {
	CreateProject(createProjectSchema models.CreateProjectSchema) (models.ProjectSelected, error)
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
		log.Fatal("Insert data project failed", err)

		return models.ProjectSelected{}, err
	}

	// get last inserted project id
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal("get data just have been created is fail", err)
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
		log.Fatal("Insert relation", err)
	}

	return project, nil
}

package repositories

import (
	database "duongdx/example/initializers"
	"duongdx/example/models"
	"log"
	"time"
)

type TaskInterface interface {
	CreateTask(createTaskSchema models.CreateTaskSchema) (models.TaskSelected, error)
}

type TaskRepository struct {
	SQL *database.SQL
}

func (db *TaskRepository) CreateTask(createTaskSchema models.CreateTaskSchema) (models.TaskSelected, error) {
	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()
	var task models.TaskSelected
	now := time.Now()

	statement := `INSERT INTO tasks(title, description, project_id, created_at, updated_at) 
		VALUES (:title, :description, :project_id, :created_at, :updated_at)`
	task.Title = createTaskSchema.Title
	task.Description = createTaskSchema.Description
	task.CreatedAt = &now
	task.UpdatedAt = &now

	result, err := db.SQL.DB.NamedExec(statement, task)

	if err != nil {
		log.Fatal("Insert data task failed", err)

		return models.TaskSelected{}, err
	}

	// get last inserted task id
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal("get data just have been created is fail", err)
	}

	task.TaskId = lastId
	task.Users = createTaskSchema.Users

	return task, nil
}

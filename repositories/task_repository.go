package repositories

import (
	"context"
	database "duongdx/example/initializers"
	"duongdx/example/models"
	"log"
	"time"
)

type TaskInterface interface {
	CreateTask(createTaskSchema models.CreateTaskSchema) (models.TaskSelected, error)
	FindAll(ctx context.Context, ProjectId int64) ([]models.TaskSelected, error)
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
	task.ProjectId = createTaskSchema.ProjectId
	task.CreatedAt = &now
	task.UpdatedAt = &now

	result, err := db.SQL.DB.NamedExec(statement, task)

	if err != nil {
		log.Println("Insert data task failed", err)

		return models.TaskSelected{}, err
	}

	// get last inserted task id
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("get data just have been created is fail", err)
	}

	task.TaskId = lastId
	task.Users = createTaskSchema.Users

	return task, nil
}

func (db *TaskRepository) FindAll(ctx context.Context, ProjectId int64) ([]models.TaskSelected, error) {
	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	var tasks []models.Task
	query := `
		SELECT 
			tasks.task_id,
			title,
			description,
			tasks.project_id,
			temp.user_id
		FROM tasks 
		JOIN project_user ON tasks.project_id = project_user.project_id
		JOIN (SELECT 
				project_user.project_id,
				project_user.user_id
			FROM project_user
		) AS temp ON temp.project_id = tasks.project_id
		WHERE tasks.project_id = ? 
			AND tasks.deleted_at IS NULL
		ORDER BY tasks.created_at DESC
	`

	err := db.SQL.DB.SelectContext(ctx, &tasks, query, ProjectId)

	if err != nil {
		log.Println("Get projects errorrrr: ", err)

		return []models.TaskSelected{}, err
	}

	tasksResult := make(map[int64]models.TaskSelected)
	temUser := make(map[int64][]int64)

	for _, task := range tasks {
		temUser[task.TaskId] = append(temUser[task.TaskId], int64(task.UserId))

		tasksResult[task.TaskId] = models.TaskSelected{
			EventName:   "create/task",
			TaskId:      task.TaskId,
			Title:       task.Title,
			Description: task.Description,
			ProjectId:   task.ProjectId,
			Users:       temUser[task.TaskId],
		}
	}

	v := make([]models.TaskSelected, 0)
	for _, value := range tasksResult {
		v = append(v, value)
	}

	return v, nil
}

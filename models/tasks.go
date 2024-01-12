package models

import (
	"encoding/json"
	"time"
)

type Task struct {
	TaskId      int64      `json:"task_id,omitempty" db:"task_id,omitempty"`
	Description string     `json:"description" db:"description,omitempty"`
	Title       string     `json:"title" db:"title,omitempty"`
	ProjectId   int64      `json:"project_id" db:"project_id,omitempty"`
	UserId      int64      `json:"user_id" db:"user_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at,omitempty"`
}

type CreateTaskSchema struct {
	Title       string  `json:"title" db:"title,omitempty"`
	Description string  `json:"description" db:"description,omitempty"`
	ProjectId   int64   `json:"project_id" db:"project_id,omitempty"`
	Users       []int64 `json:"users"`
}

type UpdateTaskSchema struct {
	Title       string `json:"title" db:"title,omitempty"`
	Description string `json:"description" db:"description,omitempty"`
}

type TaskSelected struct {
	EventName   string     `json:"event_name,omitempty" db:"event_name,omitempty"`
	TaskId      int64      `json:"task_id,omitempty" db:"task_id,omitempty"`
	Title       string     `json:"title" db:"title,omitempty"`
	Description string     `json:"description" db:"description,omitempty"`
	Users       []int64    `json:"users"`
	ProjectId   int64      `json:"project_id,omitempty" db:"project_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at,omitempty"`
}

func (s *TaskSelected) Raw() []byte {
	raw, _ := json.Marshal(s)
	return raw
}

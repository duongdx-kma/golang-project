package models

import (
	"encoding/json"
	"time"
)

type Project struct {
	ProjectId int64      `json:"project_id,omitempty" db:"project_id,omitempty"`
	Name      string     `json:"name" db:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at,omitempty"`
}

type ProjectSelected struct {
	EventName string     `json:"event_name,omitempty" db:"event_name,omitempty"`
	ProjectId int64      `json:"project_id,omitempty" db:"project_id,omitempty"`
	Name      string     `json:"name" db:"name,omitempty"`
	Users     []int64    `json:"users" db:"users"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at,omitempty"`
}

type CreateProjectSchema struct {
	Name  string  `json:"name" db:"name"`
	Users []int64 `json:"users" db:"users"`
}

type ProjectUser struct {
	ProjectId int64 `json:"project_id"`
	UserId    int64 `json:"user_id" db:"user_id"`
}

type ProjectSocketSchema struct {
	ProjectId int64           `json:"project_id,omitempty"`
	Name      string          `json:"name"`
	Tasks     map[int64]*Task `json:" tasks"`
	Users     []int64         `json:"users" db:"users"`
}

func (s *ProjectSelected) Raw() []byte {
	raw, _ := json.Marshal(s)
	return raw
}

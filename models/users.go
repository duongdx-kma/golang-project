package models

import (
	"time"
	// "github.com/go-playground/validator/v10"
)

type User struct {
	ID        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Address   string     `json:"address,omitempty" db:"address,omitempty"`
	Password  string     `json:"password,omitempty" db:"password,omitempty"`
	Age       uint8      `json:"age,omitempty" db:"age, omitempty"`
	IsAdmin   bool       `json:"is_admin" db:"is_admin, omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at,omitempty"`
}

type AuthSchema struct {
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
	Message string `json:"message"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateUserSchema struct {
	Name     string `json:"name" validate:"required,min=1,max=255" errormgs:"invalid user name"`
	Age      int64  `json:"age" validate:"gte=1,lte=120" errormgs:"invalid user age"`
	Password string `json:"password" validate:"required,min=6,max=255" errormgs:"invalid user password"`
	Address  string `json:"address,omitempty" validate:"min=1,max=255" errormgs:"invalid user addr"`
}

type UpdateUserSchema struct {
	Age      int64  `json:"age" validate:"gte=1,lte=120" errormgs:"invalid user age"`
	Password string `json:"password" validate:"required,min=6,max=255" errormgs:"invalid user password"`
	Address  string `json:"address,omitempty" validate:"min=1,max=255" errormgs:"invalid user addr"`
	IsAdmin  bool   `json:"is_admin" db:"is_admin, omitempty"`
}

type DeleteUserSchema struct {
	Id string `json:"id" validate:"required"`
}

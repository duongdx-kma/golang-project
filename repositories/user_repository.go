package repositories

import (
	"context"
	database "duongdx/example/initializers"
	"duongdx/example/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	Store(context context.Context, createUserSchema models.CreateUserSchema) (models.User, error)
	FindAll(context context.Context) ([]models.User, error)
	Detail(context context.Context, id string) (models.User, error)
	GetUserByName(context context.Context, name string) (models.User, error)
	Update(context context.Context, id string, updateSchema models.UpdateUserSchema) (models.User, error)
}

type UserRepository struct {
	SQL *database.SQL
}

func (db *UserRepository) Store(
	ctx context.Context,
	createUserSchema models.CreateUserSchema,
) (models.User, error) {
	// variables
	user := models.User{
		Name:    createUserSchema.Name,
		Age:     uint8(createUserSchema.Age),
		Address: createUserSchema.Address,
	}
	now := time.Now()

	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserSchema.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Hashing password failed")
	}

	statement := `INSERT INTO users(name, address, age, password, created_at, updated_at)
	VALUES (:name, :address, :age, :password, :created_at, :updated_at)`

	user.Password = string(hashedPassword)
	user.CreatedAt = &now
	user.UpdatedAt = &now

	log.Printf("%+v", user)

	result, err := db.SQL.DB.NamedExecContext(ctx, statement, user)

	if err != nil {
		log.Fatal("Insert data user failed", err)

		return models.User{}, err
	}

	// get last inserted user id
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal("get data just have been created is fail", err)
	}
	user.ID = lastId

	return user, nil
}

func (db *UserRepository) Update(
	ctx context.Context,
	id string,
	updateSchema models.UpdateUserSchema,
) (models.User, error) {
	// variables
	user := models.User{
		Age:     uint8(updateSchema.Age),
		Address: updateSchema.Address,
		IsAdmin: updateSchema.IsAdmin,
	}
	now := time.Now()

	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	statement := `UPDATE users SET 
		address	   = ?,
		age 	   = ?,
		is_admin   = ?,
		updated_at = ?
		WHERE id   = ?`

	user.UpdatedAt = &now

	_, err := db.SQL.DB.Exec(
		statement,
		user.Address,
		user.Age,
		user.IsAdmin,
		user.UpdatedAt,
		id,
	)

	if err != nil {
		log.Fatal("Update user failed", err)

		return models.User{}, err
	}

	return user, nil
}

func (db *UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	var users []models.User
	query := "SELECT * FROM users WHERE deleted_at IS NULL"
	err := db.SQL.DB.SelectContext(ctx, &users, query)

	if err != nil {
		log.Fatal(err)

		return users, err
	}

	return users, nil
}

func (db *UserRepository) Detail(ctx context.Context, id string) (models.User, error) {
	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	user := models.User{}
	query := `SELECT id, name, address, password, age, is_admin FROM users WHERE id=? AND deleted_at IS NULL`
	err := db.SQL.DB.Get(&user, query, id)

	if err != nil {
		log.Fatal(err)

		return user, err
	}

	return user, nil
}

func (db *UserRepository) GetUserByName(ctx context.Context, name string) (models.User, error) {
	// Open mysql connection
	db.SQL.Connect()
	// Close mysql connection
	defer db.SQL.Close()

	user := models.User{}
	query := `SELECT * FROM users WHERE name=?`
	err := db.SQL.DB.Get(&user, query, name)

	if err != nil {
		return user, err
	}

	return user, nil
}

package controllers

import (
	database "duongdx/example/initializers"
	"duongdx/example/models"
	"duongdx/example/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserRepository repositories.UserInterface
}

type JwtCustomClaims struct {
	models.User
	jwt.StandardClaims
}

func NewUserController(sql *database.SQL) *UserController {
	return &UserController{
		UserRepository: &repositories.UserRepository{
			SQL: sql,
		},
	}
}

func (u *UserController) Authentication(e echo.Context) error {
	user := models.User{}

	form := new(models.LoginRequest)

	if err := e.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("Authentication - Binding request failed %s", err))
	}

	user, err := u.UserRepository.GetUserByName(e.Request().Context(), (*form).UserName)

	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Username is incorrect !")
	}

	// Hashing the password with the default cost of 10
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte((*form).Password))

	if result != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Password is incorrect !")
	}

	// load secret key from .env
	secretKey := os.Getenv("JWT_SECRET")

	// Set custom claims
	claims := JwtCustomClaims{
		models.User{
			ID:      user.ID,
			Name:    user.Name,
			Age:     user.Age,
			Address: user.Address,
			IsAdmin: user.IsAdmin,
		},
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	token, err := tokenAlgorithm.SignedString([]byte(secretKey))
	log.Println("aaaaaaaaaaaaa", secretKey, []byte(secretKey))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, models.AuthSchema{
		Token:   token,
		IsAdmin: user.IsAdmin,
		Message: "Login successful",
	})
}

// Create user
func (u *UserController) Create(e echo.Context) (models.User, error) {
	user := models.User{}
	form := new(models.CreateUserSchema)

	if err := e.Bind(form); err != nil {
		return user, echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("Create - Binding request failed %s", err))
	}

	// Validate request
	validate := validator.New()
	err := validate.Struct(form)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)

		return user, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Validation error: %s", errors))
	}

	// check existing user
	user, err = u.UserRepository.GetUserByName(e.Request().Context(), (*form).Name)
	if user.ID != 0 {
		return user, echo.NewHTTPError(http.StatusUnprocessableEntity, "user already exists")
	}

	user, err = u.UserRepository.Store(e.Request().Context(), (*form))

	if err != nil {
		log.Fatal("Insert failed: ", err)
	}

	if user.Name == "" {
		return user, echo.NewHTTPError(http.StatusUnprocessableEntity, "Create user failed")
	}

	return user, err
}

// Get all users
func (u *UserController) FindAll(e echo.Context) ([]models.User, error) {
	user, err := u.UserRepository.FindAll(e.Request().Context())

	return user, err
}

// Find user by id
func (u *UserController) FindById(e echo.Context, id string) (models.User, error) {
	user, err := u.UserRepository.Detail(e.Request().Context(), id)

	return user, err
}

// Update user
func (u *UserController) Update(e echo.Context, id string) (models.User, error) {
	user := models.User{}
	form := new(models.UpdateUserSchema)

	if err := e.Bind(form); err != nil {
		return user, echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("Update - Binding request failed %s", err))
	}

	user, err := u.UserRepository.Update(e.Request().Context(), id, *form)

	return user, err
}

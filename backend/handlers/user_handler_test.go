package handlers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nrmadi02/go_react_monorepo/backend/handlers"
	"github.com/nrmadi02/go_react_monorepo/backend/models"
	"github.com/nrmadi02/go_react_monorepo/backend/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// mockUserRepository implements the same methods as UserRepository for error simulation
type mockUserRepository struct {
	CreateUserFunc  func(user *models.User) error
	GetUsersFunc    func() ([]models.User, error)
	GetUserByIDFunc func(id uint) (models.User, error)
	UpdateUserFunc  func(user *models.User) error
	DeleteUserFunc  func(id uint) error
}

func (m *mockUserRepository) CreateUser(user *models.User) error {
	return m.CreateUserFunc(user)
}
func (m *mockUserRepository) GetUsers() ([]models.User, error) {
	return m.GetUsersFunc()
}
func (m *mockUserRepository) GetUserByID(id uint) (models.User, error) {
	return m.GetUserByIDFunc(id)
}
func (m *mockUserRepository) UpdateUser(user *models.User) error {
	return m.UpdateUserFunc(user)
}
func (m *mockUserRepository) DeleteUser(id uint) error {
	return m.DeleteUserFunc(id)
}

func setupTestApp(t *testing.T) (*fiber.App, *repository.UserRepository) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.User{})
	repo := repository.NewUserRepository(db)
	h := handlers.NewUserHandler(repo)
	app := fiber.New()
	app.Post("/users", h.CreateUser)
	app.Get("/users", h.GetUsers)
	app.Get("/users/:id", h.GetUser)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
	return app, repo
}

func TestCreateUserHandler_DBError(t *testing.T) {
	app := fiber.New()
	h := handlers.NewUserHandler(&mockUserRepository{
		CreateUserFunc: func(user *models.User) error {
			return errors.New("db error")
		},
	})
	app.Post("/users", h.CreateUser)
	payload := `{"name":"Test User","email":"test@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestCreateUserHandler(t *testing.T) {
	app, _ := setupTestApp(t)
	payload := `{"name":"Test User","email":"test@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestCreateUserHandler_InvalidBody(t *testing.T) {
	app, _ := setupTestApp(t)
	payload := "not a json"
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestCreateUserHandler_ValidationError(t *testing.T) {
	app, _ := setupTestApp(t)
	payload := `{"name":"","email":"salah"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestGetUsersHandler_DBError(t *testing.T) {
	app := fiber.New()
	h := handlers.NewUserHandler(&mockUserRepository{
		GetUsersFunc: func() ([]models.User, error) {
			return nil, errors.New("db error")
		},
	})
	app.Get("/users", h.GetUsers)
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestGetUsersHandler(t *testing.T) {
	app, repo := setupTestApp(t)
	repo.CreateUser(&models.User{Name: "User1", Email: "user1@example.com"})
	repo.CreateUser(&models.User{Name: "User2", Email: "user2@example.com"})
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.True(t, result["status"].(bool))
}

func TestGetUserHandler(t *testing.T) {
	app, repo := setupTestApp(t)
	user := &models.User{Name: "User1", Email: "user1@example.com"}
	repo.CreateUser(user)
	url := fmt.Sprintf("/users/%d", user.ID)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestGetUserHandler_InvalidID(t *testing.T) {
	app, _ := setupTestApp(t)
	req := httptest.NewRequest(http.MethodGet, "/users/abc", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestGetUserHandler_NotFound(t *testing.T) {
	app, _ := setupTestApp(t)
	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestUpdateUserHandler_DBError(t *testing.T) {
	app := fiber.New()
	h := handlers.NewUserHandler(&mockUserRepository{
		GetUserByIDFunc: func(id uint) (models.User, error) {
			return models.User{ID: id, Name: "User1", Email: "user1@example.com"}, nil
		},
		UpdateUserFunc: func(user *models.User) error {
			return errors.New("db error")
		},
	})
	app.Put("/users/:id", h.UpdateUser)
	payload := `{"name":"Updated User","email":"updated@example.com"}`
	req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestUpdateUserHandler(t *testing.T) {
	app, repo := setupTestApp(t)
	user := &models.User{Name: "User1", Email: "user1@example.com"}
	repo.CreateUser(user)
	url := fmt.Sprintf("/users/%d", user.ID)
	payload := `{"name":"Updated User","email":"updated@example.com"}`
	req := httptest.NewRequest(http.MethodPut, url, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestUpdateUserHandler_InvalidID(t *testing.T) {
	app, _ := setupTestApp(t)
	payload := `{"name":"Test","email":"test@example.com"}`
	req := httptest.NewRequest(http.MethodPut, "/users/abc", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestUpdateUserHandler_NotFound(t *testing.T) {
	app, _ := setupTestApp(t)
	payload := `{"name":"Test","email":"test@example.com"}`
	req := httptest.NewRequest(http.MethodPut, "/users/999", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestUpdateUserHandler_InvalidBody(t *testing.T) {
	app, repo := setupTestApp(t)
	user := &models.User{Name: "User1", Email: "user1@example.com"}
	repo.CreateUser(user)
	url := fmt.Sprintf("/users/%d", user.ID)
	req := httptest.NewRequest(http.MethodPut, url, strings.NewReader("notjson"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestUpdateUserHandler_ValidationError(t *testing.T) {
	app, repo := setupTestApp(t)
	user := &models.User{Name: "User1", Email: "user1@example.com"}
	repo.CreateUser(user)
	url := fmt.Sprintf("/users/%d", user.ID)
	payload := `{"name":"","email":"salah"}`
	req := httptest.NewRequest(http.MethodPut, url, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestDeleteUserHandler_DBError(t *testing.T) {
	app := fiber.New()
	h := handlers.NewUserHandler(&mockUserRepository{
		DeleteUserFunc: func(id uint) error {
			return errors.New("db error")
		},
	})
	app.Delete("/users/:id", h.DeleteUser)
	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestDeleteUserHandler(t *testing.T) {
	app, repo := setupTestApp(t)
	user := &models.User{Name: "User1", Email: "user1@example.com"}
	repo.CreateUser(user)
	url := fmt.Sprintf("/users/%d", user.ID)
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
}

func TestDeleteUserHandler_InvalidID(t *testing.T) {
	app, _ := setupTestApp(t)
	req := httptest.NewRequest(http.MethodDelete, "/users/abc", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestDeleteUserHandler_NotFound(t *testing.T) {
	app, _ := setupTestApp(t)
	req := httptest.NewRequest(http.MethodDelete, "/users/999", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

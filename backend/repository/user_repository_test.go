package repository_test

import (
	"testing"
	"github.com/nrmadi02/go_react_monorepo/backend/models"
	"github.com/nrmadi02/go_react_monorepo/backend/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.User{})
	return db
}

func TestCreateAndGetUser(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)
	user := &models.User{Name: "Test User", Email: "test@example.com"}
	err := repo.CreateUser(user)
	assert.NoError(t, err)
	users, err := repo.GetUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "Test User", users[0].Name)
}

func TestGetUserByID(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)
	user := &models.User{Name: "Test User", Email: "test@example.com"}
	repo.CreateUser(user)
	fetched, err := repo.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, fetched.Email)
}

func TestUpdateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)
	user := &models.User{Name: "Test User", Email: "test@example.com"}
	repo.CreateUser(user)
	user.Name = "Updated Name"
	err := repo.UpdateUser(user)
	assert.NoError(t, err)
	fetched, _ := repo.GetUserByID(user.ID)
	assert.Equal(t, "Updated Name", fetched.Name)
}

func TestDeleteUser(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)
	user := &models.User{Name: "Test User", Email: "test@example.com"}
	repo.CreateUser(user)
	err := repo.DeleteUser(user.ID)
	assert.NoError(t, err)
	users, _ := repo.GetUsers()
	assert.Len(t, users, 0)
}

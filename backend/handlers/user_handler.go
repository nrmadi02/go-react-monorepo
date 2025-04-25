package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nrmadi02/go_react_monorepo/backend/dto"
	"github.com/nrmadi02/go_react_monorepo/backend/models"
	"github.com/nrmadi02/go_react_monorepo/backend/repository"
	"gorm.io/gorm"
)

type UserHandler struct {
	Repo repository.UserRepositoryIface
}

func NewUserHandler(repo repository.UserRepositoryIface) *UserHandler {
	return &UserHandler{Repo: repo}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User Data"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
			Errors:  []string{err.Error()},
		})
	}
	// VALIDASI INPUT
	errs := dto.ValidateRequest(req)
	if len(errs) > 0 {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Validation failed",
			Errors:  errs,
		})
	}
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := h.Repo.CreateUser(user); err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Failed to create user",
			Errors:  []string{err.Error()},
		})
	}
	return c.Status(201).JSON(dto.SuccessResponse{
		Status:  true,
		Message: "User created successfully",
		Data: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} dto.UserResponse
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.Repo.GetUsers()
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch users",
			Errors:  []string{err.Error()},
		})
	}
	return c.JSON(dto.SuccessResponse{
		Status:  true,
		Message: "Users fetched successfully",
		Data: func() []dto.UserResponse {
			resp := make([]dto.UserResponse, len(users))
			for i, u := range users {
				resp[i] = dto.UserResponse{
					ID:        u.ID,
					Name:      u.Name,
					Email:     u.Email,
					CreatedAt: u.CreatedAt,
					UpdatedAt: u.UpdatedAt,
				}
			}
			return resp
		}(),
	})
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Invalid user ID",
			Errors:  []string{err.Error()},
		})
	}
	user, err := h.Repo.GetUserByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "User not found",
			Errors:  []string{err.Error()},
		})
	}
	return c.JSON(dto.SuccessResponse{
		Status:  true,
		Message: "User fetched successfully",
		Data: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update details of a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body dto.CreateUserRequest true "User Data"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Invalid user ID",
			Errors:  []string{err.Error()},
		})
	}
	user, err := h.Repo.GetUserByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "User not found",
			Errors:  []string{err.Error()},
		})
	}
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
			Errors:  []string{err.Error()},
		})
	}
	// VALIDASI INPUT
	errs := dto.ValidateRequest(req)
	if len(errs) > 0 {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Validation failed",
			Errors:  errs,
		})
	}
	user.Name = req.Name
	user.Email = req.Email
	if err := h.Repo.UpdateUser(&user); err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Failed to update user",
			Errors:  []string{err.Error()},
		})
	}
	return c.JSON(dto.SuccessResponse{
		Status:  true,
		Message: "User updated successfully",
		Data: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Delete a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Invalid user ID",
			Errors:  []string{err.Error()},
		})
	}
	err = h.Repo.DeleteUser(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(dto.ErrorResponse{
				Status:  false,
				Message: "User not found",
				Errors:  []string{err.Error()},
			})
		}
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Failed to delete user",
			Errors:  []string{err.Error()},
		})
	}
	return c.SendStatus(204)
}

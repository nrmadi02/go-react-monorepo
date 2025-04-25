package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nrmadi02/go_react_monorepo/backend/dto"
	"github.com/nrmadi02/go_react_monorepo/backend/repository"
	"github.com/nrmadi02/go_react_monorepo/backend/utils"
)

type AuthHandler struct {
	AccountRepo repository.AccountRepositoryIface
}

func NewAuthHandler(accountRepo repository.AccountRepositoryIface) *AuthHandler {
	return &AuthHandler{AccountRepo: accountRepo}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and get access token
// @Tags auth
// @Accept json
// @Produce json
// @Param data body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.SuccessResponse{data=dto.AuthResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
			Errors:  []string{err.Error()},
		})
	}
	errs := dto.ValidateRequest(req)
	if len(errs) > 0 {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Validation failed",
			Errors:  errs,
		})
	}
	account, err := h.AccountRepo.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Invalid credentials",
			Errors:  []string{err.Error()},
		})
	}
	token, err := utils.GenerateJWT(account.User.ID)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Failed to generate token",
			Errors:  []string{err.Error()},
		})
	}
	h.AccountRepo.UpdateAccessToken(account.ID, token)
	account.AccessToken = token
	resp := dto.NewAuthResponse(account)
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  true,
		Message: "Login successful",
		Data:    resp,
	})
}

// Register godoc
// @Summary Register user
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param data body dto.RegisterRequest true "Register Request"
// @Success 200 {object} dto.SuccessResponse{data=dto.AuthResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
			Errors:  []string{err.Error()},
		})
	}
	errs := dto.ValidateRequest(req)
	if len(errs) > 0 {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Validation failed",
			Errors:  errs,
		})
	}
	account, err := h.AccountRepo.Register(req.Email, req.Password, req.Name)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Failed to register",
			Errors:  []string{err.Error()},
		})
	}

	token, err := utils.GenerateJWT(account.User.ID)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Failed to generate token",
			Errors:  []string{err.Error()},
		})
	}
	h.AccountRepo.UpdateAccessToken(account.ID, token)
	account.AccessToken = token
	resp := dto.NewAuthResponse(account)
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  true,
		Message: "Register successful",
		Data:    resp,
	})
}

// Me godoc
// @Summary Get current user info
// @Description Get profile of the authenticated user (ID dari token)
// @Tags auth
// @Produce json
// @Success 200 {object} dto.SuccessResponse{data=dto.MeResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /auth/me [get]
func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return c.Status(401).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "Unauthorized",
			Errors:  []string{"Unauthorized"},
		})
	}
	account, err := h.AccountRepo.Me(userID)
	if err != nil {
		return c.Status(404).JSON(dto.ErrorResponse{
			Status:  false,
			Message: "User not found",
			Errors:  []string{err.Error()},
		})
	}
	resp := dto.NewMeResponse(account)
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  true,
		Message: "User fetched successfully",
		Data:    resp,
	})
}

package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nrmadi02/go_react_monorepo/backend/dto"
	"github.com/nrmadi02/go_react_monorepo/backend/utils"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Status:  false,
				Message: "Missing or invalid Authorization header",
				Errors:  []string{"Missing or invalid Authorization header"},
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Status:  false,
				Message: "Invalid or expired token",
				Errors:  []string{"Invalid or expired token"},
			})
		}

		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}

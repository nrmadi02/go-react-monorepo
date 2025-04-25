package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nrmadi02/go_react_monorepo/backend/handlers"
	"github.com/nrmadi02/go_react_monorepo/backend/middleware"
)

func RegisterAuthRoutes(app *fiber.App, handler *handlers.AuthHandler) {
	auth := app.Group("/auth")

	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
	auth.Get("/me", middleware.JWTProtected(), handler.Me)
}

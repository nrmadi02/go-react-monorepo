package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nrmadi02/go_react_monorepo/backend/handlers"
)

func RegisterUserRoutes(app *fiber.App, handler *handlers.UserHandler) {
	user := app.Group("/users")
	user.Post("/", handler.CreateUser)
	user.Get("/", handler.GetUsers)
	user.Get("/:id", handler.GetUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)
}

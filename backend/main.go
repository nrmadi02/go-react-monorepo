package main

import (
	"github.com/nrmadi02/go_react_monorepo/backend/config"
	"github.com/nrmadi02/go_react_monorepo/backend/handlers"
	"github.com/nrmadi02/go_react_monorepo/backend/repository"
	"github.com/nrmadi02/go_react_monorepo/backend/routes"

	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
	_ "github.com/nrmadi02/go_react_monorepo/backend/docs"
)

func main() {
	app := fiber.New()
	swagger := swagger.New(swagger.Config{
		Title: "Go React Monorepo API",
	})

	config.ConnectDatabase()
	config.MigrateDatabase()

	repo := repository.NewUserRepository(config.DB)
	handler := handlers.NewUserHandler(repo)

	routes.RegisterUserRoutes(app, handler)

	app.Get("/swagger/*", swagger)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":8000"))
}

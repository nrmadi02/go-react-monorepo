package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nrmadi02/go_react_monorepo/backend/config"
	"github.com/nrmadi02/go_react_monorepo/backend/handlers"
	"github.com/nrmadi02/go_react_monorepo/backend/repository"
	"github.com/nrmadi02/go_react_monorepo/backend/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/swagger"
	_ "github.com/nrmadi02/go_react_monorepo/backend/docs"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Makassar",
	}))

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

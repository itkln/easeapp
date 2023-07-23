package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itkln/ease-backend/app/handlers"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("/api")
	route.Post("/register", handlers.Register)
	route.Post("/login", handlers.Login)
	route.Get("/profile", handlers.Profile)
}

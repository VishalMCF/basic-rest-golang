package routes

import (
	"educative-rest-api-course/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// get items
	app.Get("/api/v1/items", handlers.GetAllItems)

	// get item by id
	app.Get("/api/v1/items/:id", handlers.GetItemById)

	// create an item
	app.Post("/api/v1/items", handlers.CreateItem)

	// update an item
	app.Put("/api/v1/items/:id", handlers.UpdateItem)

	// delete an item
	app.Delete("/api/v1/items/:id", handlers.DeleteItem)
}

func SetupAuthRoutes(app *fiber.App) {

	// signup handler for the signup request
	app.Post("/api/v1/signup", handlers.SignUp)

	// login handler for the login request
	app.Post("/api/v1/login", handlers.Login)
}

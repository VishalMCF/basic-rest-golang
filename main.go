package main

import (
	"educative-rest-api-course/database"
	"educative-rest-api-course/routes"
	"educative-rest-api-course/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

const DEFAULT_PORT = "8090"

func newFibreApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello000!")
	})

	// registering the inventory routes
	routes.SetupRoutes(app)

	// registering the authentication routes
	routes.SetupAuthRoutes(app)

	return app
}

func main() {
	app := newFibreApp()

	database.InitDatabase(utils.GetValue("DB_NAME"))

	var port string = os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	// start the port at the application 8090 port
	app.Listen(fmt.Sprintf(":%s", port))
}

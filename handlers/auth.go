package handlers

import (
	"educative-rest-api-course/models"
	services "educative-rest-api-course/service"
	"github.com/gofiber/fiber/v2"
)

// SignUp returns the JWT token
func SignUp(c *fiber.Ctx) error {
	var userInput *models.UserRequest = new(models.UserRequest)

	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// validate the request
	errors := userInput.ValidateStruct()

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "Validation failed",
			Data:    errors,
		})
	}

	// perform the signup a.k.a store the user details into the database
	token, err := services.SignUp(*userInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "SignUp failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response[any]{
		Success: true,
		Message: "Successfully Signed Up",
		Data:    token,
	})
}

// Login returns the JWT token for the registered token only
func Login(c *fiber.Ctx) error {
	var userInput *models.UserRequest = new(models.UserRequest)
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "Request could not be parsed",
		})
	}

	errors := userInput.ValidateStruct()

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "Validation failed",
			Data:    errors,
		})
	}

	token, err := services.Login(*userInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "Login failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response[any]{
		Success: true,
		Message: "Successfully Logged In",
		Data:    token,
	})
}

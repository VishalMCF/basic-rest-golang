package handlers

import (
	"educative-rest-api-course/models"
	"educative-rest-api-course/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// get all items
func GetAllItems(c *fiber.Ctx) error {
	var items []models.Item
	items = services.GetAllItems()

	// return the response
	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "All items data",
		Data:    items,
	})
}

// get ite m by id
func GetItemById(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item
	item, err := services.GetItemById(id)

	// return the error response if err is present
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// return the item found
	return c.JSON(models.Response[any]{
		Success: true,
		Message: "Item found",
		Data:    item,
	})
}

// create an item
func CreateItem(c *fiber.Ctx) error {
	var itemInput = new(models.ItemRequest)

	// parse the request into itemInput variable
	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// validate the request
	errors := itemInput.ValidateStruct()

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	var createdItem = services.CreateItem(*itemInput)

	return c.JSON(models.Response[any]{
		Success: true,
		Message: "Item created",
		Data:    createdItem,
	})
}

// update an item
func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var itemInput = new(models.ItemRequest)
	// parse the request into itemInput variable
	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := itemInput.ValidateStruct()

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	var updatedItem, err = services.UpdateItem(*itemInput, id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[any]{
		Success: true,
		Message: "Item updated",
		Data:    updatedItem,
	})
}

// delete an item
// DeleteItem returns deletion result
func DeleteItem(c *fiber.Ctx) error {
	// get the item's ID from the request parameter
	var itemID string = c.Params("id")

	// delete the item data
	var result bool = services.DeleteItem(itemID)

	if result {
		// if successful, return the result
		return c.JSON(models.Response[any]{
			Success: true,
			Message: "item deleted",
		})
	}

	// return the failed result
	return c.Status(http.StatusNotFound).JSON(models.Response[any]{
		Success: false,
		Message: "item failed to delete",
	})
}

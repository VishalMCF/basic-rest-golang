package services

import (
	"educative-rest-api-course/models"
	"errors"
	"github.com/google/uuid"
	"time"
)

var storage []models.Item = []models.Item{}

// get all the items
func GetAllItems() []models.Item {
	return storage
}

// get item by ID
func GetItemById(id string) (models.Item, error) {
	for _, item := range storage {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Item{}, errors.New("Item could not be found")
}

// create a new item
func CreateItem(request models.ItemRequest) models.Item {
	var newItem models.Item = models.Item{
		ID:        uuid.New().String(),
		Name:      request.Name,
		Price:     request.Price,
		Quantity:  request.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	storage = append(storage, newItem)
	return newItem
}

// update item by id
func UpdateItem(request models.ItemRequest, id string) (models.Item, error) {
	for _, item := range storage {
		if item.ID == id {
			item.UpdatedAt = time.Now()
			item.Name = request.Name
			item.Price = request.Price
			item.Quantity = request.Quantity
			return item, nil
		}
	}
	return models.Item{}, errors.New("Item could not be updated or found")
}

// delete an item by id
func DeleteItem(id string) bool {
	var tempStorage []models.Item = []models.Item{}
	for _, item := range storage {
		if item.ID != id {
			tempStorage = append(tempStorage, item)
		}
	}
	storage = tempStorage
	return true
}

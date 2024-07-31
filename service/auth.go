package services

import (
	"educative-rest-api-course/database"
	"educative-rest-api-course/models"
	"educative-rest-api-course/utils"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// register the user during the sign up process and return the token
func SignUp(request models.UserRequest) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	var user models.User = models.User{
		ID:       uuid.New().String(),
		Email:    request.Email,
		Password: string(password),
	}

	database.DB.Create(&user)

	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(request models.UserRequest) (string, error) {
	var user models.User

	result := database.DB.Where(&user, "email = ?", request.Email).First(&user)

	if result.RowsAffected == 0 {
		return "", errors.New("User not found")
	}

	database.DB.Where("email = ?", user.Email).First(&models.User{})

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	// if the password do not match the return the err
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}

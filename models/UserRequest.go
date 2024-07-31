package models

import "github.com/go-playground/validator/v10"

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (userRequest *UserRequest) ValidateStruct() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(userRequest)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Field = err.Field()
			errors = append(errors, &element)
		}
	}

	return errors
}

package errors

import (
	"blogProject/internal/providers/validation"
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = make(map[string]string)
}

func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
			//["name": "the field is required"]
		}
	}
}

func Add(key string, value string) {
	errorsList[strings.ToLower(key)] = value
}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorsList
}

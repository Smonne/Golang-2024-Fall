package utils

import (
    "github.com/go-playground/validator/v10"
)

// Validator instance
var validate *validator.Validate

func init() {
    validate = validator.New()
}

// ValidateStruct validates a struct based on tags
func ValidateStruct(data interface{}) error {
    return validate.Struct(data)
}

package types

import (
	"github.com/maelsilva96/lambda_utils/entities"
)

type UserNameAlreadyExists struct {
	Message string
}

func (err *UserNameAlreadyExists) Error() string {
	return err.Message
}

type FieldError struct {
	fieldsWithError []entities.FieldError
}

func NewFieldError(errors []entities.FieldError) *FieldError {
	return &FieldError{
		fieldsWithError: errors,
	}
}

func (m *FieldError) Error() string {
	return "Fields with error!"
}

func (m *FieldError) GetErrors() []entities.FieldError {
	return m.fieldsWithError
}

type TokenNotFound struct{}

func NewTokenNotFound() *TokenNotFound {
	return &TokenNotFound{}
}

func (err *TokenNotFound) Error() string {
	return ""
}

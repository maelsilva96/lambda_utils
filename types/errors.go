package types

import (
	"fmt"
	"strings"
)

type UserNameAlreadyExists struct {
	Message string
}

func (err *UserNameAlreadyExists) Error() string {
	return err.Message
}

type FieldError struct {
	errorMessages map[string][]string
}

func NewFieldError(errors map[string][]string) *FieldError {
	return &FieldError{
		errorMessages: errors,
	}
}

func (m *FieldError) Error() string {
	message := ""
	for n, messages := range m.errorMessages {
		message += fmt.Sprintf("O campo (%s) tem os seguintes erros: %s\n", n, strings.Join(messages, ","))
	}
	return message
}

func (m *FieldError) GetErrors() map[string][]string {
	return m.errorMessages
}

type TokenNotFound struct{}

func NewTokenNotFound() *TokenNotFound {
	return &TokenNotFound{}
}

func (err *TokenNotFound) Error() string {
	return ""
}

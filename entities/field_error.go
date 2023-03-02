package entities

type ErrorCode int

const (
	None                ErrorCode = 0
	UnMappedCode        ErrorCode = 1
	NoAuthorizedCode    ErrorCode = 2
	ResetPasswordCode   ErrorCode = 3
	FieldErrorCode      ErrorCode = 4
	ChangePasswordCode  ErrorCode = 5
	ExpiredCode         ErrorCode = 6
	TokenNotFoundCode   ErrorCode = 7
	ExistingUser        ErrorCode = 8
	InvalidPasswordCode ErrorCode = 9
	UserNotFoundCode    ErrorCode = 10
)

type FieldError struct {
	Field        string `json:"field"`
	ErrorMessage string `json:"error_message"`
}

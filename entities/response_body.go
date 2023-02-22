package entities

type ResponseBody struct {
	Success    bool         `json:"success"`
	ErrorCode  ErrorCode    `json:"error_code"`
	Message    string       `json:"message"`
	Data       interface{}  `json:"data"`
	FieldError []FieldError `json:"field_error"`
}

func NewResponseBodyFieldError(fieldError []FieldError) *ResponseBody {
	return &ResponseBody{
		Success:    false,
		ErrorCode:  FieldErrorCode,
		FieldError: fieldError,
	}
}

package response

// Error Simple error response
func Error(code int, message string) Base {
	return Base{
		Status:     code,
		Validation: make(map[string]string),
		Data:       make(map[string]interface{}),
		Message:    message,
	}
}

// ErrorValidation Error response with validation error
func ErrorValidation(validation map[string]string, message string) Base {
	return Base{
		Status:     400,
		Validation: validation,
		Data:       make(map[string]interface{}),
		Message:    message,
	}
}

// Success Success response
func Success(i interface{}) Base {
	return Base{
		Status:     200,
		Validation: make(map[string]string),
		Data:       i,
		Message:    "success",
	}
}

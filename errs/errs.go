package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}
func NewAppError(code int, message string) AppError {
	return AppError{Code: http.StatusBadGateway, Message: message}
}
func NewValidationError(message string) AppError {
	return AppError{Code: 400, Message: message}
}

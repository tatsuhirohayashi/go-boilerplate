package apperrors

import "fmt"

type ErrorType string

const (
	NotFound          ErrorType = "NOT_FOUND"
	ValidationError   ErrorType = "VALIDATION_ERROR"
	PermissionDenied  ErrorType = "PERMISSION_DENIED"
	AlreadyExists     ErrorType = "ALREADY_EXISTS"
	Unauthorized      ErrorType = "UNAUTHORIZED"
	BusinessRuleError ErrorType = "BUSINESS_RULE_ERROR"
	InternalError     ErrorType = "INTERNAL_ERROR"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func NewValidationError(message string, err error) *AppError {
	return &AppError{
		Type:    ValidationError,
		Message: message,
		Err:     err,
	}
}

func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
		Err:     err,
	}
}

func NewUnauthorizedError(message string, err error) *AppError {
	return &AppError{
		Type:    Unauthorized,
		Message: message,
		Err:     err,
	}
}

func NewAlreadyExistsError(message string, err error) *AppError {
	return &AppError{
		Type:    AlreadyExists,
		Message: message,
		Err:     err,
	}
}

func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Type:    InternalError,
		Message: message,
		Err:     err,
	}
}

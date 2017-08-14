package models

import (
	"errors"
)

type HTTPError struct {
	error
	statusCode int
}

func NewHTTPError(status int, description string) *HTTPError {
	return &HTTPError{
		error:      errors.New(description),
		statusCode: status,
	}
}

func (e *HTTPError) StatusCode() int {
	return e.statusCode
}

package domain

import "fmt"

type AppError struct {
	msg       string
	exception error
}

// Error implements the Error method of the error interface
func (s *AppError) Error() string {
	return fmt.Sprintf("%s: exception: %v", s.msg, s.exception)
}

// Is implements the Is method of the errors.Is interface
func (s *AppError) Is(target error) bool {
	_, ok := target.(*AppError)
	return ok
}

func (s *AppError) Unwrap() error {
	return s.exception
}

func NewAppError(s string, err error) *AppError {
	return &AppError{
		msg:       s,
		exception: err,
	}
}

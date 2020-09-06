package database

import "fmt"

const (
	UniqueConstraintViolation = 1
)

type UserFriendlyDatabaseError struct {
	ErrorType      int
	FieldName      string
	AttemptedValue string
	UserFault      bool
	RootError      error `json:"-"`
}

func (e UserFriendlyDatabaseError) Error() string {
	switch e.ErrorType {
	case UniqueConstraintViolation:
		return fmt.Sprintf("There is already a %s with value %s", e.FieldName, e.AttemptedValue)
	default:
		return "Sorry, an error occurred during a database operation."
	}
}

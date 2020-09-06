package validation

import "fmt"

type Error struct {
	Field   string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("Validation failed for field %s: %s", e.Field, e.Message)
}

package validation

import (
	"github.com/GoosvandenBekerom/go-cms/pkg/pages"
	"strings"
)

func Validate(page pages.Page) *Error {
	if strings.HasPrefix(page.Path, "/api") {
		return &Error{
			Field:   "Path",
			Message: "Paths starting with /api are reserved.",
		}
	}
	return nil
}

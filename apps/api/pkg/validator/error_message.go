package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ErrorMessage(errs validator.ValidationErrors) (messages []string) {
	for _, e := range errs {
		switch e.Tag() {
		case "required":
			messages = append(messages, fmt.Sprintf("%s is required", e.Field()))
		case "email":
			messages = append(messages, "invalid Email")
		case "min":
			messages = append(messages, fmt.Sprintf("%s should have at least %s characters", e.Field(), e.Param()))
		default:
			messages = append(messages, fmt.Sprintf("%s invalid", e.Field()))
		}
	}
	return
}

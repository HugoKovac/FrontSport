package user

import (
	"GoNext/base/pkg/fiber/fibercontext"
	"GoNext/base/pkg/templ"
	customvalidator "GoNext/base/pkg/validator"
	"GoNext/base/templ/components"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	u := fibercontext.GetUserToContext(c)

	var userUpdate struct {
		Email       string `json:"email" validate:"email"`
		OldPassword string `json:"oldPassword" validate:"password"`
		NewPassword string `json:"newPassword" validate:"password"`
	}

	if err := c.BodyParser(&userUpdate); err != nil {
		log.Println(err)
		c.Status(fiber.StatusUnprocessableEntity)
		if errs, ok := err.(validator.ValidationErrors); ok {
			return templ.Render(c, components.ErrorMessage(customvalidator.ErrorMessage(errs)))
		} else {
			return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
		}
	}

	if err := h.Validate.Struct(userUpdate); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.Status(fiber.StatusUnprocessableEntity)
			return templ.Render(c, components.ErrorMessage(customvalidator.ErrorMessage(errs)))
		}
	}

	_, err := h.UserService.Update(u.Id, userUpdate.Email, userUpdate.OldPassword, userUpdate.NewPassword)
	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return templ.Render(c, components.ErrorMessage([]string{err.Error()}))
	}

	templ.Render(c, components.Toast(components.ToastAttributes{
		T:       "success",
		Message: "Profile updated",
	}))

	return nil
}

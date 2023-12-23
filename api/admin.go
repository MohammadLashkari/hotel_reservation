package api

import (
	"errors"
	"hotel-reservation/models"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*models.User)
	if !ok {
		return errors.New("unauthorized")
	}
	if !user.Admin {
		return errors.New("unauthorized")
	}
	return c.Next()
}

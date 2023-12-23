package api

import (
	"errors"
	"hotel-reservation/models"

	"github.com/gofiber/fiber/v2"
)

func getAuthUser(c *fiber.Ctx) (*models.User, error) {
	user, ok := c.Context().UserValue("user").(*models.User)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	return user, nil
}

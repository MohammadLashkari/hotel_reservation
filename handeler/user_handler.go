package handeler

import (
	"hotel-reservation/models"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := models.User{
		FirstName: "ali",
		LastName:  "lashkari",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("user1")
}

package handeler

import (
	"context"
	"hotel-reservation/db"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		UserStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("")
}

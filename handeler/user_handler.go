package handeler

import (
	"errors"
	"hotel-reservation/db"
	"hotel-reservation/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
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
	id := c.Params("id")
	user, err := h.GetUserById(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found"})
		}
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params models.UserPrams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errs := params.Validate(); len(errs) > 0 {
		return c.JSON(errs)
	}
	user, err := models.NewUserFromPrams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.DeleteUser(c.Context(), id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": id})
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.DeleteUser(c.Context(), id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": id})
}

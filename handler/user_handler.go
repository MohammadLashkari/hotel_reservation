package handler

import (
	"errors"
	"hotel-reservation/db"
	"hotel-reservation/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	*db.Store
}

func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{
		Store: store,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.UserStore.GetById(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found"})
		}
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.UserStore.GetAll(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params models.InsertUserPrams
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
	insertedUser, err := h.UserStore.Insert(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Delete(c.Context(), id); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": id})
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		params models.UpdateUserPrams
		id     = c.Params("id")
	)
	objecId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&params); err != nil {
		return nil
	}
	filter := bson.M{"_id": objecId}
	if err := h.UserStore.Update(c.Context(), filter, params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": id})
}

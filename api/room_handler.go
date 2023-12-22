package api

import (
	"hotel-reservation/db"

	"github.com/gofiber/fiber/v2"
)

type RoomHandler struct {
	*db.Store
}

func NewRoomHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		Store: store,
	}
}

func (h *BookingHandler) HandleGetRooms(c *fiber.Ctx) error {
	bookedRooms, err := h.RoomStore.GetAll(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(bookedRooms)
}

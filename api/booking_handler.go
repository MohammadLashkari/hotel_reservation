package api

import (
	"context"
	"hotel-reservation/db"
	"hotel-reservation/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingHandler struct {
	*db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		Store: store,
	}
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	booking, err := h.BookingStore.GetAll(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(booking)
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.BookingStore.GetById(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(booking)
}

func (h *BookingHandler) isRoomAvailable(ctx context.Context, roomId primitive.ObjectID, params models.BookRoomParams) (bool, error) {
	filter := bson.M{
		"roomId": roomId,
		"fromDate": bson.M{
			"$gte": params.FromDate,
		},
		"tillDate": bson.M{
			"$gte": params.TillDate,
		},
	}
	res, err := h.BookingStore.GetAll(ctx, filter)
	if err != nil {
		return false, err
	}
	ok := len(res) == 0
	return ok, nil

}

func (h *BookingHandler) HandleBookRoom(c *fiber.Ctx) error {
	var (
		params models.BookRoomParams
		id     = c.Params("id")
	)
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	// validate fromDate and tilDate
	if err := params.Validate(); err != nil {
		return err
	}
	roomId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// check the room already booded or not
	ok, err := h.isRoomAvailable(c.Context(), roomId, params)
	if err != nil {
		return err
	}
	if !ok {
		return c.JSON("room already booked")
	}
	user, ok := c.Context().Value("user").(*models.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON("lol")
	}
	booking := models.Booking{
		UserId:     user.Id,
		RoomId:     roomId,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
		NumPersons: params.NumPersons,
	}
	inserted, err := h.BookingStore.Insert(c.Context(), &booking)
	if err != nil {
		return err
	}
	return c.JSON(inserted)
}

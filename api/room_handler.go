package api

import (
	"errors"
	"hotel-reservation/db"
	"hotel-reservation/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	*db.Store
}

type BookRoomParams struct {
	FromDate   time.Time `bson:"fromDate" json:"fromDate"`
	TillDate   time.Time `bson:"tillDate" json:"tillDate"`
	NumPersons int       `bson:"numPersons" json:"NumPersons"`
}

func (p BookRoomParams) validate() error {
	if time.Now().After(p.FromDate) || time.Now().After(p.TillDate) {
		return errors.New("cannot book a room in the past")
	}
	return nil
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		Store: store,
	}
}

func (h *RoomHandler) HandleBookedRooms(c *fiber.Ctx) error {
	bookedRooms, err := h.BookingStore.GetAll(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(bookedRooms)
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var (
		param BookRoomParams
		id    = c.Params("id")
	)
	if err := c.BodyParser(&param); err != nil {
		return err
	}
	if err := param.validate(); err != nil {
		return err
	}
	if res, _ := h.RoomStore.GetById(c.Context(), id); res != nil {
		return c.JSON("room already has been booked")
	}
	roomId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	user, ok := c.Context().Value("user").(*models.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON("lol")
	}
	booking := models.Booking{
		UserId:     user.Id,
		RoomId:     roomId,
		FromDate:   param.FromDate,
		TillDate:   param.TillDate,
		NumPersons: param.NumPersons,
	}
	inserted, err := h.BookingStore.Insert(c.Context(), &booking)
	if err != nil {
		return err
	}
	return c.JSON(inserted)
}

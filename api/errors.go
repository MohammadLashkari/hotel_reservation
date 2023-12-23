package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.StatusCode).JSON(apiError)
	}
	apiErr := Error{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}
	return c.Status(apiErr.StatusCode).JSON(apiErr)
}

type Error struct {
	StatusCode int    `json:"code"`
	Message    string `json:"error"`
}

func (e Error) Error() string {
	return e.Message
}

func ErrUnAuthorized() Error {
	return Error{
		StatusCode: http.StatusUnauthorized,
		Message:    "unauthorized request",
	}
}

func ErrInvalidId() Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		Message:    "invalid id given",
	}
}

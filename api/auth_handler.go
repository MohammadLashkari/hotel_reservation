package api

import (
	"errors"
	"fmt"
	"hotel-reservation/db"
	"hotel-reservation/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	*db.Store
}

func NewAuthHandler(store *db.Store) *AuthHandler {
	return &AuthHandler{
		Store: store,
	}
}

type AuthPrams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

func (h *AuthHandler) HandleAuth(c *fiber.Ctx) error {
	var params AuthPrams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	user, err := h.UserStore.GetByEmail(c.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("invalid credentials")
		}
		return err
	}

	if !models.IsPasswordValid(user.EncryptedPassword, params.Password) {
		return errors.New("invalid credentials")

	}
	resp := AuthResponse{
		User:  user,
		Token: createTokenFromUser(user),
	}
	return c.JSON(resp)

}

func createTokenFromUser(user *models.User) string {
	expire := time.Now().Add(time.Hour * 4).Unix()
	claim := jwt.MapClaims{
		"id":     user.Id,
		"email":  user.Email,
		"expire": expire,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("failed to sign token with secret", err)
	}
	return tokenString
}

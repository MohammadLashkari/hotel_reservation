package middleware

import (
	"errors"
	"fmt"
	"hotel-reservation/db"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthenticaion(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok {
			return errors.New("unauthorized")
		}
		claims, err := validateToken(token[0])
		if err != nil {
			return err
		}
		expireFloat := claims["expire"].(float64)
		expire := int64(expireFloat)
		if time.Now().Unix() > expire {
			return errors.New("token expired")
		}
		userId := claims["id"].(string)
		user, err := userStore.GetById(c.Context(), userId)
		if err != nil {
			return errors.New("unauthorized")
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", t.Header["alg"])
			return nil, errors.New("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET_KEY")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("faild to parse JWT token :", err)
		return nil, errors.New("unauthorized")
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, errors.New("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	return claims, nil
}

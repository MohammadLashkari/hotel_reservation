package main

import (
	"flag"
	"hotel-reservation/handeler"

	"github.com/gofiber/fiber/v2"
)

var portAddr string

func init() {
	flag.StringVar(&portAddr, "port", ":8080", "port address")
	flag.Parse()
}

func main() {

	app := fiber.New()
	app.Get("/user", handeler.HandleGetUsers)
	app.Listen(portAddr)

}

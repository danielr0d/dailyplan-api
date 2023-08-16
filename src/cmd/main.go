package main

import (
	"github.com/gofiber/fiber/v2"
)

func HelloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func main() {
	app := fiber.New()

	app.Get("/", HelloHandler)

	app.Listen(":3000")
}

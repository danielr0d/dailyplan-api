package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func HelloHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": fmt.Sprintf("Hello, %s", c.Params("name"))})
}

func main() {
	app := fiber.New()

	app.Get("/:name", HelloHandler)

	app.Listen(":3000")

}

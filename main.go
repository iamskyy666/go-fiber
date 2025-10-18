package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


func main() {
	// Send, SendStatus, SendFile & Status
	// Status sets the HTTP status for the response. This method is chainable.
	app := fiber.New()
	
	app.Get("/fiber", func(c *fiber.Ctx) error {
  c.Status(fiber.StatusOK)
  return nil
})

app.Get("/hello", func(c *fiber.Ctx) error {
  return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
})

app.Get("/world", func(c *fiber.Ctx) error {
  return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
})
	
	fmt.Println("Server Up N Running.. ✅")
	log.Fatal(app.Listen(":3000"))
}
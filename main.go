package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()


// forcefully destroying server / graceful shutdown

app.Get("/api",func (c *fiber.Ctx)error  {
	if c.IP() != "127.0.0.1"{
		return app.Shutdown()
	}
	return c.SendString("Correct IP Address connected.. ✅")
})

    log.Fatal(app.Listen(":3000"))

	fmt.Println("Server Up N Running.. ✅")
}
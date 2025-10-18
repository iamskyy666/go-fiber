package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


func main() {
	// Send, SendStatus, SendFile & Status
	// SendFile transfers the file from the given path. The file is not compressed by default, enable this by passing a 'true' argument Sets the Content-Type response HTTP header field based on the filenames extension.
	app := fiber.New()
	
	app.Get("/",func (ctx *fiber.Ctx) error {
		
		return ctx.SendFile("./files/sample.txt")
	})

	app.Get("/not-found",func (ctx *fiber.Ctx) error {
		
		return ctx.SendFile("./files/404.html")
	})
	fmt.Println("Server Up N Running.. ✅")
	log.Fatal(app.Listen(":3000"))
}
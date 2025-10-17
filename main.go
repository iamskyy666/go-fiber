package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// download

	app := fiber.New()
	app.Get("/download", func (ctx *fiber.Ctx)error  {
		return ctx.Download("./files/sample.txt")
	})


	fmt.Println("Server Up N Running.. ✅")
	log.Fatal(app.Listen(":3000"))
}
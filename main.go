package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()
	
app.Post("/upload-document", func(ctx *fiber.Ctx) error {
	// FormFile returns the first file by key from a MultipartForm.
	// SaveFile saves any multipart file to disk.
	
	file,err:=ctx.FormFile("document")
	if err!=nil{
		return err
	}

	return ctx.SaveFile(file,fmt.Sprintf("./%s",file.Filename))


})
	fmt.Println("Server Up N Running.. ✅")
	log.Fatal(app.Listen(":3000"))
}

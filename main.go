package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()
	
app.Post("/upload-doc", func(c *fiber.Ctx) error {
  // Parse the multipart form:
  if form, err := c.MultipartForm(); err == nil {
    // => *multipart.Form

    // Get all files from "documents" key:
    files := form.File["documents"]
    // => []*multipart.FileHeader

    // Loop through files:
    for _, file := range files {
      //name:=fmt.Sprintf(file.Filename, file.Size, time.Now().String())
      // => "tutorial.pdf" 360641 "application/pdf"

      // Save the files to disk:
      if err := c.SaveFile(file, fmt.Sprintf("./%s", time.Now().String())); err != nil {
        return err
      }
    }
    return err
  }
  return c.SendStatus(fiber.StatusOK)
})
	fmt.Println("Server Up N Running.. ✅")
	log.Fatal(app.Listen(":3000"))
}

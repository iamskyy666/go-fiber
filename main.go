package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main(){
app:= InitApp()
log.Fatal(app.Listen(":3000"))
fmt.Println("Server Up N Running.. ✅")
}


func InitApp() *fiber.App{
    app := fiber.New()


// testing

app.Get("/check-token",func (c *fiber.Ctx)error  {
	token:=c.Get("Authorization")
	if token == ""{
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.SendStatus(fiber.StatusOK)
})

	return app
}
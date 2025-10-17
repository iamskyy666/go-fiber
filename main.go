package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// get cookies 🍪
	app := fiber.New()
	app.Get("/login", func (ctx *fiber.Ctx)error  {
		// create cookie
		cookieAuth := new(fiber.Cookie)
		cookieAuth.Name = "username"
		cookieAuth.Value = "john doe"
		cookieAuth.Expires = time.Now().Add(24 * time.Hour)

		cookieTheme:= new(fiber.Cookie)
		cookieTheme.Name = "app_theme"
		cookieTheme.Value = "Dark"

		// set cookie
		ctx.Cookie(cookieAuth)
		ctx.Cookie(cookieTheme)

		return ctx.SendStatus(fiber.StatusOK)
		
	})

	// Read Cookies 
	app.Get("/checkout", func (ctx *fiber.Ctx)error  {
		fmt.Println("username:",ctx.Cookies("username","no user found!"))
		fmt.Println("app_theme:",ctx.Cookies("app_theme","light"))
		return ctx.SendStatus(fiber.StatusOK)
		
	})

	// Clear Cookies
	app.Get("/logout", func (ctx *fiber.Ctx)error  {
		ctx.ClearCookie() // clear all cookies
		// ctx.ClearCookie("app_theme") // clear specific cookies
			return ctx.SendStatus(fiber.StatusOK)
	})


	fmt.Println("Server Up N Running.. ✅")
	log.Fatal(app.Listen(":3000"))
}


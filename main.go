package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type QueryProduct struct{
	Sort string `query:"sort"`
	Date int `query:"date"`
	New bool `query:"new"`
	Maxprice float64 `query:"maxprice"`
}


func main() {
	// 💡 QueryParser : Binds the query string to a struct.
	app := fiber.New()
	// GET http://example.com/?sort=asc&date=today&new=true&maxprice=999.99

app.Get("/", func(ctx *fiber.Ctx) error {
	queryProduct:=new(QueryProduct)
	if err := ctx.QueryParser(queryProduct); err!=nil{
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"sort": queryProduct.Sort,
		"date":queryProduct.Date,
		"new":queryProduct.New,
		"maxprice":queryProduct.Maxprice,
	})

})
	fmt.Println("Server Up N Running.. ✅")
	log.Fatal(app.Listen(":3000"))
}

// {
//   "date": 234567,
//   "maxprice": 999.99,
//   "new": true,
//   "sort": "asc"
// }
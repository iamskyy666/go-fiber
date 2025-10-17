package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Pokemon struct{
	Name string `json:"name"`
	Trainer string `json:"trainer"`
}

func main() {
	// bodyParser

	app := fiber.New()
	app.Post("/api/pokemon", func (ctx *fiber.Ctx)error  {
		seasonNum:=9
		p:=new(Pokemon)

		if err:=ctx.BodyParser(p); err!=nil{
			return err
		}

		// log.Println(p.Name)
		// log.Println(p.Trainer)

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"name":p.Name,
			"trainer":p.Trainer,
			"season":seasonNum,
		})
	})


	fmt.Println("Server Up N Running.. ✅")
	log.Fatal(app.Listen(":3000"))
}

// JSON-Body:
// {
//   "name": "Blaziken 🔥",
//   "season": 9,
//   "trainer": "May"
// }
package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
	"os"
	"fmt"
)

func main() {
    engine := pug.New("./views", ".pug")
    app := fiber.New(fiber.Config{
    Views: engine,
})

    app.Static("/", "./public") // nastaví složku "public" jako zdroj statických souborů, které bude server posílat pod cestou "/" 

    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
			"Title": "Tour de App",
		}) // vrátí views/index.pug, kde Title je nahrazen tím, co zde zadáme
    })

    app.Get("/api", func(c *fiber.Ctx) error {
        return c.SendString("Ahoj Světe") // vrátí čistý text
    })
    fmt.Print(os.Getenv("PORT"))
    app.Listen(os.Getenv("PORT"))
}

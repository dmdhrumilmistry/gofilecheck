package main

import (
	"encoding/json"
	"log"

	"github.com/dmdhrumilmistry/gofilecheck/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "gofilecheck",
		CaseSensitive: true,
		ServerHeader:  "gofilecheck",
		StrictRouting: true,
		Prefork:       false,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		XMLEncoder:    nil, // disable XML support
	})

	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(requestid.New())

	router.SetupRoutes(app)

	log.Fatalln(app.Listen(":3000"))
}

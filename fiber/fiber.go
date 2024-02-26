package fiber

import (
	"gallery/core"

	"github.com/gofiber/fiber/v2"
)

var (
	log = core.GetLogger()
)

func Start() {
	fiberConfig := fiber.Config{}
	app := fiber.New(fiberConfig)
	middleware(app)
	// auth.InitJWTAuthenticationBackend()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ayoooo")
	})

	app.Post("/generate", generate)
	app.Get("/getimage/:id", getImage)
	app.Get("/collection/:id", getCollection)
	app.Get("/collection/:id/add", AddToCollection)
	app.Delete("/collection/:id/:job", rmFromCollection)

	// app.Get("/login", login)

	// app.Get("/register", register)

	app.Listen(":3000")
}

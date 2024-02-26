package fiber

import (
	"gallery/api"
	"gallery/core"
	"gallery/core/auth"
	"gallery/db"
	"gallery/fiber/validators"

	"errors"

	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {
	var p api.LoginPayload
	if err := c.BodyParser(&p); err != nil {
		return core.SendErr(c, errors.New("parse error"), false)
	}

	if errs := validators.ValidateInterface(p); len(errs) > 0 {
		return core.SendValidError(c, errs, false)
	}
	jwtBackend := auth.InitJWTAuthenticationBackend()

	lp := api.LoginPayload{Email: p.Email, Password: p.Password}

	user, err := jwtBackend.Authenticate(lp)
	if err != nil {
		return core.SendErr(c, err, true)
	}

	tok, err := jwtBackend.GenerateToken(user.Email, user.ID.Hex())
	if err != nil {
		return core.SendErr(c, errors.New("error generating token"), true)
	}
	return c.JSON(fiber.Map{
		"success": true,
		"token":   tok,
	})

}
func fakeLogin(c *fiber.Ctx) error {
	email := "myemail57715@gmail.com"
	user, err := db.UserByEmail(email)
	if err != nil {
		return core.SendErrWithMsg(c, err, "error finding user")
	}
	jwtBackend := auth.InitJWTAuthenticationBackend()
	tok, err := jwtBackend.GenerateToken(user.Email, user.ID.Hex())
	if err != nil {
		return core.SendErrWithMsg(c, err, "error creating token")
	}
	return c.JSON(fiber.Map{
		"success": true,
		"token":   tok,
	})
}

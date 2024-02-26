package fiber

import (
	"gallery/api"
	"gallery/core"
	"gallery/core/auth"
	"gallery/db"
	"gallery/fiber/validators"
	"gallery/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(c *fiber.Ctx) error {
	var p api.LoginPayload
	// p := new(authPayload)
	err := c.BodyParser(&p)
	if err != nil {
		return core.SendErrString(c, "parsing error")
	}
	// log.Debug(p)

	if errs := validators.ValidateInterface(&p); len(errs) > 0 {
		return core.SendValidError(c, errs, false)
	}
	hashed, err := auth.GenPassword(p.Password)
	if err != nil {
		return core.SendErr(c, err, true)
	}

	u := models.User{Email: p.Email, Password: hashed}
	u.DateCreated = time.Now()
	_id, err := db.AddUser(u)
	if err != nil {
		return core.SendErr(c, err, true)
	}

	u.ID = _id

	jb := auth.InitJWTAuthenticationBackend()

	tok, err := jb.GenerateToken(u.Email, u.ID.Hex())

	if err != nil {
		return core.SendErr(c, err, true)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   tok,
	})

}

func _valid(p authPayload) bool {
	if p.Email != "" && p.Password != "" {
		return true
	}
	return false

}

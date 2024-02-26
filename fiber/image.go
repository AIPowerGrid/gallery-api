package fiber

import (
	"encoding/base64"
	"errors"
	"gallery/core"
	"gallery/endpoints"

	"github.com/gofiber/fiber/v2"
)

type genData struct {
	Prompt string `json:"prompt"`
	Seed   int64  `json:"seed"`
}

func generate(c *fiber.Ctx) error {
	// prompt := "Lil wayne eating rice"
	var p genData
	if err := c.BodyParser(&p); err != nil {
		return core.SendErr(c, errors.New("parse error"), false)
	}
	// id := uuid.NewString()
	/* go func() {
		_, err := endpoints.GenerateImage(p.Prompt, id, p.Seed)
		if err != nil {
			log.Error(err)
		}
	}() */
	// job := endpoints.MakeJob(p.Prompt, id, p.Seed)
	return nil
	// return c.JSON(job)
}
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// func
func getImage(c *fiber.Ctx) error {

	id := c.Params("id")
	// data, err := endpoints.GetImage(id)
	// filename := fmt.Sprintf("%s_00001_.png", id)
	data, err := endpoints.GetImage(id)
	// log.Debug(data)
	if err != nil {
		// ret
		// return c.SendString("f")
		return c.Status(200).JSON(fiber.Map{"message": "not found", "success": false})
	}

	// b := toBase64([]byte(data.ImageData))
	// c.Response().Header.Set("Content-Type", "image/png")
	// return c.SendString(string(b))
	// return c.SendFile(data.ImageData)
	return c.JSON(fiber.Map{"success": true, "image": data.ImageData})

}

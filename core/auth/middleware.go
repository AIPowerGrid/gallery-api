package auth

import (
	"gallery/db"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func sendErr(ctx *fiber.Ctx, msg string) error {
	// log.Error(msg)
	return ctx.Status(400).JSON(fiber.Map{"success": false, "message": msg})
}

func sendErrStatus(ctx *fiber.Ctx, status int, msg string) error {
	return ctx.Status(status).JSON(fiber.Map{"success": false, "message": msg})
}

func RequireAuth() fiber.Handler {

	return func(c *fiber.Ctx) error {
		auth := c.Get("authorization")
		// fmt.Println(auth,5555)

		if auth == "" {
			return sendErrStatus(c, 401, "unauthenticated")
			// return c.Status(401).JSON(fiber.Map{"message":"unauthenticated"})
		}

		split := strings.Split(auth, " ")
		if len(split) == 0 {
			return sendErr(c, "no token")
		}
		// log.Debug(split)
		last := split[len(split)-1]
		if last == "" {
			return sendErr(c, "no token")
			// return c.Status(400).JSON(fiber.Map{"message":"no token"})
		}
		last = strings.TrimSpace(last)
		// token, err := jwt.ParseWithClaims(last,)
		token, err := jwt.Parse(last, func(token *jwt.Token) (interface{}, error) {
			secretKey := appSettings.JWTSecret
			// log.Error(secretKey)
			return []byte(secretKey), nil
		})
		// token ,err:= jwt.Parse(last)
		if err != nil {
			log.Error(err)
			log.Debug(last, last[0], 3)
			return sendErr(c, err.Error())

		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error("wrong signing method for token", last, token)
			return sendErr(c, "wrong signature")
		}

		if err == nil && token.Valid {
			// log.Debug("token is valid..")
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// log.Info(claims)
				id := claims["id"].(string)
				user, err := db.UserByID(id)
				if err != nil {
					log.Error(err)
					log.Info(id)
					return sendErr(c, err.Error())
				}
				c.Locals("user", user)
				return c.Next()
			} else {
				return sendErr(c, "unknown error")
			}
		}
		return sendErr(c, "token is not valid")
	}
}

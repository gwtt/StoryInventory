package middlewares

import (
	"fiberTest/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error{
	cookies := c.Cookies("jwt")


	if _, err := util.ParseJwt(cookies);err!=nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status": 401,
			"msg": "unauthenticated",
			"data": "",
		})
	}

	return c.Next()
}
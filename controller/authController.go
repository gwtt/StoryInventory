package controller

import (
	"fiberTest/database"
	//"fiberTest/models"
	"fiberTest/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil{
		return err
	}
	if data["password"] != data["password_confirm"]{
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "密码不匹配",

		})
	}
	user := models.User{
		Name: data["name"],
		Email: data["email"],
	}
	user.SetPassword(data["password"])
	database.DB.Create(&user)
	return c.JSON(user)
}
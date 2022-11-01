package util

import "github.com/gofiber/fiber/v2"

func Map(status int,msg string,data interface{}) fiber.Map {
	return fiber.Map{
		"status":status,
		"msg":msg,
		"data":data,
	}
}
package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB,entity Entity,page int) fiber.Map {
	limit := 10
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)
	//total := entity.Count(db)

	return fiber.Map{
		"status":0,
		"msg":"获取成功",
		"data":fiber.Map{
			"object": data,

		},
	}
}
func PaginateOut(db *gorm.DB,entity Object,page int) fiber.Map {
	limit := 10
	offset := (page - 1) * limit

	data := entity.TakeOut(db, limit, offset)
	//total := entity.Count(db)

	return fiber.Map{
		"status":0,
		"msg":"获取成功",
		"data":fiber.Map{
			"object": data,
			},
			}
}
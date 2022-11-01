package controller

import (

	"fiberTest/database"
	"fiberTest/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func AllObjects(c *fiber.Ctx) error{
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB,&models.Object{},page))
}
func AllObjectsOut(c *fiber.Ctx)  error{
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.PaginateOut(database.DB,models.Object{},page))
}
func CreateObject(c *fiber.Ctx) error {
	var object models.Object
	if err := c.BodyParser(&object);err != nil{
		return err
	}

	database.DB.Create(&object)
	return c.JSON(object)
}

func GetObject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("id"))
	object := models.Object{
		Id: uint(id),
		}
		database.DB.Find(&object)
	return c.JSON(object)
}

func UpdateObject(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("id"))

	object := models.Object{
		Id: uint(id),
		}
		if err:=c.BodyParser(&object);err!=nil{
			return err
		}
	database.DB.Model(&object).Where("id = ?", object.Id).Updates(object)
	database.DB.Find(&object)
	return c.JSON(object)
}
func DeleteObject(c *fiber.Ctx) error { //出货
	id, _ := strconv.Atoi(c.Query("id"))

	object := models.Object{
		Id: uint(id),
		}
		database.DB.Delete(object)
	return nil
}

func OutObject(c *fiber.Ctx) error { //出货
	id, _ := strconv.Atoi(c.Query("id"))

	object := models.Object{
		Id: uint(id),
		}
	year, month, day := time.Now().Date()
	object.OutTime = strconv.Itoa(year)+"-"+strconv.Itoa(int(month))+"-"+strconv.Itoa(day)
	database.DB.Model(&object).Where("id = ?", object.Id).Updates(object)
		database.DB.Find(&object)
		return c.JSON(object)
}
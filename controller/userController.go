package controller

import (
	"fiberTest/database"
	"fiberTest/models"
	"fiberTest/util"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func AllUsers(c *fiber.Ctx)  error{
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB,&models.User{},page))
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user);err != nil{
		return err
	}
	user.SetPassword("1234")
	database.DB.Create(&user)
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("id"))

	user := models.User{
		Id: uint(id),
	}
	database.DB.Find(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("id"))
	log.Println(id)
	user := models.User{
		Id: uint(id),
	}

	if err:=c.BodyParser(&user);err!=nil{
		return err
	}
	log.Println(&user)

	database.DB.Model(&user).Where("id = ?",user.Id).Updates(user)
	return c.JSON(user)
}
func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("id"))

	user := models.User{
		Id: uint(id),
	}
	database.DB.Delete(user)
	return nil
}
func Alert(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)

	var user models.User
	database.DB.Where("id = ?",id).First(&user)
	return c.JSON(util.SearchPastDue(user.Email))
}
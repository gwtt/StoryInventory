package controller

import (
	"fiberTest/database"
	"fiberTest/models"
	"fiberTest/util"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx) error{
	var data map[string]string
	if err := c.BodyParser(&data);err != nil{
		return err
	}
	var user models.User
	database.DB.Where("email = ?",data["email"]).First(&user)

	if user.Id == 0{
		c.Status(404)
		return c.JSON(util.Map(404,"用户未找到",""))
	}
	if err :=bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]));err!=nil{
		c.Status(400)
		return c.JSON(util.Map(400,"密码不正确",""))
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil{
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(util.Map(0,"登录成功",token))
}



func User(c *fiber.Ctx) error{
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)

	var user models.User
	database.DB.Where("id = ?",id).First(&user)
	return c.JSON(util.Map(200,"查询成功",user))
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		}
	c.Cookie(&cookie)
	return c.JSON(util.Map(200,"成功登出",""))
}
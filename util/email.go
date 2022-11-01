package util

import (
	"fiberTest/database"
	"fiberTest/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strconv"
	"time"
)

func Alert(content string,Email string) {
	newEmail := email.NewEmail()
	newEmail.From = "顾文韬 <1213605030@qq.com>"
	newEmail.To=[]string{Email}
	newEmail.Subject = "商品过期报警"
	newEmail.Text =[]byte(content)
	newEmail.Send("smtp.qq.com:587",smtp.PlainAuth("","1213605030@qq.com","hepbotbzyssbiece","smtp.qq.com"))
}
func SearchPastDue(Email string) interface{} {
	var objects []models.Object
	year, month, day := time.Now().Date()
	Time := strconv.Itoa(year)+"-"+strconv.Itoa(int(month))+"-"+strconv.Itoa(day)
	database.DB.Where("expiration_time < ? and out_time =''",Time).Find(&objects) //找没出货且过期的商品
	var information string
	for _,o:= range objects{
		information += o.Toshow()
	}
	Alert(information,Email)
	return fiber.Map{
		"status":0,
		"msg":"获取过期产品成功",
		"data":fiber.Map{
			"object": objects,
			},
	}
}
package service

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func Alert(content string,Email string) {
	newEmail := email.NewEmail()
	newEmail.From = "顾文韬 <1213605030@qq.com>"
	newEmail.To=[]string{Email}
	newEmail.Subject = "商品过期报警"
	newEmail.Text =[]byte(content)
	newEmail.Send("smtp.qq.com:587",smtp.PlainAuth("","1213605030@qq.com","hepbotbzyssbiece","smtp.qq.com"))
}
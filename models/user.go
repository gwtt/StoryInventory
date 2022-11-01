package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	 Id uint    		`json:"id"`
	 Name string		`json:"name"`
	 Email string		`json:"email"`
	 Password []byte	`json:"-"`
 }

func (user *User) SetPassword(password string) {
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = fromPassword
}

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

func (user *User) Take(db *gorm.DB,limit int,offset int) interface{} {
	var users []User
	db.Offset(offset).Limit(limit).Find(&users)
	return users
}
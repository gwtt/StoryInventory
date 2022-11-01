package models

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Object struct{
	Id 			 uint
	Specie 		 string `json:"specie"`//种类
	Name 		 string `json:"name"`//商品名字
	Price 		 float64 `json:"price"`//商品价格
	Number int `json:"number"`
	Unit string `json:"unit"`
	PurChaseTime  string `json:"pur-chase-time"`// 进货时间
	ExpirationTime  string `json:"expiration-time"` //过期时间
	OutTime string `json:"out-time"` //出货时间
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at,omitempty"`
	UpdateAt  time.Time `gorm:"column:update_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp" json:"update_at,omitempty"`
}
func (object *Object) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

func (object *Object) Take(db *gorm.DB,limit int,offset int) interface{} {
	var objects []Object
	db.Where("out_time = ''").Find(&objects)
	return objects
}

func (object *Object) TakeOut(db *gorm.DB,limit int,offset int) interface{} {
	var objects []Object
	db.Where("out_time != ''").Find(&objects)
	return objects
}
func (object *Object) Toshow() string{
	return "产品ID :"+strconv.Itoa(int(object.Id))+"，种类: "+object.Specie + ",名称: " +object.Name + "过期了\n"
}
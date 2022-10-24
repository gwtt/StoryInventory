package dao

import (
	"strings"
	"time"
)

type Object struct{
	Specie string //种类
	Name string //商品名字
	PurChaseTime string // 进货时间
	PastDue string//过期时间

}
type Varieable interface {
	ToString() string
	ToShow() string
}

func (o Object) ToString() string{ //为了存放数据
	var information string
	information = information + o.Specie + ","
	information = information + o.Name +","
	timeNow := time.Now()
	if o.PurChaseTime == ""{
		information = information + timeNow.Format("2006-01-02") + ","
	}else{
		information = information + o.PurChaseTime + ","
	}
	information = information + o.PastDue
	return information
}
func (o Object) ToShow() string{ //为了展示数据
	var information string
	information = information + "商品种类:"+ o.Specie + " "
	information = information + "商品名字:"+ o.Name +" "
	split := strings.Split(o.PurChaseTime, "-")
	information = information + "进货时间:"+split[0]+"年"+split[1]+"月"+split[2]+"日" + " "
	split = strings.Split(o.PastDue, "-")
	information = information + "过期时间:"+split[0]+"年"+split[1]+"月"+split[2]+"日"+"\n"
	return information
}

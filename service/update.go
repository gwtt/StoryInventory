package service

import (
	"StoreInventoryManagementSystem/dao"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Update(origin dao.Object,to dao.Object)  (string,error,dao.Object) {
	var sum []dao.Object
	file, err := os.OpenFile("./fyne.txt", os.O_APPEND|os.O_RDWR, 07777)
	if err != nil{
		return "",err,dao.Object{}
	}
	readBytes, err := ioutil.ReadAll(file)

	if err != nil{
		return "",err,dao.Object{}
	}
	results := strings.Split(string(readBytes), "\n")
	for i :=0;i<len(results)-1;i++{
		split := strings.Split(results[i], ",")
		var object dao.Object
		object.Specie = split[0]
		object.Name = split[1]
		object.PurChaseTime = split[2]
		object.PastDue = split[3]
		count := 0
		if object.Specie == origin.Specie{
			count++
		}
		if object.Name == origin.Name{
			count++
		}
		if object.PurChaseTime == origin.PurChaseTime{
			count++
		}
		if object.PastDue == origin.PastDue{
			count++
		}
		if count < 4{
			sum = append(sum,object)
		}else{
			sum = append(sum,to)
		}

	}
	file.Close()
	if len(sum) == 0{
		return "没有数据，无法修改哦！", nil,dao.Object{}
	}
	os.Truncate("./fyne.txt",0)
	for _,one :=range sum{
		PurChase(one)
	}
	var information string
	information = "修改后数据量为:"+strconv.Itoa(len(sum))+"个 下面是修改后的数据\n\n"
	var length = len(sum)
	if length > 10{
		length = 10
	}
	for  i := 0;i<length;i++{
		information +=sum[i].ToShow()
	}
	if len(sum) > 10{
		information +="......"
	}
	return information, nil,to
}
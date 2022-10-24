package service

import (
	"StoreInventoryManagementSystem/dao"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Query(query dao.Object)  (string,error,dao.Object){
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
		count := 0
		if query.Specie != ""{
			if split[0] == query.Specie{
				object.Specie = split[0]
				count++
			}
		}else{
			object.Specie = split[0]
			count++
		}

		if query.Name != ""{
			if split[1] == query.Name{
				object.Name = split[1]
				count++
			}
		}else{
			object.Name = split[1]
			count++
		}

		if query.PurChaseTime != ""{
			if split[2] == query.PurChaseTime{
				object.PurChaseTime = split[2]
				count++
			}
		}else{
			object.PurChaseTime = split[2]
			count++
		}

		if query.PastDue != ""{
			if split[3] == query.PastDue{
				object.PastDue = split[3]
				count++
			}
		}else{
			object.PastDue = split[3]
			count++
		}

		if count==4{
			sum = append(sum,object)
		}
	}
	file.Close()
	if len(sum) == 0{
		return "对不起，没有数据喔!", nil,dao.Object{}
	}
	var information string
	information = "现在数据量为:"+strconv.Itoa(len(sum))+"个 下面是查询的数据\n\n"
	var length = len(sum)
	if length > 10{
		length = 10
	}
	for  i := 0;i<length;i++{
		fmt.Println(sum[i])
		information +=sum[i].ToShow()
	}
	if len(sum) > 10{
		information +="......"
	}
	return information, nil,sum[0]
}
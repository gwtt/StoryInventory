package service

import (
	"StoreInventoryManagementSystem/dao"
	"log"
	"sort"
	"strconv"
	"time"
	"io/ioutil"
	"os"
	"strings"
)

func OpenAndWrite(information string) error { //进货
	file, err := os.OpenFile("./fyne.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 07777)
	if err != nil{
		return err
	}
	file.WriteString(information + "\n")
	return nil
}
func OpenAndDelete() (string,error) { //出货
	var sum []dao.Object
	var SaledObject dao.Object
	file, err := os.OpenFile("./fyne.txt", os.O_APPEND|os.O_RDWR, 07777)
	if err != nil{
		return "",err
	}
	readBytes, err := ioutil.ReadAll(file)

	if err != nil{
		return "",err
	}
	results := strings.Split(string(readBytes), "\n")
	for i :=0;i<len(results)-1;i++{
		split := strings.Split(results[i], ",")
		var object dao.Object
		object.Specie = split[0]
		object.Name = split[1]
		object.PurChaseTime = split[2]
		object.PastDue = split[3]
		if SaledObject == (dao.Object{}){
			SaledObject = object
		}else {
			Pucompare := Compare(object.PurChaseTime, SaledObject.PurChaseTime)
			if Pucompare==0{
				sum = append(sum,SaledObject)
				SaledObject = object
			}else if Pucompare == -1{
				PaCompare := Compare(object.PastDue, SaledObject.PastDue)
				if(PaCompare==0){
					sum = append(sum,SaledObject)
					SaledObject = object
				}else{
					sum = append(sum,object)
				}
			}else{
				sum = append(sum,object)
			}
		}
	}
	file.Close()
	log.Println("现在数据量:"+strconv.Itoa(len(sum)))
	if len(sum) == 0{
		return "现在数据量为0",nil
	}
	os.Truncate("./fyne.txt",0)
	for _,one :=range sum{
		PurChase(one)
	}
	log.Println("出货操作:"+SaledObject.ToShow())
	return "出货操作:"+SaledObject.ToShow()+"\n",nil
}
func SearchPastDue() (string,error) {
	var sum []dao.Object
	file, err := os.OpenFile("./fyne.txt", os.O_APPEND|os.O_RDWR, 07777)
	if err != nil{
		return "",err
	}
	readBytes, err := ioutil.ReadAll(file)

	if err != nil{
		return "",err
	}
	results := strings.Split(string(readBytes), "\n")
	for i :=0;i<len(results)-1;i++{
		split := strings.Split(results[i], ",")
		if Compare(time.Now().Format("2006-01-02"),split[3]) == 1{
			var object dao.Object
			object.Specie = split[0]
			object.Name = split[1]
			object.PurChaseTime = split[2]
			object.PastDue = split[3]
			sum = append(sum,object)
		}
	}
	file.Close()
	if len(sum) == 0{
		return "", nil
	}
	var information string
	for _,o := range sum{
		information +=o.ToShow()+"过期了\n"
	}
	return information, nil
}
func Compare(date1 string,date2 string) int { //比较,如果第一个小返回false
	t1, _ := time.Parse("2006-01-02", date1) //改变两个日期数据格式，方便比较
	t2, _ := time.Parse("2006-01-02", date2)
	dat1 := t1.Unix()
	dat2 := t2.Unix()
	switch {
	case dat1 == dat2:
		return -1
		case dat1 < dat2:
			return 0
			default :
				return 1
	}
}
func SearchSum() (string,error) {
	var sum []dao.Object
	file, err := os.OpenFile("./fyne.txt", os.O_APPEND|os.O_RDWR, 07777)
	if err != nil{
		return "",err
	}
	readBytes, err := ioutil.ReadAll(file)

	if err != nil{
		return "",err
	}
	results := strings.Split(string(readBytes), "\n")
	for i :=0;i<len(results)-1;i++{
		split := strings.Split(results[i], ",")
		var object dao.Object
		object.Specie = split[0]
		object.Name = split[1]
		object.PurChaseTime = split[2]
		object.PastDue = split[3]
		sum = append(sum,object)
	}
	file.Close()
	if len(sum) == 0{
		return "对不起，没有数据喔!", nil
	}
	var information string
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
	return information, nil
}
func Sort(String string) (string,error)  {
	var sum []dao.Object
	file, err := os.OpenFile("./fyne.txt", os.O_APPEND|os.O_RDWR, 07777)
	if err != nil{
		return "",err
	}
	readBytes, err := ioutil.ReadAll(file)
	if err != nil{
		return "",err
	}
	results := strings.Split(string(readBytes), "\n")
	for i :=0;i<len(results)-1;i++{
		split := strings.Split(results[i], ",")
		var object dao.Object
		object.Specie = split[0]
		object.Name = split[1]
		object.PurChaseTime = split[2]
		object.PastDue = split[3]
		sum = append(sum,object)
	}
	if len(sum) == 0{
		log.Println("对不起，没有数据喔!")
	}
	switch String{
	case "按照种类排序":
		sort.Slice(sum, func(i, j int) bool {
			compare := strings.Compare(sum[i].Specie, sum[j].Specie)
			if compare == -1{
				return true
			}else{
				return false
			}
		})
	case "按照名字排序":sort.Slice(sum, func(i, j int) bool {
		compare := strings.Compare(sum[i].Name, sum[j].Name)
		if compare == -1{
			return true
		}else{
			return false
		}
	})
	case "按照进货时间排序":sort.Slice(sum, func(i, j int) bool {
		compare := strings.Compare(sum[i].PurChaseTime, sum[j].PurChaseTime)
		if compare == -1{
			return true
		}else{
			return false
		}
	})
	case "按照过期时间排序":sort.Slice(sum, func(i, j int) bool {
		compare := strings.Compare(sum[i].PastDue, sum[j].PastDue)
		if compare == -1{
			return true
		}else{
			return false
		}
	})
	}
	file.Close()
	os.Truncate("./fyne.txt",0)
	for _,one :=range sum{
		PurChase(one)
	}
	var information string
	information = "总数据量:"+strconv.Itoa(len(sum))+"\n\n排序后:\n"
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
	return information, nil
}
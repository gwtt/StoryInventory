package service

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func WriteToExcel() error{
	f, err := os.Create("Data.csv")
	if err != nil{
		log.Println(err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	writer.Write([]string{"商品种类","商品名字","进货时间","过期时间"})
	file, err := os.OpenFile("./fyne.txt", os.O_APPEND|os.O_RDWR, 07777)
	if err != nil{
		return err
	}
	readBytes, err := ioutil.ReadAll(file)
	if err != nil{
		return err
	}
	results := strings.Split(string(readBytes), "\n")
	for i :=0;i<len(results)-1;i++{
		split := strings.Split(results[i], ",")
		writer.Write([]string{split[0],split[1],split[2],split[3]})
	}
	writer.Flush()
	file.Close()
	return nil
}
package service

import (
	"StoreInventoryManagementSystem/dao"
)

func PurChase(goods dao.Object) error{ //进货使用
	err := OpenAndWrite(goods.ToString())
	if err != nil{
		return err
	}
	return nil
}
func Shipment() error { //出货使用
	_,err := OpenAndDelete()
	if err != nil{
		return err
	}
	return nil
}
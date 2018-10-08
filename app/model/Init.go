package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/demo?parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	} else {
		db.LogMode(true)
	}

	//fmt.Println("实例对象的信息和地址", sin, &sin)
	return db
}

//func init() {
//	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/demo?parseTime=true&loc=Local")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	db.LogMode(true)
//	defer db.Close()
//}

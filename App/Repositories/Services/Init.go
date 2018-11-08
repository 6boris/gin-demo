package Services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	db.LogMode(true)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}

package Services

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Asia%2FShanghai")

	if err != nil {
		fmt.Println(err)
	}

	return db
}

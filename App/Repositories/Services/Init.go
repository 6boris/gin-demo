package Services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	//	open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

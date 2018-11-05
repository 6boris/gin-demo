package MySQL

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Asia%2FShanghai")

	if err != nil {
		fmt.Println(err)
	}

	return db
}

func TestArticle_TableName(t *testing.T) {
	db := GetDB()
	articles := []Article{}

	db.Find(&articles)

	for _, v := range articles {
		fmt.Println(v.Id, v.Title, v.Created_at, v.Updated_at)
	}
}
func TestArticle_TableName2(t *testing.T) {
	db := GetDB()
	articles := []Article{}

	db.Find(&articles)

	for _, v := range articles {
		fmt.Println(v.Id, v.Title, v.Created_at, v.Updated_at, v.Description)
	}
}

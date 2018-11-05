package Services

import (
	"github.com/kylesliu/gin-demo/App/Repositories/MySQL"
	//"github.com/kylesliu/gin-demo/Bootstrap"
)

func GetAllArticle() *[]MySQL.Article {
	//db, err := gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//db := GetDB()

	articles := []MySQL.Article{}
	db.Find(&articles)
	return &articles
}

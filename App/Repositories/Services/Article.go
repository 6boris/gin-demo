package Services

import (
	"github.com/kylesliu/gin-demo/App/Repositories/MySQL"
	//"github.com/kylesliu/gin-demo/Bootstrap"
)

func GetAllArticle(group_id int) *[]MySQL.Article {

	articles := []MySQL.Article{}

	if group_id == 0 {
		db.Find(&articles)

	} else {
		db.Where("g_id = ?", group_id).Find(&articles)
	}

	return &articles
}

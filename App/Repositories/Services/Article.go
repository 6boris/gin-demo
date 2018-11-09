package Services

import (
	"github.com/kylesliu/gin-demo/App/Repositories/MySQL"
)

func GetOneArticle(id int) *MySQL.Article {
	article := MySQL.Article{}
	//article.Id = id

	db.Where("id = ?", id).Find(&article)

	return &article
}

func GetAllArticle(group_id int) *[]MySQL.Article {

	articles := []MySQL.Article{}

	if group_id == 0 {
		db.Order("updated_at desc").Find(&articles)
	} else {
		db.Where("g_id = ?", group_id).Order("updated_at desc").Find(&articles)
	}

	return &articles
}

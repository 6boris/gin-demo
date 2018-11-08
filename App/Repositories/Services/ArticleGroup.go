package Services

import (
	"github.com/kylesliu/gin-demo/App/Repositories/MySQL"
	//"github.com/kylesliu/gin-demo/Bootstrap"
)

func GetAllArticleGroup() *[]MySQL.ArticleGroup {
	groups := []MySQL.ArticleGroup{}
	db.Table("blog_article_group").
		Select("" +
			"blog_article_group.id, " +
			"blog_article_group.name," +
			"count(blog_article.id) AS count, " +
			"blog_article_group.created_at, " +
			"blog_article_group.updated_at").
		Joins("left join blog_article on blog_article_group.id = blog_article.g_id").
		Group("blog_article_group.id").
		Find(&groups)
	return &groups
}

package MySQL

import "time"

type ArticleGroup struct {
	Id         int        `gorm:"column:id" json:"id"`
	Name       string     `gorm:"column:name" json:"name"`
	Count      int        `gorm:"column:count" json:"count"`
	Created_at *time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (self ArticleGroup) TableName() string {
	return "article_group"
}

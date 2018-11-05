package MySQL

import "time"

type Article struct {
	Id          int        `gorm:"column:id" json:"id"`
	Title       string     `gorm:"column:title" json:"title"`
	Content     string     `gorm:"column:content" json:"content"`
	Cover       string     `gorm:"column:cover" json:"cover"`
	Description string     `gorm:"column:description" json:"description"`
	Created_at  *time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (self Article) TableName() string {
	return "blog_article"
}

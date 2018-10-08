package model

type Group struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (Group) TableName() string {
	return "blog_article_group"
}

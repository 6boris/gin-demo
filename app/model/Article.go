package model

type Article struct {
	ID      int    `gorm:"primary_key" json:"id"`
	GroupId int    `gorm:"column:g_id" json:"g_id"`
	Title   string `json:"title"`

	Group Group `gorm:"ForeignKey:GroupId" json:"group"`
}

func (Article) TableName() string {
	return "blog_article"
}

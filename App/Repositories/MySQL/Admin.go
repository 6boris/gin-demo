package MySQL

import "time"

type Admin struct {
	Id          int        `gorm:"column:id; primary_key ; AUTO_INCREMENT" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Email       string     `gorm:"column:email" json:"email"`
	Password    string     `gorm:"column:password" json:"password"`
	Created_at  *time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (self Admin) TableName() string {
	return "demo"
}

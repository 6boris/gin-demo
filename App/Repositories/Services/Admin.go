package Services

import (
	"fmt"
	"github.com/kylesliu/gin-demo/App/Repositories/MySQL"
)

func GetAdminInfo() []*MySQL.Admin {
	admins := []*MySQL.Admin{}

	db.Find(&admins)
	return admins
}

func AdminLogin(email, password string) (interface{}, error) {
	var err error
	var count int
	admin := MySQL.Admin{}

	res := db.Table("blog_demo").Where("email = ?", email).Find(&admin).Count(&count)

	if count == 0 {
		err.Error()
		return nil, err
	}

	fmt.Println(admin)

	if res.Error != nil {
		err = res.Error
	}

	return &res, err
}

func AdminRegister(emal, password string) (interface{}, error) {
	var err error
	admin := MySQL.Admin{
		Email:    emal,
		Password: Encryption(password),
	}
	res := db.Table("blog_demo").Create(&admin)

	if res.Error != nil {
		err = res.Error
	}

	return &res, err
}

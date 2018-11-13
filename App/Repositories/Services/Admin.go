package Services

import (
	"errors"
	"fmt"
	"github.com/kylesliu/gin-demo/App/Repositories/MySQL"
)

func GetAdminInfo() []*MySQL.Admin {
	admins := []*MySQL.Admin{}

	db.Table("blog_admins").Find(&admins)
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

func AdminLoginCheck(email, password string) (interface{}, error) {
	admin := MySQL.Admin{Email: email}
	count := 0
	db.Table("blog_admins").Where("email = ?", email).Find(&admin).Count(&count)

	if count == 0 {
		return nil, errors.New("user is not exits")
	} else if admin.Password != Encryption(password) {
		return nil, errors.New("incorrect Username or Password")
	}
	return &admin, nil
}

func AdminRegister(emal, password string) (interface{}, error) {
	var err error
	admin := MySQL.Admin{
		Email:    emal,
		Password: Encryption(password),
	}
	res := db.Table("blog_admins").Create(&admin)

	if res.Error != nil {
		err = res.Error
	}

	return &res, err
}

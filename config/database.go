package Config

import (
	"gin-laravel/app/Models"
	"github.com/jinzhu/gorm"
)

func InitDatabase() {
	var err error
	Db, err = gorm.Open("sqlite3", "test.db?cache=shared")
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}

	Db.AutoMigrate(&Models.UserModel{})

}

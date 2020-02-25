package Config

import (
	"fmt"
	"gin-laravel/app/Models"
	"github.com/jinzhu/gorm"
)

func InitDatabase() {
	var err error
	Db, err = gorm.Open("sqlite3", App.GetString("sqlite.url"))
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}

	Db.AutoMigrate(&Models.UserModel{})
	fmt.Printf("sqlite3数据库连接成功: %s \r\n",
		App.GetString("sqlite.url"),
	)
}

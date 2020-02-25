package Models

import (
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model

	Name string `json:"name" gorm:"type:varchar(100);"`
	Pass string `json:"pass" gorm:"type:varchar(100)"`
}

func (UserModel) TableName() string {
	return "users"
}

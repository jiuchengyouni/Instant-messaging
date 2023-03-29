package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	//Name   string `gorm:"column:name;type:varchar(50);" json:"name"`
	//Info   string `gorm:"column:info;type:varchar(50);" json:"info"`
	UserId int `gorm:"column:user_id;type:int(11);" json:"user_id"` //创建房间的人
}

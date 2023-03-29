package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserId int    `gorm:"column:user_id;type:int(11);" json:"user_id"`
	RoomId int    `gorm:"column:room_id;type:int(11);" json:"room_id"`
	Data   string `gorm:"column:data;type:varchar(50);" json:"data"`
}

type MessageBasic struct {
	Message string `json:"message"`
	RoomId  int    `json:"room_id"`
}

package models

import "gorm.io/gorm"

type UserRoom struct {
	gorm.Model
	UserId   int `gorm:"column:user_id;type:int(11);" json:"user_id"`
	RoomType int `gorm:"column:room_type;type:int(11);" json:"room_type"` // 房间 类型 【1-独聊房间 2-群聊房间】
	RoomId   int `gorm:"column:room_id;type:int(11);" json:"room_id"`
}

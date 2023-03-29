package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"column:nickname;type:varchar(50);" json:"nickname"`
	Account  string `gorm:"column:account;type:varchar(50) not null;" json:"account"`
	Password string `gorm:"column:password;type:varchar(50) not null;" json:"password"`
	Email    string `gorm:"column:email;type:varchar(50) not null;" json:"email"`
	Sex      int    `gorm:"column:sex;type:int(11) not null;" json:"sex"`
}

type UserQueryResult struct {
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"` // 是否是好友 【true-是，false-否】
}

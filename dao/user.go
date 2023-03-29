package dao

import (
	"IM/models"
	"strconv"
)

func GetUserBasicByAccountPassword(account, password string) (user *models.User, err error) {
	err = DB.Model(&models.User{}).
		Where("account=? AND password=?", account, password).
		First(&user).Error
	return
}

func GetUserBasicById(id string) (user *models.User, err error) {
	id2, _ := strconv.Atoi(id)
	err = DB.Model(&models.User{}).
		Where("id=?", id2).
		First(&user).Error
	return
}

func GetUserBasicCountByEmail(email string) (count int64, err error) {
	err = DB.Model(&models.User{}).
		Where("email=?", email).
		Count(&count).Error
	return
}

func GetUserBasicCountByAccount(account string) (count int64, err error) {
	err = DB.Model(&models.User{}).
		Where("account=?", account).
		Count(&count).Error
	return
}

func InsertOneUser(user *models.User) (err error) {
	err = DB.Model(&models.User{}).
		Save(&user).Error
	return
}

func GetUserBasicByAccount(account string) (user *models.User, err error) {
	err = DB.Model(&models.User{}).
		Where("account=?", account).
		First(&user).Error
	return
}

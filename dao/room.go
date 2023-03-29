package dao

import "IM/models"

func InsertOneRoom(room *models.Room) (err error) {
	err = DB.Model(&models.Room{}).
		Save(&room).Error
	return
}

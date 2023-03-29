package dao

import "IM/models"

func InsertOneMessage(message *models.Message) (err error) {
	err = DB.Model(&models.Message{}).
		Save(&message).Error
	return
}

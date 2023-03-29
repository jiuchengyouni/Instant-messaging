package dao

import (
	"IM/models"
	"log"
	"strconv"
)

func GetUserRoomByUserIdRoomId(userIdS string, roomId int) (userRoom *models.UserRoom, err error) {
	userId, _ := strconv.Atoi(userIdS)
	err = DB.Model(&models.UserRoom{}).
		Where("user_id=? AND room_id=?", userId, roomId).
		First(&userRoom).Error
	return
}

func GetUserRoomByRoomId(roomId int) (userRooms []*models.UserRoom, err error) {
	err = DB.Model(&models.UserRoom{}).
		Where("room_id=?", roomId).
		Find(&userRooms).Error
	return
}

func JudgeUserIsFriend(userId1 int, userIdS2 string) bool {
	rooms1 := make([]int, 0)
	err := DB.Model(&models.UserRoom{}).
		Where("user_id=? AND room_type=?", userId1, 1).
		Select("room_id").
		Find(&rooms1).Error
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return false
	}
	rooms2 := make([]int, 0)
	userId2, _ := strconv.Atoi(userIdS2)
	err = DB.Model(&models.UserRoom{}).
		Where("user_id=? AND room_type=?", userId2, 1).
		Select("room_id").
		Find(&rooms2).Error
	maps := make(map[int]int)
	for _, v := range rooms1 {
		maps[v] = 1
	}
	for _, v := range rooms2 {
		if maps[v] == 1 {
			return true
		}
	}
	return false
}

func InsertOneUserRoom(userRoom *models.UserRoom) (err error) {
	err = DB.Model(&models.UserRoom{}).
		Save(&userRoom).Error
	return
}

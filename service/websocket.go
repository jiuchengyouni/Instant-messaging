package service

import (
	"IM/dao"
	"IM/models"
	"IM/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{}
var wc = make(map[string]*websocket.Conn)

func WebsocketMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("[1ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误" + err.Error(),
		})
		return
	}
	//关闭连接
	defer conn.Close()
	user := c.MustGet("user_claims").(*utils.UserClaims)
	wc[user.Identity] = conn
	for {
		messageBasic := models.MessageBasic{}
		err = conn.ReadJSON(&messageBasic)
		if err != nil {
			log.Printf("[2ERROR]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "系统错误" + err.Error(),
			})
			return
		}
		err = conn.ReadJSON(&messageBasic)
		if err != nil {
			log.Printf("[2ERROR]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "系统错误" + err.Error(),
			})
			return
		}
		_, err := dao.GetUserRoomByUserIdRoomId(user.Identity, messageBasic.RoomId)
		if err != nil {
			log.Printf("[DB ERROR]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据库查询错误" + err.Error(),
			})
			return
		}
		//保存消息
		userId, _ := strconv.Atoi(user.Identity)
		message := models.Message{
			UserId: userId,
			RoomId: messageBasic.RoomId,
			Data:   messageBasic.Message,
		}
		err = dao.InsertOneMessage(&message)
		if err != nil {
			log.Printf("[DB ERROR]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据库错误" + err.Error(),
			})
			return
		}
		//获取在线用户
		userRooms, err := dao.GetUserRoomByRoomId(messageBasic.RoomId)
		if err != nil {
			log.Printf("[DB ERROR]:%v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据库查询错误" + err.Error(),
			})
			return
		}
		//发送消息
		for _, room := range userRooms {
			fmt.Println(room.UserId)
			if cc, ok := wc[strconv.Itoa(room.UserId)]; ok {
				err = cc.WriteMessage(websocket.TextMessage, []byte(messageBasic.Message))
				if err != nil {
					log.Printf("[ERROR]:%v\n", err)
					c.JSON(http.StatusOK, gin.H{
						"code": -1,
						"msg":  "系统错误" + err.Error(),
					})
					return
				}
			}
		}
	}
}

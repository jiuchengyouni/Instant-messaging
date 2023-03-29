package routes

import (
	"IM/middleware"
	"IM/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/login", service.Login)

	//未完善
	r.POST("/send/code", service.SendCode)

	r.POST("/register", service.Register)
	auth := r.Group("/u", middleware.AuthCheck())
	auth.GET("/user/detail", service.UserDetail)
	auth.GET("/websocket/message", service.WebsocketMessage)
	auth.POST("/user/add", service.UserAdd)
	return r
}

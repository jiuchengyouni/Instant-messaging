package middleware

import (
	"IM/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := utils.AnalyseToken(token)
		if err != nil {
			c.Abort() //请求中断
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证未通过",
			})
			return
		}
		c.Set("user_claims", userClaims)
		c.Next() //进入下一层service
	}
}

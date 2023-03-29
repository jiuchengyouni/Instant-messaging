package service

import (
	"IM/dao"
	"IM/models"
	"IM/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	ub, err := dao.GetUserBasicByAccountPassword(account, utils.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "未找到该用户",
		})
		return
	}
	token, err := utils.GenerateToken(strconv.Itoa(int(ub.ID)), ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserDetail(c *gin.Context) {
	u, _ := c.Get("user_claims")
	uc := u.(*utils.UserClaims)
	userBasic, err := dao.GetUserBasicById(uc.Identity)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": userBasic,
	})
}

// 未完善
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	count, err := dao.GetUserBasicCountByEmail(email)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常:" + err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前邮箱已被注册",
		})
		return
	}
	code := utils.GetCode()
	err = utils.SendCode(email, code)
	if err != nil {
		log.Printf("[ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": "",
	})
}

func Register(c *gin.Context) {
	account := c.PostForm("account")
	email := c.PostForm("email")
	password := c.PostForm("password")
	if account == "" || email == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	count, err := dao.GetUserBasicCountByAccount(account)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常:" + err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号已被注册",
		})
		return
	}
	user := models.User{
		Account:  account,
		Password: utils.GetMd5(password),
		Email:    email,
	}
	err = dao.InsertOneUser(&user)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据插入异常:" + err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(string(user.ID), user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserAdd(c *gin.Context) {
	account := c.PostForm("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	user, err := dao.GetUserBasicByAccount(account)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常:" + err.Error(),
		})
		return
	}
	userClaim := c.MustGet("user_claims").(*utils.UserClaims)
	if dao.JudgeUserIsFriend(int(user.ID), userClaim.Identity) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "已为好友",
		})
		return
	}
	userClaimId, _ := strconv.Atoi(userClaim.Id)
	room := &models.Room{
		UserId: userClaimId,
	}
	err = dao.InsertOneRoom(room)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常:" + err.Error(),
		})
		return
	}
	userRoom1 := &models.UserRoom{
		UserId:   int(user.ID),
		RoomType: 1,
		RoomId:   int(room.ID),
	}
	err = dao.InsertOneUserRoom(userRoom1)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常:" + err.Error(),
		})
		return
	}
	userRoom2 := &models.UserRoom{
		UserId:   userClaimId,
		RoomType: 1,
		RoomId:   int(room.ID),
	}
	err = dao.InsertOneUserRoom(userRoom2)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

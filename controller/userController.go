package controller

import (
	"douyin-lite/comm"
	"douyin-lite/dao"
	"douyin-lite/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context)  {
	db := comm.GetDB()
	//	获取参数
	name := c.PostForm("username")
	password := c.PostForm("password")
	//	表单校验
	if len(name) == 0 || len(password) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg": "注册失败",
		})
		return
	}
	if dao.IsNameExist(db, name) {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg": "用户已存在",
		})
		return
	}
	//	保存用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"status_msg": "注册出现错误",
		})
	}
	user := model.User{Name: name, Password: string(hashedPassword)}
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg": "注册成功",
		"user_id": user.ID,
		"token": "string",
	})
}

func Login(c *gin.Context) {
	db := comm.GetDB()
	// 获取参数
	name := c.PostForm("username")
	password := c.PostForm("password")
	//	表单校验
	if len(name) == 0 || len(password) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg": "登录失败",
		})
		return
	}
	//	查询用户
	user := dao.FindByUsername(db, name)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg": "用户不存在",
		})
		return
	}
	//	密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg": "密码错误",
		})
		return
	}
	//	发放 token
	token, err := comm.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": -1,
			"status_msg": "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg": "登录成功",
		"user_id": user.ID,
		"token": token,
	})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg": "获取成功",
		"user": user,
	})
}

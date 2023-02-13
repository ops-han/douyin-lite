package router

import (
	"douyin-lite/controller"
	"douyin-lite/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/douyin/user/register", controller.Register)
	r.POST("/douyin/user/login", controller.Login)
	r.GET("/douyin/user", middleware.AuthMiddleWare(), controller.Info)

	return r
}
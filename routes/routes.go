package routes

import (
	"bluebell/controllers"
	"bluebell/logger"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 测试框架搭建是否成功
	r.GET("/", controllers.LinkTestHandler)

	// 注册业务路由
	r.POST("/signUp", controllers.SignUpHandler)

	// 登录业务路由
	r.POST("/login", controllers.LoginHandler)

	return r
}

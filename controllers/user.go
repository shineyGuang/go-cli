package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.UserSignUp) // 返回指针
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("注册传入参数错误！", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 402,
			"msg":  "注册输入参数错误！",
		})
		return
	}
	//fmt.Println(p)
	// 2. 业务处理【放在logic层】
	res := logic.SignUp(p)
	// 3. 返回响应
	if res {
		zap.L().Info("注册成功！", zap.String("userName", p.UserName))
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "注册成功",
		})
	} else {
		zap.L().Error("注册失败！", zap.String("userName", p.UserName))
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "注册失败！",
		})
	}
}
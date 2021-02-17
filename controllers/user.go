package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// LinkTestHandler 测试框架搭建是否成功
func LinkTestHandler(c *gin.Context) {
	Res(c, CodeSuccess, "")
}

// SignUpHandler 处理注册请求函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.UserSignUp) // 返回指针
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResWithMsg(c, CodeServerBusy, errors.New("非参数类型错误！"), "")
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResWithMsg(c, CodeInvalidParams, errs.Translate(trans), "")
		return
	}
	//fmt.Println(p)
	// 2. 业务处理【放在logic层】
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("注册失败！", zap.String("userName", p.UserName))
		Res(c, CodeSignUpFailed, "")
	} else {
		zap.L().Info("注册成功！", zap.String("userName", p.UserName))
		Res(c, CodeSuccess, map[string]string{"userName": p.UserName, "pwd": p.PassWord})
	}
}

// LoginHandler 处理登录业务函数
func LoginHandler(c *gin.Context) {
	loginUser := new(models.UserLogin)
	// 1. 参数校验
	if err := c.ShouldBindJSON(loginUser); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 返回非参数校验类型错误
			ResWithMsg(c, CodeServerBusy, errors.New("非参数类型错误！"), "")
		} else {
			ResWithMsg(c, CodeInvalidParams, errs.Translate(trans), "")
			return
		}
	}
	// 2. 查库看用户是否存在
	if logic.Login(loginUser) {
		// 此处表示登录成功
		Res(c, CodeSuccess, map[string]string{"userName": loginUser.UserName, "pwd": loginUser.PassWord, "token": "以后会有的"})
		zap.L().Info("登录成功", zap.String("username", loginUser.UserName))
	} else {
		Res(c, CodeLoginFailed, "")
		zap.L().Info("登录失败", zap.String("username", loginUser.UserName))
	}
	// 3. 返回响应
}

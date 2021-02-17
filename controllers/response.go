package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func Res(c *gin.Context, code ResCode, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.getMsg(),
		Data: data,
	})
}

func ResWithMsg(c *gin.Context, code ResCode, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

package logic

import (
	"bluebell/dao/dealt"
	"bluebell/models"
	"bluebell/pkg/md5"
	"bluebell/pkg/snowflake"
	"errors"

	"go.uber.org/zap"
)

// SignUp 此处通过参数校验，处理注册逻辑
func SignUp(p *models.UserSignUp) (err error) {
	// 1. 查询是否存在，如果存在直接返回
	if !dealt.CheckUserExist(p) {
		return errors.New("用户已存在")
	} else {
		// 2. 生成user_id
		p.UserId = snowflake.GenID()
		p.PassWord = md5.EncryptPassword(p.PassWord)
		if err := dealt.InsertUser(p); err != nil {
			zap.L().Error("insert data failed", zap.Error(err))
			return err
		} else {
			return nil
		}
	}
}

// Login 处理登录逻辑
func Login(p *models.UserLogin) bool {
	p.PassWord = md5.EncryptPassword(p.PassWord)
	return dealt.LoginCheck(p)
}

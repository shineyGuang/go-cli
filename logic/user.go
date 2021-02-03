package logic

import "bluebell/models"

// SignUp 此处通过参数校验，处理注册逻辑
func SignUp(p *models.UserSignUp) bool {
	if p.PassWord != p.RePassWord {
		return false
	} else {
		return true
	}
}

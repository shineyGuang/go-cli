package models

type UserSignUp struct {
	UserName   string `json:"user_name"`
	PassWord   string `json:"pwd"`
	RePassWord string `json:"re_pwd"`
}

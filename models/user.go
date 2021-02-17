package models

type UserSignUp struct {
	UserId     int64  `json:"user_id" db:"user_id"`
	UserName   string `json:"userName" db:"username" binding:"required"`
	PassWord   string `json:"pwd" db:"password" binding:"gte=3,lte=8,required"`
	RePassWord string `json:"re_pwd" binding:"required,eqfield=PassWord"`
	Email      string `json:"email" db:"email" binding:"required,email"`
	Gender     int    `json:"gender" db:"gender"`
}

type UserLogin struct {
	UserName string `json:"userName" db:"username" binding:"required"`
	PassWord string `json:"pwd" db:"password" binding:"gte=3,lte=8,required"`
}

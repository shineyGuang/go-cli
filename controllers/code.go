package controllers

type ResCode int64

const (
	CodeLoginUserNotExists ResCode = 1001 + iota
	CodeSignUpUserExists
	CodeInvalidParams
	CodeServerBusy
	CodeLoginFailed
	CodeSignUpFailed
	CodeSuccess = 0
)

var codeResMap = map[ResCode]string{
	CodeInvalidParams:      "参数错误！",
	CodeLoginUserNotExists: "登录用户不存在！",
	CodeSignUpUserExists:   "注册用户已存在！",
	CodeServerBusy:         "服务器繁忙",
	CodeSuccess:            "Success!",
	CodeLoginFailed:        "登录失败！",
	CodeSignUpFailed:       "注册失败！",
}

func (rc ResCode) getMsg() string {
	msg, ok := codeResMap[rc]
	if !ok {
		msg = codeResMap[CodeServerBusy]
	}
	return msg
}

package tool

const (
	SUCCSE = 200
	ERROR = 500

	//用户错误
	ERROR_USERNAME_USED = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_USER_EXIST = 1004
	NO_AUTHED = 1005
	AUTHED = 1006

	//文章错误
	ERROR_ARTICLE_NOT_EXIST = 2001
	ERROR_CHANNEL_USED = 2002
	ERROR_CHANNEL_NOT_EXIST = 2003
)

var codeMsg = map[int]string{
	SUCCSE: "成功",
	ERROR: "错误",

	//用户
	ERROR_USERNAME_USED: "用户名已存在！",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_USER_EXIST: "用户名已存在",
	NO_AUTHED: "认证失败",
	AUTHED: "认证成功",

	//文章
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
	ERROR_CHANNEL_USED:  "频道已存在",
	ERROR_CHANNEL_NOT_EXIST: "频道不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

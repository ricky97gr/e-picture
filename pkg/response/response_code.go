package response

const (
	ErrStruct = iota + 100000
	ErrDB
	ErrRegisterInfo
	ErrAuth
	ErrRedis
	ErrUserNameOrPassword
)

var errCodeMap = map[int32]struct {
	msgCn string
	msgEn string
}{
	200: {msgCn: "已成功", msgEn: "handle successfully"},

	ErrStruct:             {msgCn: "结构体有误", msgEn: "error struct"},
	ErrDB:                 {msgCn: "数据库有误", msgEn: "error db"},
	ErrRegisterInfo:       {msgCn: "注册信息格式有误", msgEn: "error register info format"},
	ErrAuth:               {msgCn: "认证信息有误，请重新登录", msgEn: "err auth info, please login again"},
	ErrRedis:              {msgCn: "redis数据库有误", msgEn: "error redis db"},
	ErrUserNameOrPassword: {msgCn: "用户名或者密码有误", msgEn: "username or password is wrong"},
}

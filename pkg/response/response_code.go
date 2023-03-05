package response

const (
	ErrStruct = iota + 100000
	ErrDB
	ErrRegisterInfo
)

var errCode = map[int32]struct {
	msgCn string
	msgEn string
}{
	200: {msgCn: "已成功", msgEn: "handle successfully"},

	ErrStruct:       {msgCn: "结构体有误", msgEn: "error struct"},
	ErrDB:           {msgCn: "数据库有误", msgEn: "error db"},
	ErrRegisterInfo: {msgCn: "注册信息格式有误", msgEn: "error register info format"},
}

package response

var errCode = map[int32]struct {
	msgCn string
	msgEn string
}{
	200: {msgCn: "已成功", msgEn: "handle successfully"},
}

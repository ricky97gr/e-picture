package uimodel

type Login struct {
	//phone
	UserName string `json:"userName"`
	ID       string `json:"id"`
	Passwd   string `json:"passwd"`
}

type Register struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

type RgisterConfirm struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
	Code  string `json:"code"`
}

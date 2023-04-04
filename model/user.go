package model

import "my-admin/model/uimodel"

type User struct {
	UUID     string `json:"uuid" gorm:"column:uuid;index;comment:the uuid of user"`
	UserName string `json:"userName" gorm:"column:userName;comment: the name of user"`
	Password string `json:"-" gorm:"column:password; comment: the password of user"`
	NickName string `json:"nickName" gorm:"column:nickName; comment: the nickName of user"`
	Theme    string `json:"theme" gorm:"column:theme; comment: the theme of user"`
	Avatar   string `json:"avatar" gorm:"column:avatar; comment: the avator of user"`
	Role     string `json:"role" gorm:"column:role; comment: the role of user"`
	RoleID   string `json:"roleID" gorm:"column:roleID; comment: the roleID of user"`
	Phone    string `json:"phone" gorm:"column:phone; comment: the phone of user"`
	Email    string `json:"email" gorm:"column:email; comment: the email of user"`
	Status   int    `json:"status" gorm:"column:status; comment: the status of user"`
}

func (u *User) Convert(info uimodel.Register) {
	u.UserName = info.Name
	u.Email = info.Email
	u.Phone = info.Phone
	u.Password = info.Passwd
}

func (User) TableName() string {
	return "user"
}

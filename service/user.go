package service

import (
	"errors"
	"my-admin/global"
	"my-admin/model"
	"my-admin/pkg/errs"
	"my-admin/pkg/password"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func insertUser(user model.User) error {
	return global.DBClient.Create(&user).Error
}

func UpdateUser() {

}

func DeleteUser() {

}

func UpdateUserEmail() {

}

func Register(user model.User) error {
	if IsUserExist(user) {
		return errs.ErrUserExist
	}
	user.UUID = uuid.New().String()
	//TODO:forgocode need generated passws
	user.Password = password.GetPassword(user.Password)
	if insertUser(user) != nil {
		return errs.ErrDbError
	}
	return nil
}

func IsUserExist(user model.User) bool {
	var record model.User
	result := global.DBClient.Where("phone = ?", user.Phone).First(&record)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func CheckPasswdByUser(user model.User) (string, bool) {
	var record model.User
	result := global.DBClient.Where("phone = ?", user.Phone).First(&record)
	if result.Error != nil {
		return "", false
	}
	return record.UserName, password.GetPassword(user.Password) == record.Password
}

func GetUserByPhone(phone string) model.User {
	var user model.User
	global.DBClient.Where("phone = ?", phone).First(&user)
	return user
}

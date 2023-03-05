package service

import (
	"errors"
	"my-admin/global"
	"my-admin/model"
	"my-admin/pkg/errs"

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
	user.Password = ""
	if insertUser(user) != nil {
		return errs.ErrDbError
	}
	return nil
}

func IsUserExist(user model.User) bool {
	var record model.User
	result := global.DBClient.Where("phone = ?", user.Phone).First(&record)
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

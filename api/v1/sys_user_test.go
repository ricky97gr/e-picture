package v1

import (
	"fmt"
	"my-admin/model"
	"my-admin/model/uimodel"
	"testing"
)

func TestConvert(t *testing.T) {
	info := uimodel.Register{
		Name:   "1",
		Phone:  "12345678",
		Email:  "test@qq.com",
		Passwd: "password",
	}
	var user model.User
	user.Convert(info)
	fmt.Printf("info:%+v, \n user:%+v\n", info, user)
}

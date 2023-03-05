package v1

import (
	"my-admin/model"
	"my-admin/model/uimodel"
	"my-admin/pkg/response"
	"my-admin/pkg/verify"
	"my-admin/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var info uimodel.Register
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.Failed(c, response.ErrStruct, "struct error")
		return
	}
	if !verify.IsEmailFormat(info.Email) {
		return

	}
	if !verify.IsMobileFormat(info.Phone) {
		return
	}
	var user model.User
	user.Convert(info)
	err = service.Register(user)
	if err != nil {
		return
	}
	response.Success(c)
	return

}

func Login(c *gin.Context) {

}

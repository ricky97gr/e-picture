package v1

import (
	"my-admin/global"
	"my-admin/middleware"
	"my-admin/model"
	"my-admin/model/uimodel"
	"my-admin/pkg/response"
	"my-admin/pkg/verify"
	"my-admin/service"
	"net/http"

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
	response.Success(c, "", 1)

}

func Login(c *gin.Context) {
	var info uimodel.Login
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.Failed(c, response.ErrStruct, "struct error")
		return
	}
	if info.Phone == "" {
		response.Failed(c, response.ErrStruct, "please input phone")
		return
	}
	user := model.User{
		Phone:    info.Phone,
		Password: info.Passwd,
	}
	userName, ok := service.CheckPasswdByUser(user)
	if !ok {
		service.SendOperationLog(userName, user.Phone, "Login", "登陆失败")
		response.Failed(c, response.ErrUserNameOrPassword)
		return
	}
	global.Logger.Infof("user <%s> login successfully\n", userName)
	token, err := middleware.GenerateToken(user.Phone, userName)
	if err != nil {
		response.Failed(c, response.ErrUserNameOrPassword)
		return
	}
	err = middleware.StoreToken(token)
	if err != nil {
		response.Failed(c, response.ErrRedis)
		return
	}
	service.SendOperationLog(userName, user.Phone, "Login", "登陆成功")
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  200,
			"msg":   "handle successfully",
			"token": token,
		})

}

func GetUserInfo(ctx *gin.Context) {
	phone := ctx.Request.Header.Get("userID")
	user := service.GetUserByPhone(phone)
	response.Success(ctx, user, 1)
}

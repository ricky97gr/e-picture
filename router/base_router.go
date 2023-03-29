package router

import (
	v1 "my-admin/api/v1"
	"my-admin/middleware"
)

func (r *Router) InitBaseRouter() {
	r.POST("/login", v1.Login)

}

func (r *Router) InitSystemRouter() {
	userGroup := r.Group("user")

	userGroup.Use(middleware.Auth())
	//TODO:test need delete
	userGroup.GET("/admin_register", v1.TestRespose)

	userGroup.POST("/register", v1.Register)

	userGroup.GET("/operationLog", v1.GetOperationLog)
	userGroup.GET("/info", v1.GetUserInfo)
	userGroup.POST("/upload", v1.UpLoadPhoto)

}

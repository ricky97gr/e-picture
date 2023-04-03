package router

import (
	v1 "my-admin/api/v1"
	"my-admin/middleware"
)

func (r *Router) InitPhotoRouter() {
	photoGroup := r.Group("photo")
	photoGroup.Use(middleware.Auth())
	photoGroup.GET("/list", v1.ListPhoto)
}

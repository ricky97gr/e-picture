package router

import (
	//"my-admin/middleware"
	// "my-admin/middleware"
	"my-admin/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin.RouterGroup
}

func StartServer() error {
	// engine := gin.New()

	// engine.Use(middleware.Logger(), gin.Recovery())
	engine := gin.Default()
	//TODO:(forgocode) need to update path and dir
	engine.StaticFS("/file", http.Dir(""))
	r := new(Router)
	r.RouterGroup = engine.RouterGroup
	engine.GET("/health", func(ctx *gin.Context) {
		response.Success(ctx, "system is ready", 1)
	})
	r.InitBaseRouter()
	r.InitSystemRouter()
	r.InitBucketRouter()
	r.InitPhotoRouter()

	//桶的增删改查
	//图片的增删改查
	//用户的增删改查
	//权限设置

	return engine.Run(":8080")

}

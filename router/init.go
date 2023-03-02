package router

import (
	//"my-admin/middleware"
	"my-admin/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin.RouterGroup
}

func StartServer() error {
	// engine := gin.New()

	// engine.Use(middleware.Logger(), middleware.Recovery())
	engine := gin.Default()
	//TODO:(forgocode) need to update path and dir
	engine.StaticFS("/file", http.Dir(""))
	r := new(Router)
	r.RouterGroup = engine.RouterGroup
	engine.GET("/health", func(ctx *gin.Context) {
		response.Success(ctx, "system is ready")
	})
	r.InitBaseRouter()
	//engine.POST("login")

	// InitRouter()
	// InitRouter()
	// baseRouter := engine.Group("base")

	// systemGroup := engine.Group("system")

	return engine.Run(":8080")

}

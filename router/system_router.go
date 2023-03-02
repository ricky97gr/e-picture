package router

import "github.com/gin-gonic/gin"

type SystemRouter struct {
	gin.RouterGroup
}

func (r *SystemRouter) InitRouter() {
	r.Group("system")
}

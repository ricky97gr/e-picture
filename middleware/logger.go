package middleware

import (
	"my-admin/global"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		host := ctx.RemoteIP()
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		ctx.Next()
		raw := ctx.Request.URL.RawQuery
		status := ctx.Writer.Status()
		global.Logger.Debugf("src ip: %s | path: %s | method: %s | consume: %+v | payload: %s | status: %d", host, path, method, time.Now().Sub(start), raw, status)
	}
}

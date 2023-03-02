package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, detail interface{}) {

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":   200,
			"msg":    "handle successfully",
			"detail": detail,
		})
}

func Failed(ctx *gin.Context, errCode int32, detail interface{}) {

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":   errCode,
			"msg":    "",
			"detail": detail,
		})
}

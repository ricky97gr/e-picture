package v1

import (
	"my-admin/pkg/response"
	"my-admin/service"

	"github.com/gin-gonic/gin"
)

func TestRespose(ctx *gin.Context) {
	service.SendOperationLog(ctx, "Test", "测试日志")
	response.Success(ctx, "test successfully", 1)
}

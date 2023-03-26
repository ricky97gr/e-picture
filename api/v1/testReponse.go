package v1

import (
	"my-admin/pkg/response"
	"my-admin/service"

	"github.com/gin-gonic/gin"
)

func TestRespose(ctx *gin.Context) {
	service.SendOperationLog("admin", "1234567", "Test", "测试日志")
	response.Success(ctx, "test successfully", 1)
}

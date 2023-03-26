package v1

import (
	"my-admin/pkg/response"
	"my-admin/service"

	"github.com/gin-gonic/gin"
)

func GetOperationLog(ctx *gin.Context) {
	result := service.GetOperationLog()
	response.Success(ctx, result, len(result))
}

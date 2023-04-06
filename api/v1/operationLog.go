package v1

import (
	"my-admin/pkg/page"
	"my-admin/pkg/response"
	"my-admin/service"

	"github.com/gin-gonic/gin"
)

func GetOperationLog(ctx *gin.Context) {
	result := service.GetOperationLog()
	page.GetPagination(ctx)
	response.Success(ctx, result, len(result))
}

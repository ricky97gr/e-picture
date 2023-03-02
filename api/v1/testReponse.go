package v1

import (
	"my-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

func TestRespose(ctx *gin.Context) {
	response.Success(ctx, "test successfully")
}

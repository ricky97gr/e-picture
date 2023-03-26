package v1

import (
	"my-admin/model"
	"my-admin/model/uimodel"
	"my-admin/pkg/response"
	"my-admin/service"

	"github.com/gin-gonic/gin"
)

func CreateBucket(ctx *gin.Context) {
	var info uimodel.Bucket
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	var bucket model.Bucket
	bucket.Covert(info)

	userID := ctx.Request.Header.Get("userID")
	userName := ctx.Request.Header.Get("userName")
	bucket.OwnerID = userID
	bucket.OwnerName = userName
	err = service.CreateBucket(bucket)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "create bucket successfully", 1)
}

func UpdateBucket(ctx *gin.Context) {

}

func DeleteBucket(ctx *gin.Context) {
	var info uimodel.Bucket
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	var bucket model.Bucket
	bucket.Covert(info)
	userID := ctx.Request.Header.Get("userID")
	bucket.OwnerID = userID
	err = service.DeleteBucket(bucket)
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "create bucket successfully", 1)

}

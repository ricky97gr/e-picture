package v1

import (
	"fmt"
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
	//TODO: 需要区分db还是minio的error
	err = service.CreateBucket(bucket)
	if err != nil {
		response.Failed(ctx, response.ErrDB, err)
		return
	}
	service.SendOperationLog(ctx, "Bucket", fmt.Sprintf("创建存储桶{ %s }", bucket.Name))
	response.Success(ctx, "create bucket successfully", 1)
}

func GetBucketList(ctx *gin.Context) {
	buckets := service.ListBucket()
	response.Success(ctx, buckets, len(buckets))
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
	//TODO: 需要区分db还是minio的error
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	service.SendOperationLog(ctx, "Bucket", fmt.Sprintf("删除存储桶{ %s }", bucket.Name))
	response.Success(ctx, "create bucket successfully", 1)

}

func GetBucketCount(ctx *gin.Context) {
	count, err := service.GetBucketCount()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "", int(count))

}

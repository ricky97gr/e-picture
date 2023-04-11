package v1

import (
	"context"
	"fmt"
	"my-admin/global"
	"my-admin/model"
	"my-admin/pkg/response"
	"my-admin/pkg/uuid"
	"my-admin/service"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func UpLoadPhoto(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Printf("failed to get multipart form, err:%+v\n", err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	fmt.Printf("receive form:%+v\n", form)
	photoName := form.Value["name"][0]
	bucketName := form.Value["bucketName"][0]
	if bucketName == "" {
		response.Failed(ctx, response.ErrStruct, "桶名称不能为空")
		return
	}

	if photoName != "" {
		if len(form.File["file"]) != 1 {
			response.Failed(ctx, response.ErrStruct, "文件命名时，上传文件数不能超过1")
			return
		}
	}
	bucket, err := service.GetBucketByName(bucketName)
	if err != nil {
		fmt.Printf("failed to get bucket by bucketName:%s, err:%+v\n", bucketName, err)
		response.Failed(ctx, response.ErrDB, "查询桶信息异常")
		return
	}

	files := form.File["file"]
	//BUG: 上传多个文件时有可能会出现同名的照片
	for _, file := range files {
		fmt.Println(file.Filename)
		var newFile string
		f, err := file.Open()
		if err != nil {
			global.Logger.Infof("failed to insert file, err:%+v\n", err)
		}
		newFile = fmt.Sprintf("%d", time.Now().UnixMilli()) + path.Ext(file.Filename)
		if photoName != "" {
			newFile = photoName + path.Ext(file.Filename)
		}
		if service.IsPhotoOutOfCapacity(bucket, file.Size) {
			response.Failed(ctx, response.ErrDB, "桶容量超出上限")
			return
		}
		service.CreatePhoto(model.Photo{
			UUID:       uuid.GetUUID(),
			Name:       newFile,
			BucketName: bucketName,
			Size:       file.Size,
			Url:        fmt.Sprintf("http://%s:%d/%s/%s", global.Config.Minio.IP, global.Config.Minio.Port, bucketName, newFile),
		})
		err = service.AddPhotoCapacityByBucketName(bucketName, file.Size)
		if err != nil {
			global.Logger.Errorf("failed to update(add photo) bucket used info, err:%+v\n", err)
			response.Failed(ctx, response.ErrDB)
			return
		}
		_, err = global.MinioClient.PutObject(context.Background(), bucketName, newFile, f, file.Size, minio.
			PutObjectOptions{})
		if err != nil {
			fmt.Printf("failed to upload file to minio, err:%+v\n", err)
			continue
		}
		service.SendOperationLog(ctx, "Photo", fmt.Sprintf("上传了照片 %s", newFile))
	}

	response.Success(ctx, "上传成功", len(files))

}

func DeletePhoto(ctx *gin.Context) {
	var info model.Photo
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.Failed(ctx, response.ErrStruct)
		return
	}
	photo, ok := service.GetPhotoByUUID(info.UUID)
	if !ok {
		response.Failed(ctx, response.ErrRecordNotFound)
		return
	}
	err = global.MinioClient.RemoveObject(context.Background(), photo.BucketName, photo.Name, minio.RemoveObjectOptions{})
	if err != nil {
		global.Logger.Errorf("failed to remove file, info:%+v, err:%+v\n", photo, err)
	}
	err = service.DeletePhoto(photo)
	if err != nil {
		global.Logger.Errorf("failed to delete file record from db, info:%+v, err:%+v\n", photo, err)
	}
	err = service.DeletePhotoCapacityByBucketName(photo.BucketName, photo.Size)
	if err != nil {
		global.Logger.Errorf("failed to update(delete photo) bucket used info, err:%+v\n", err)
	}
	service.SendOperationLog(ctx, "Photo", fmt.Sprintf("删除了照片 %s", photo.Name))
	response.Success(ctx, "删除成功", 1)

}

func ListPhoto(ctx *gin.Context) {
	photos := service.ListPhoto()
	response.Success(ctx, photos, len(photos))
}

func GetPhotoCount(ctx *gin.Context) {
	count, err := service.GetPhotoCount()
	if err != nil {
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, "", int(count))
}

func GetPhotoInfo(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == ""{
		response.Failed(ctx, response.ErrStruct, "failed to get photo by uuid")
		return
	}
	photo, ok := service.GetPhotoByUUID(uuid)
	if !ok{
		response.Failed(ctx, response.ErrDB)
		return
	}
	response.Success(ctx, photo, 1)
}

package v1

import (
	"fmt"
	"my-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

func UpLoadPhoto(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		fmt.Println(file.Filename)
		ctx.SaveUploadedFile(file, "/root/tmp/"+file.Filename)
	}

	response.Success(ctx, "上传成功", len(files))

}

func DeletePhoto(ctx *gin.Context) {

}

package v1

import (
	"fmt"
	"my-admin/pkg/response"

	"github.com/gin-gonic/gin"
)

func UpLoadPhoto(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Printf("failed to get multipart form, err:%+v\n", err)
		response.Failed(ctx, response.ErrStruct)
		return
	}
	fmt.Printf("receive form:%+v\n", form)
	files := form.File["file"]
	for _, file := range files {
		fmt.Println(file.Filename)
		ctx.SaveUploadedFile(file, "/root/tmp/"+file.Filename)
	}

	response.Success(ctx, "上传成功", len(files))

}

func DeletePhoto(ctx *gin.Context) {

}

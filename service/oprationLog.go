package service

import (
	"my-admin/global"
	"my-admin/model"
	"my-admin/pkg/uuid"

	"github.com/gin-gonic/gin"
	// "time"
)

func SendOperationLog(ctx *gin.Context, moduleName, message string) {
	userName := ctx.Request.Header.Get("userName")
	userID := ctx.Request.Header.Get("userID")
	log := &model.OperationLog{
		UUID: uuid.GetUUID(),
		// CreateTime: time.Now(),
		Module:   moduleName,
		UserName: userName,
		UserID:   userID,
		Message:  message,
		IP:       ctx.ClientIP(),
	}
	insertOperationLog(*log)
}
func DeleteOperationLog(uuid string) {

}

func GetOperationLog() []model.OperationLog {
	var logs []model.OperationLog
	global.DBClient.Find(&logs)
	return logs
}

func insertOperationLog(log model.OperationLog) {
	global.DBClient.Create(&log)
}

func deleteOperationLogByUUID(uuid string) {
	//global.DBClient.Delete()
}

func GetOperationCount() (int64, error) {
	var count int64
	result := global.DBClient.Model(&model.OperationLog{}).Count(&count)
	return count, result.Error
}

package service

import (
	"my-admin/global"
	"my-admin/model"
	"my-admin/pkg/uuid"
	"time"
)

func SendOperationLog(userName, userID, moduleName, message string) {
	log := &model.OperationLog{
		UUID:       uuid.GetUUID(),
		CreateTime: time.Now(),
		Module:     moduleName,
		UserName:   userName,
		UserID:     userID,
		Message:    message,
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

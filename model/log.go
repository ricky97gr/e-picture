package model

import "time"

type OperationLog struct {
	CreateTime time.Time `json:"createTime"`
	UUID       string    `json:"uuid"`
	Module     string    `json:"module"`
	Message    string    `json:"message"`
	UserName   string    `json:"userName"`
	UserID     string    `json:"userID"`
}

func (OperationLog) TableName() string {
	return "operationLog"
}

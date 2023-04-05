package model

import (
	"my-admin/pkg/time"
)

type OperationLog struct {
	CreateTime *time.LocalTime `json:"createTime" gorm:"column:createTime; autoCreateTime"`
	UUID       string          `json:"uuid" gorm:"column:uuid"`
	Module     string          `json:"module" gorm:"column:module"`
	Message    string          `json:"message" gorm:"column:message"`
	UserName   string          `json:"userName" gorm:"column:userName"`
	UserID     string          `json:"userID" gorm:"column:userID"`
	IP         string          `json:"ip" gorm:"column:ip"`
	Place      string          `json:"place" gorm:"column:place"`
}

func (OperationLog) TableName() string {
	return "operationLog"
}

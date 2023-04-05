package model

import (
	"my-admin/pkg/time"
)

type Photo struct {
	UUID       string          `json:"uuid" gorm:"column:uuid"`
	Name       string          `json:"name" gorm:"column:name"`
	BucketID   string          `json:"bucketID" gorm:"column:bucketID"`
	BucketName string          `json:"bucketName" gorm:"column:bucketName"`
	Url        string          `json:"url" gorm:"column:url"`
	Status     string          `json:"status" gorm:"column:status"`
	CreateTime *time.LocalTime `json:"createTime" gorm:"column:createTime; autoCreateTime"`
}

func (Photo) TableName() string {
	return "photo"
}

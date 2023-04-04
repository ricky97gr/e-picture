package model

import "my-admin/model/uimodel"

type Bucket struct {
	ID           int64  `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:name"`
	CapacityMode string `json:"capacityMode" gorm:"column:capacityMode"`
	Status       string `json:"status" gorm:"column:status"`
	OwnerID      string `json:"ownerID" gorm:"column:ownerID"`
	OwnerName    string `json:"ownerName" gorm:"column:ownerName"`
	TotalSize    string `json:"totalSize" gorm:"column:totalSize"`
	UsedSize     string `json:"usedSize" gorm:"column:usedSize"`
	TotalNumber  int32  `json:"totalNumber" gorm:"column:totalNumber"`
	UsedNumber   int32  `json:"usedNumber" gorm:"column:usedNumber"`
}

func (b *Bucket) Covert(bucket uimodel.Bucket) {
	b.Name = bucket.Name
	b.TotalSize = bucket.TotalSize
	b.CapacityMode = bucket.CapacityMode
	b.TotalNumber = bucket.TotalNumber
	b.Status = bucket.Status
	b.ID = int64(bucket.ID)
}

func (Bucket) TableName() string {
	return "bucket"
}

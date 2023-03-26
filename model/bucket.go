package model

import "my-admin/model/uimodel"

type Bucket struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Size         string `json:"size"`
	CapacityMode string `json:"capacityMode"`
	Number       int32  `json:"number"`
	Status       string `json:"status"`
	OwnerID      string `json:"ownerID"`
	OwnerName    string `json:"ownerName"`
}

func (b *Bucket) Covert(bucket uimodel.Bucket) {
	b.Name = bucket.Name
	b.Size = bucket.Size
	b.CapacityMode = bucket.CapacityMode
	b.Number = bucket.Number
	b.Status = bucket.Status
}

func (Bucket) TableName() string {
	return "bucket"
}

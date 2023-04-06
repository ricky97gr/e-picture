package uimodel

import (
	"my-admin/pkg/time"
)

type Bucket struct {
	Name         string          `json:"name"`
	TotalSize    int64           `json:"totalSize"`
	CapacityMode string          `json:"capacityMode"`
	TotalNumber  int32           `json:"totalNumber"`
	Status       string          `json:"status"`
	ID           int32           `json:"id"`
	UsedSize     string          `json:"usedSize"`
	UsedNumber   int32           `json:"usedNumber"`
	CreateTime   *time.LocalTime `json:"createTime"`
}

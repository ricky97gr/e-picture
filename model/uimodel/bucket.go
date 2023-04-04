package uimodel

type Bucket struct {
	Name         string `json:"name"`
	TotalSize    string `json:"totalSize"`
	CapacityMode string `json:"capacityMode"`
	TotalNumber  int32  `json:"totalNumber"`
	Status       string `json:"status"`
	ID           int32  `json:"id"`
	UsedSize     string `json:"usedSize"`
	UsedNumber   int32  `json:"usedNumber"`
}

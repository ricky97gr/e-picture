package uimodel

type Bucket struct {
	Name         string `json:"name"`
	Size         string `json:"size"`
	CapacityMode string `json:"capacityMode"`
	Number       int32  `json:"number"`
	Status       string `json:"status"`
	ID           int32  `json:"id"`
}

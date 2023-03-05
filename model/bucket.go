package model

type Bucket struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Size         string `json:"size"`
	CapacityMode string `json:"capacityMode"`
	Number       int32  `json:"number"`
	Status       string `json:"status"`
	OwnerID      string `json:"ownerID"`
	OwnerName    string `json:"ownerName"`
}

func (Bucket) TableName() string {
	return "bucket"
}

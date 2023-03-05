package model

type Photo struct {
	UUID       string `json:"uuid" gorm:"uuid"`
	Name       string `json:"name"`
	BucketID   string `json:"bucketID" gorm:"index"`
	BucketName string `json:"bucketName"`
	Url        string `json:"url"`
	Status     string `json:"status"`
}

func (Photo) TableName() string {
	return "photo"
}

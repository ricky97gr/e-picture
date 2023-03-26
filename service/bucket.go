package service

import (
	"my-admin/global"
	"my-admin/model"
	"my-admin/pkg/errs"
)

const (
	BucketCapacity = 32
)

func CreateBucket(bucket model.Bucket) error {
	bucketNumber := getBucketCount()
	if bucketNumber >= BucketCapacity {
		return errs.ErrOutOfCapacity
	}
	bucket.ID = bucketNumber + 1

	insertBucket(bucket)
	return nil
}

func DeleteBucket(bucket model.Bucket) error {
	b := getBucketByID(bucket.ID)
	if b == nil {
		return errs.ErrDbError
	}
	if b.OwnerID != bucket.OwnerID {

	}
	return deleteBucket(bucket)
}

func insertBucket(bucket model.Bucket) error {
	return global.DBClient.Create(&bucket).Error
}

func deleteBucket(bucket model.Bucket) error {
	return global.DBClient.Delete(bucket).Error
}

func getBucketCount() int64 {
	result := global.DBClient.Find(&model.Bucket{})
	return result.RowsAffected
}

func getBucketByID(id int64) *model.Bucket {
	var bucket model.Bucket
	result := global.DBClient.Find(&model.Bucket{}).Where("id = ?", id).First(&bucket)
	if result.Error != nil {
		return nil
	}
	return &bucket
}

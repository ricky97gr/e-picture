package service

import (
	"context"
	"errors"
	"my-admin/global"
	"my-admin/model"
	"my-admin/pkg/errs"

	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

const (
	BucketCapacity = 32
)

func CreateBucket(bucket model.Bucket) error {
	bucketNumber := getBucketCount()
	if bucketNumber >= BucketCapacity {
		return errs.ErrOutOfCapacity
	}
	if isBucketExist(bucket.Name) {
		return errs.ErrBucketExisted
	}
	bucket.ID = bucketNumber + 1

	if err := insertBucket(bucket); err != nil {
		return errs.ErrDbError
	}
	return createStoreBucket(bucket.Name)
}

func DeleteBucket(bucket model.Bucket) error {
	// b := getBucketByID(bucket.ID)
	// if b == nil {
	// 	return errs.ErrDbError
	// }
	// if b.OwnerID != bucket.OwnerID {
	// 	return errs.ErrBucketOwner
	// }
	if err := deleteBucket(bucket); err != nil {
		return errs.ErrDbError
	}
	return deleteStoreBucket(bucket.Name)

}

func ListBucket() []model.Bucket {
	var buckets []model.Bucket
	global.DBClient.Find(&buckets)
	return buckets
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

func isBucketExist(bucketName string) bool {
	var bucket model.Bucket
	result := global.DBClient.Where("name = ?", bucketName).First(&bucket)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func getBucketByID(id int64) *model.Bucket {
	var bucket model.Bucket
	result := global.DBClient.Find(&model.Bucket{}).Where("id = ?", id).First(&bucket)
	if result.Error != nil {
		return nil
	}
	return &bucket
}

func createStoreBucket(bucketName string) error {
	ctx := context.Background()
	return global.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{ObjectLocking: false})
}

func deleteStoreBucket(bucketName string) error {
	ctx := context.Background()
	return global.MinioClient.RemoveBucket(ctx, bucketName)
}

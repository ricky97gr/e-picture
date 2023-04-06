package service

import (
	"context"
	"errors"
	"fmt"
	"my-admin/global"
	"my-admin/model"
	"my-admin/pkg/errs"

	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

const (
	BucketCapacity = 32
)

type BucketPolicy string

const (
	BucketPolicyReadOnly = `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetBucketLocation"],"Resource":["arn:aws:s3:::test"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:ListBucket"],"Resource":["arn:aws:s3:::%s"],"Condition":{"StringEquals":{"s3:prefix":["*"]}}},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::%s/**"]}]}`
	// BucketPolicyWriteRead = `
	// {
	// 	"Version": "2012-10-17",
	// 	"Statement": [
	// 		{
	// 			"Effect": "Allow",
	// 			"Action": [
	// 				"s3:GetBucketLocation",
	// 				"s3:GetObject"
	// 			],
	// 			"Resource": [
	// 				"arn:aws:s3:::*"
	// 			]
	// 		}
	// 	]
	// }
	// `
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

func updateBucketUsedInfo(bucket model.Bucket) error {
	return global.DBClient.Model(&model.Bucket{}).Where("name = ?", bucket.Name).Updates(model.Bucket{UsedNumber: bucket.UsedNumber, UsedSize: bucket.UsedSize}).Error
}

func isBucketExist(bucketName string) bool {
	var bucket model.Bucket
	result := global.DBClient.Where("name = ?", bucketName).First(&bucket)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func GetBucketByName(name string) (model.Bucket, error) {
	var bucket model.Bucket
	result := global.DBClient.Find(&model.Bucket{}).Where("name = ?", name).First(&bucket)
	if result.Error != nil {
		return bucket, result.Error
	}
	return bucket, nil
}

func createStoreBucket(bucketName string) error {
	ctx := context.Background()
	err := global.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{ObjectLocking: false})
	if err != nil {
		return err
	}
	return addPolicyToBucket(bucketName, fmt.Sprintf(BucketPolicyReadOnly, bucketName, bucketName))
}

func deleteStoreBucket(bucketName string) error {
	ctx := context.Background()
	return global.MinioClient.RemoveBucket(ctx, bucketName)
}

func addPolicyToBucket(bucketName, policy string) error {
	return global.MinioClient.SetBucketPolicy(context.Background(), bucketName, policy)
}

func GetBucketCount() (int64, error) {
	var count int64
	result := global.DBClient.Model(&model.Bucket{}).Count(&count)
	return count, result.Error
}

func AddPhotoCapacityByBucketName(bucketName string, photoSize int64) error {
	var bucket model.Bucket
	result := global.DBClient.Model(&model.Bucket{}).Where("name = ?", bucketName).First(&bucket)
	if result.Error != nil {
		return result.Error
	}
	bucket.UsedNumber = bucket.UsedNumber + 1
	bucket.UsedSize = bucket.UsedSize + photoSize
	return updateBucketUsedInfo(bucket)
}

func DeletePhotoCapacityByBucketName(bucketName string, photoSize int64) error {
	var bucket model.Bucket
	result := global.DBClient.Model(&model.Bucket{}).Where("name = ?", bucketName).First(&bucket)
	if result.Error != nil {
		return result.Error
	}
	bucket.UsedNumber = bucket.UsedNumber - 1
	bucket.UsedSize = bucket.UsedSize - photoSize
	//NOTE: 应该不需要，是个保障
	if bucket.UsedSize < 0 {
		bucket.UsedSize = 0
	}
	return updateBucketUsedInfo(bucket)

}

func GetBucketUsedSize() (int64, error) {
	var sizes []int64
	result := global.DBClient.Model(&model.Bucket{}).Pluck("usedSize", &sizes)
	sum := int64(0)
	for _, size := range sizes {
		sum = sum + size
	}
	return sum, result.Error
}

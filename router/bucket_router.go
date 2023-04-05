package router

import (
	v1 "my-admin/api/v1"
	"my-admin/middleware"
)

func (r *Router) InitBucketRouter() {
	bucketGroup := r.Group("bucket")

	bucketGroup.Use(middleware.Auth())
	bucketGroup.POST("/create", v1.CreateBucket)
	bucketGroup.DELETE("/delete", v1.DeleteBucket)
	bucketGroup.GET("/list", v1.GetBucketList)
	bucketGroup.GET("/count", v1.GetBucketCount)
}

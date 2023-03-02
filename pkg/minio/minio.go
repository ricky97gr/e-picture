package minio

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio(ip string, port uint16, accessKeyID, secretAccessKey string, useSSL bool) (*minio.Client, error) {
	endpoint := fmt.Sprintf("%s:%d", ip, port)
	// Initialize minio client object.
	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
}

package minio

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
)

func TestInitMinio(t *testing.T) {
	c, err := InitMinio("10.182.34.112", 9000, "K62MVxUhRgCO0dtd", "jH82hj7TvFewMQ8zCLVAfv6tBt6s4xuT", false)
	if err != nil {
		fmt.Printf("failed to init minio client, err: %+v\n", err)
		return
	}
	l, err := c.ListBuckets(context.Background())
	if err != nil {
		fmt.Printf("failed to list bucket, err:%+v\n", err)
		return
	}
	fmt.Printf("bucket list:%+v\n", l)
	f, err := os.Open("images.jpg")
	if err != nil {
		fmt.Printf("failed to open file:%s, err:%+v\n", "", err)
		return
	}
	defer f.Close()
	fStat, err := f.Stat()
	if err != nil {
		fmt.Printf("failed to get file stat, err:%+v\n", err)
		return
	}
	ti := time.Now().UnixMilli()
	result, err := c.PutObject(context.Background(), l[0].Name, fmt.Sprintf("%d.jpg", ti), f, fStat.Size(), minio.PutObjectOptions{})

	if err != nil {
		fmt.Printf("failed to up load .jpg, err:%+v\n", err)
	}
	fmt.Println("Uploaded:", "my-objectname", " of size: ", result, "Successfully.")
}

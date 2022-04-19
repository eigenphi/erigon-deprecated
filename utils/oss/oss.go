package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type ParquetSaver struct {
	bucket *oss.Bucket
}

var (
	ParquetBucketName = "default-erigon-parquet"
	objectIdFunc      = func(height int64) string {
		return fmt.Sprintf("%d", height)
	}
)

func NewParquetSaver(endpoint, accessKeyId, accessKeySecret string) (*ParquetSaver, error) {
	return newParquetSaver(endpoint, accessKeyId, accessKeySecret)
}

func newParquetSaver(endpoint, accessKeyId, accessKeySecret string) (*ParquetSaver, error) {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(ParquetBucketName)
	if err != nil {
		return nil, fmt.Errorf("get bucket %s: %v", ParquetBucketName, err)
	}
	return &ParquetSaver{
		bucket: bucket,
	}, nil
}

func (p *ParquetSaver) Save(heihgt int64, reader io.Reader) error {
	if err := p.bucket.PutObject(objectIdFunc(heihgt), reader); err != nil {
		return err
	}
	return nil
}

package oss

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type ParquetSaver struct {
	bucket *oss.Bucket
}

const (
	ParquetBucketName = "erigon-parquet"
)

var (
	objectIdFunc = func(height int64) string {
		return fmt.Sprintf("%d", height)
	}
)

func NewParquetSaver(endpoint, accessKeyId, accessKeySecret string) (*ParquetSaver, error) {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(ParquetBucketName)
	if err != nil {
		return nil, err
	}
	return &ParquetSaver{
		bucket: bucket,
	}, nil
}

func (p *ParquetSaver) Save(heihgt int64, data []byte) error {
	return p.bucket.PutObject(objectIdFunc(heihgt), bytes.NewReader(data))
}

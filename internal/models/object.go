package models

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const DefaultBucketName = "bucket"

// S3 implementations must meet this minimum interface - useful for testing
type S3API interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

// Used to structure the database implementations in ./stores
type ObjectStorage interface {
	Upload(ctx context.Context, key string, data io.Reader) error
	Download(ctx context.Context, key string) (io.Reader, error)
	Delete(ctx context.Context, key string) error
	List(ctx context.Context, prefix string) ([]string, error)
}

type S3InitBucketEnv struct {
	AwsAccessKeyId     string
	AwsS3Region        string
	AwsSecretAccessKey string
	AwsS3BucketName    string
}

type R2InitBucketEnv struct {
	CloudflareAccountId string
	AwsAccessKeyId      string
	AwsSecretAccessKey  string
	AwsS3BucketName     string
}
type MinioInitBucketEnv struct {
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	AwsEndpoint        string
	AwsS3Region        string
}
